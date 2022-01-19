package common

type Direction int

const (
	East Direction = iota
	North
	West
	South
)

type Move int

const (
	Forward Move = iota
	Backward
	ForwardRight
	ForwardLeft
	BackwardRight
	BackwardLeft
)

// Cell : 2d vector
type Cell struct {
	Xcoord int
	Ycoord int
}

// Position : position of the robot (cell, direction)
type Position struct {
	Cell      Cell
	Direction Direction
}

// Path : sequence of consecutive positions
type Path []Position

// Transition : sequence of moves
type Transition []Move
