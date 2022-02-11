package session

import (
	"container/heap"
	"fmt"
	"gonum.org/v1/gonum/stat/combin"
	"math"
	"mdp_algo/common"
)

type costModel struct {
	moveCost map[common.Move]int
}

var moveCost = map[common.Move]int{
	common.Forward:       1,
	common.Backward:      2,
	common.ForwardLeft:   10,
	common.ForwardRight:  10,
	common.BackwardLeft:  11,
	common.BackwardRight: 11,
}

var hamiltonCostModel = costModel{
	moveCost: moveCost,
}

type node struct {
	position         common.Position
	shortestDistance int
	shortestFrom     *node
	visited          bool
}

func (n *node) Less(other common.Node) bool {
	return n.shortestDistance < other.(*node).shortestDistance
}

type expansion struct {
	position common.Position
	distance int
}

// cellsToTakeImage find cells to take image from the image cell
func (s *sessionImpl) cellsToTakeImage(cell common.Cell, cs cellState) (positions []common.Position) {
	currentPos := common.Position{
		Cell:      cell,
		Direction: cs.direction,
	}

	// Positions aligned with the image obstacles that can take image
	alignPositions := []common.Position{
		currentPos.Transition(common.Forward, common.Forward, common.Forward),
		currentPos.Transition(common.Forward, common.Forward, common.Forward, common.Forward),
	}

	robotDirection := cs.direction.Opposite()
	// Add variations to the positions because robot do not need to be aligned to take image
	variation := 0
	for _, alignPosition := range alignPositions {
		for i := -variation; i <= variation; i++ {
			tmp := alignPosition
			if tmp.Direction.IsVertical() {
				tmp.Cell.Xcoord += i
			} else {
				tmp.Cell.Ycoord += i
			}

			// Update robot direction to face the image
			tmp.Direction = robotDirection
			positions = append(positions, tmp)
		}
	}

	return
}

// initTakePositions find optimal positions to take imageCells
func (s *sessionImpl) initTakePositions() {
	takePositions := make(map[common.Cell][]common.Position)
	s.imageCells = nil

	// Use virtual arena to get the obstacles and find optimal cells to take picture
	for Ycoord := 0; Ycoord < s.height; Ycoord++ {
		for Xcoord := 0; Xcoord < s.width; Xcoord++ {
			cell := common.Cell{
				Xcoord: Xcoord,
				Ycoord: Ycoord,
			}
			cs := s.getVirtualCellState(cell)
			if cs.state == obstacle && cs.hasFace {
				takePositions[cell] = s.cellsToTakeImage(cell, cs)
				s.imageCells = append(s.imageCells, cell)
			}
		}
	}
	s.takePositions = takePositions
}

func (s *sessionImpl) setPlanCache(start common.Position, imageCell common.Cell, plan Plan) {
	if _, exist := s.planCache[start]; !exist {
		s.planCache[start] = make(map[common.Cell]Plan)
	}
	s.planCache[start][imageCell] = plan
}

func (s *sessionImpl) getPlanCache(start common.Position, imageCell common.Cell) (Plan, bool) {
	if _, exist := s.planCache[start]; !exist {
		return Plan{}, false
	}
	if _, exist := s.planCache[start][imageCell]; !exist {
		return Plan{}, false
	}
	return s.planCache[start][imageCell], true
}

func (s *sessionImpl) HamiltonPath() HamiltonPlan {
	// Init plan cache
	s.planCache = make(map[common.Position]map[common.Cell]Plan)

	optimalCost := math.MaxInt64
	var optimalPath common.Path
	var optimalOrder []common.Cell
	var optimalDetectImageIndices []int

	s.initTakePositions()
	// Generate permutations of image obstacles
	imageNum := len(s.imageCells)
	perms := combin.Permutations(imageNum, imageNum)

	for permIndex, perm := range perms {
		var imageCellOrders []common.Cell
		for _, i := range perm {
			imageCellOrders = append(imageCellOrders, s.imageCells[i])
		}

		start := s.current
		cost := 0
		path := common.Path{start}
		var detectImageIndices []int

		// Use shortest path algo to find path from one imageCell to another
		for _, imageCell := range imageCellOrders {
			takePositions := s.takePositions[imageCell]
			goalTest := func(position common.Position) bool {
				for _, takePosition := range takePositions {
					if position == takePosition {
						return true
					}
				}
				return false
			}
			moveableTest := func(cell common.Cell) bool {
				cs := s.getAugmentCellState(cell)
				return cs.state == empty
			}

			// Get plan from cache
			localPlan, exist := s.getPlanCache(start, imageCell)
			if !exist {
				localPlan = s.shortestPath(
					start, goalTest, len(takePositions), moveableTest, hamiltonCostModel,
				)
				s.setPlanCache(start, imageCell, localPlan)
			}

			if localPlan.Cost == math.MaxInt64 { // check if local plan fails
				cost = math.MaxInt64
				break
			}
			cost += localPlan.Cost
			path = append(path, localPlan.Path[1:]...)

			// The robot will detect image at these indices
			detectImageIndices = append(detectImageIndices, len(path)-2)

			start = localPlan.Path[len(localPlan.Path)-1]
		}

		// Update with best path
		if cost < optimalCost {
			optimalCost = cost
			optimalPath = path
			optimalOrder = imageCellOrders
			optimalDetectImageIndices = detectImageIndices
		}
		fmt.Println("Permutation", permIndex)
	}

	return HamiltonPlan{
		Cost:               optimalCost,
		Path:               optimalPath,
		Order:              optimalOrder,
		DetectImageIndices: optimalDetectImageIndices,
	}
}

func (s *sessionImpl) expand(
	current common.Position,
	moveableTest func(common.Cell) bool,
	model costModel,
) (next []expansion) {
	for _, move := range common.AllMoves {
		if !current.IsValidMove(move, moveableTest) {
			continue
		}

		nextPosition := current.Transition(move)
		cost := model.moveCost[move]
		next = append(next, expansion{
			position: nextPosition,
			distance: cost,
		})
	}

	return
}

func (s *sessionImpl) shortestPath(
	start common.Position,
	goalTest func(common.Position) bool,
	goalNum int,
	moveableTest func(common.Cell) bool,
	costModel costModel,
) Plan {
	// Preparation
	moveableNode := make(map[common.Position]*node)
	unvisitedPQ := common.NewPriorityQueue()
	heap.Init(unvisitedPQ)

	for Ycoord := 0; Ycoord < s.height; Ycoord++ {
		for Xcoord := 0; Xcoord < s.width; Xcoord++ {
			cell := common.Cell{
				Xcoord: Xcoord,
				Ycoord: Ycoord,
			}
			if moveableTest(cell) {
				for _, direction := range common.AllDirections {
					position := common.Position{
						Cell:      cell,
						Direction: direction,
					}
					node := &node{
						position:         position,
						shortestDistance: math.MaxInt64,
						shortestFrom:     nil,
						visited:          false,
					}
					moveableNode[position] = node
				}
			}
		}
	}
	if moveableNode[start] == nil {
		return Plan{
			Cost: math.MaxInt64,
			Path: nil,
		}
	}
	moveableNode[start].shortestDistance = 0
	for _, node := range moveableNode {
		heap.Push(unvisitedPQ, node)
	}

	// Shortest path
	var goals []*node
	for unvisitedPQ.Len() > 0 {
		// Pop node with shortest distance
		currentNode := heap.Pop(unvisitedPQ).(*node)
		currentNode.visited = true

		// Test if goal is reached
		if goalTest(currentNode.position) {
			goals = append(goals, currentNode)
			if len(goals) >= goalNum {
				break
			}
		}
		// No path
		if currentNode.shortestDistance == math.MaxInt64 {
			break
		}

		// Expand to next node
		for _, expand := range s.expand(currentNode.position, moveableTest, costModel) {
			neighbor, exist := moveableNode[expand.position]
			if exist && neighbor.visited == false {
				newShortestDistance := currentNode.shortestDistance + expand.distance
				// update shortest distance
				if newShortestDistance < neighbor.shortestDistance {
					neighbor.shortestDistance = newShortestDistance
					neighbor.shortestFrom = currentNode

					// Fix the priority queue
					unvisitedPQ.Update(neighbor)
				}
			}
		}
	}

	if len(goals) == 0 {
		return Plan{
			Cost: math.MaxInt64,
			Path: nil,
		}
	}

	goal := goals[0]
	for _, tmpGoal := range goals {
		if tmpGoal.Less(goal) {
			goal = tmpGoal
		}
	}

	// Get the path by back-tracking
	cost := goal.shortestDistance
	var path []common.Position
	for goal != nil {
		path = append([]common.Position{goal.position}, path...)
		goal = goal.shortestFrom
	}
	return Plan{
		Cost: cost,
		Path: path,
	}
}
