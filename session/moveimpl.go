package session

import "mdp_algo/common"

var stepbystep = 0

//var finalObstacleDist = 0
var forwardDist = 0

//var returnDist = 0
var startPosition common.Position
var currentPosition common.Position
var positionToReturn = common.Position{
	Cell: common.Cell{
		Xcoord: startPosition.Cell.Xcoord - 1,
		Ycoord: startPosition.Cell.Ycoord,
	},
	Direction: common.East,
}
var turnNumber int

var fastestmoveCost = map[common.Move]int{
	common.Forward:       1,
	common.Backward:      100,
	common.ForwardLeft:   1,
	common.ForwardRight:  100,
	common.BackwardLeft:  100,
	common.BackwardRight: 100,
}

var fastestCostModel = costModel{
	moveCost: fastestmoveCost,
}
var completePath = false

func (s *sessionImpl) MakeMove(sensor common.SensorPayload) (transition common.Transition) {
	// TODO(Zhi Ying): Implement this
	currentPosition = s.current
	startPosition = s.start

	moveableTest := func(cell common.Cell) bool {
		cs := s.getAugmentCellState(cell)
		return cs.state == empty
	}

	goalTest := func(position common.Position) bool {
		return position == positionToReturn
	}

	if sensor.FrontDist != 0 {
		forwardDist = sensor.FrontDist
		for i := 0; i < (forwardDist - 2); i++ {
			transition = append(transition, common.Forward)
		}
		transition = append(transition, common.ForwardRight)
		turnNumber += 1
	} else if sensor.LeftDist < 2 && turnNumber == 1 {
		transition = append(transition, common.Forward)
		stepbystep += 1
	} else {
		if stepbystep == 0 {
			for x := currentPosition.Cell.Xcoord; x > (common.ObstacleDistance + common.ObstacleBuffer); x-- {
				s.virtualArena[currentPosition.Cell.Ycoord-1][x].state = obstacle
			}
			s.augment()

			transition = s.shortestPath(currentPosition, goalTest, 1, moveableTest, fastestCostModel).Path.ToTransition()
			completePath = true

		} else {
			for x := currentPosition.Cell.Xcoord; x > (common.ObstacleDistance); x-- {
				s.virtualArena[currentPosition.Cell.Ycoord-1][x].state = obstacle
			}
			s.augment()

			transition = s.shortestPath(currentPosition, goalTest, 1, moveableTest, fastestCostModel).Path.ToTransition()
			completePath = true

		}
	}

	if completePath == true {
		transition = append(transition, common.ForwardRight)
		for i := 0; i < forwardDist-4; i++ {
			transition = append(transition, common.Forward)
		}
	}
	return
}

/*func (s *sessionImpl) MakeMove(sensor common.SensorPayload) (transition common.Transition) {
	//
	currentPosition = s.current

	if sensor.FrontDist != 0 {
		forwardDist = sensor.FrontDist
		for i := 0; i < (forwardDist - 3); i++ {
			transition = append(transition, common.Forward)
		}
		transition = append(transition, common.ForwardRight)
		turnNumber += 1
	} else if sensor.LeftDist < 2 && turnNumber == 1 {
		transition = append(transition, common.Forward)
		stepbystep += 1
	} else {
		if stepbystep == 0 {

			finalObstacleDist = common.ObstacleDistance + common.ObstacleBuffer
			returnDist = finalObstacleDist - common.RemoveBuffer
			return getFinalSprint(finalObstacleDist, returnDist)
		} else {

			finalObstacleDist = common.ObstacleDistance - common.RemoveBuffer
			returnDist = stepbystep + 2
			return getFinalSprint(finalObstacleDist, returnDist)
		}

	}

	return
}*/

/*func getFinalSprint(obDist int, returnDist int) (transition common.Transition) {
	// Uturn
	transition = append(transition, uTurn("Left")...) // add a uTurn move

	// At back of the obstacle
	for i := 0; i < (obDist - 3); i++ {
		transition = append(transition, common.Forward)
	}

	// Uturn
	transition = append(transition, uTurn("Left")...) // add a uTurn move

	// At Front of obstacle
	for i := 0; i < returnDist; i++ {
		if stepbystep == 0 {
			transition = append(transition, common.Forward)
		} else {
			transition = append(transition, common.Backward)
		}
	}

	//Turn to face parking lot
	transition = append(transition, common.ForwardRight)

	//Going back to carpark
	for i := 0; i < (forwardDist - 4); i++ {
		transition = append(transition, common.Forward)
	}

	return
}*/

/*func uTurn(dir string) (moves []common.Move) {
	if dir == "Left" {
		moves = append(moves, common.ForwardLeft)
		moves = append(moves, common.ForwardLeft)
	} else {
		moves = append(moves, common.ForwardRight)
		moves = append(moves, common.ForwardRight)
	}
	return
}*/

func (s *sessionImpl) Move(move common.Move) {
	s.current = s.current.Transition(move)
}
