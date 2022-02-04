package session

import "mdp_algo/common"

type cellStateEnum string

const (
	empty    cellStateEnum = "O"
	obstacle cellStateEnum = "X"

	// Obstacles with face
	obstacleNorth cellStateEnum = "N"
	obstacleSouth cellStateEnum = "S"
	obstacleEast  cellStateEnum = "E"
	obstacleWest  cellStateEnum = "W"
)

type cellState struct {
	state     cellStateEnum
	hasFace   bool
	direction common.Direction
}

type sessionImpl struct {
	height       int
	width        int
	virtualArena [][]cellState
	augmentArena [][]cellState
	current      common.Position

	// Hamilton path algo
	imageCells    []common.Cell
	takePositions map[common.Cell][]common.Position
	planCache     map[common.Position]map[common.Cell]Plan

	//Fastest Path
	distance  int
	pathCache map[common.Position]map[common.Cell]Plan
}

func NewSession(height int, width int, current common.Position) Session {
	s := &sessionImpl{
		height:  height,
		width:   width,
		current: current,
	}
	s.Reset()
	return s
}

func (s *sessionImpl) Reset() {
	s.virtualArena = make([][]cellState, s.height)
	for row := 0; row < s.height; row++ {
		s.virtualArena[row] = make([]cellState, s.width)
	}
	for Ycoord := 0; Ycoord < s.height; Ycoord++ {
		for Xcoord := 0; Xcoord < s.width; Xcoord++ {
			s.virtualArena[Ycoord][Xcoord] = cellState{state: empty}
		}
	}
}

func (s *sessionImpl) LoadArena(arena [][]string) {
	s.Reset()
	for Ycoord := 0; Ycoord < s.height; Ycoord++ {
		for Xcoord := 0; Xcoord < s.width; Xcoord++ {
			c := cellState{}
			value := arena[Ycoord][Xcoord]
			if value == string(empty) {
				c.state = empty
			} else {
				c.state = obstacle
				if value != string(obstacle) {
					c.hasFace = true
				}
			}

			switch value {
			case string(obstacleNorth):
				c.direction = common.North
			case string(obstacleSouth):
				c.direction = common.South
			case string(obstacleEast):
				c.direction = common.East
			case string(obstacleWest):
				c.direction = common.West
			}
			s.virtualArena[Ycoord][Xcoord] = c
		}
	}

	s.augment()
}

func (s *sessionImpl) withinCell(current common.Cell) bool {
	// outside of the map
	if current.Xcoord < 0 || current.Xcoord >= s.width || current.Ycoord < 0 || current.Ycoord >= s.height {
		return false
	}
	return true
}

// augment abstract the obstacles into 3x3 instead of 1x1 and add wall obstacles so that we can treat the robot as 1x1
func (s *sessionImpl) augment() {
	s.augmentArena = make([][]cellState, s.height)
	for i := 0; i < s.height; i++ {
		s.augmentArena[i] = make([]cellState, s.width)
		copy(s.augmentArena[i], s.virtualArena[i])
	}

	// Put virtual obstacle around real obstacle
	for Ycoord := 0; Ycoord < s.height; Ycoord++ {
		for Xcoord := 0; Xcoord < s.width; Xcoord++ {
			if s.virtualArena[Ycoord][Xcoord].state == obstacle {
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						cell := common.Cell{
							Xcoord: Xcoord + i,
							Ycoord: Ycoord + j,
						}
						if s.withinCell(cell) {
							s.augmentArena[cell.Ycoord][cell.Xcoord].state = obstacle
						}
					}
				}
			}
		}
	}

	// Put virtual obstacle around the arena
	for Ycoord := 0; Ycoord < s.height; Ycoord++ {
		for Xcoord := 0; Xcoord < s.width; Xcoord++ {
			if Xcoord == 0 || Xcoord == s.width-1 || Ycoord == 0 || Ycoord == s.height-1 {
				s.augmentArena[Ycoord][Xcoord].state = obstacle
			}
		}
	}
}

func (s *sessionImpl) getAugmentCellState(cell common.Cell) cellState {
	if s.withinCell(cell) {
		return s.augmentArena[cell.Ycoord][cell.Xcoord]
	}
	return cellState{state: obstacle}
}

func (s *sessionImpl) getVirtualCellState(cell common.Cell) cellState {
	if s.withinCell(cell) {
		return s.virtualArena[cell.Ycoord][cell.Xcoord]
	}
	return cellState{state: obstacle}
}
