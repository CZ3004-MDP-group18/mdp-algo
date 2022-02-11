package common

import (
	"fmt"
)

type Direction int

const (
	East Direction = iota
	North
	West
	South
)

var AllDirections = [4]Direction{East, North, West, South}

type Move int

const (
	Forward Move = iota
	Backward
	ForwardRight
	ForwardLeft
	BackwardRight
	BackwardLeft
)

var AllMoves = []Move{Forward, Backward, ForwardRight, ForwardLeft, BackwardRight, BackwardLeft}
var AllTurnMoves = []Move{ForwardRight, ForwardLeft, BackwardRight, BackwardLeft}

var MoveName = map[Move]string{
	Forward:       "Forward",
	Backward:      "Backward",
	ForwardRight:  "ForwardRight",
	ForwardLeft:   "ForwardLeft",
	BackwardRight: "BackwardRight",
	BackwardLeft:  "BackwardLeft",
}

// Cell : 2d vector
type Cell struct {
	Xcoord int `json:"x"`
	Ycoord int `json:"y"`
}

// Position : position of the robot (cell, direction)
type Position struct {
	Cell      Cell      `json:"cell"`
	Direction Direction `json:"direction"`
}

// Path : sequence of consecutive positions
type Path []Position

// Transition : sequence of moves
type Transition []Move

func (current Direction) IsVertical() bool {
	if current == North || current == South {
		return true
	}
	return false
}

func (current Direction) Rotate(move Move) Direction {
	next := int(current)
	switch move {
	case ForwardRight, BackwardLeft:
		next = (int(current) + 3) % 4
	case ForwardLeft, BackwardRight:
		next = (int(current) + 1) % 4
	}
	return Direction(next)
}

func (current Direction) Opposite() Direction {
	return Direction((int(current) + 2) % 4)
}

// Transition move to the next position
func (current Position) Transition(transition ...Move) (next Position) {
	next = current
	for _, move := range transition {
		if move == Forward || move == Backward {
			dist := 1
			if move == Backward {
				dist = -1
			}
			next = next.translate(dist)
		} else {
			frontDist := FrontRadius
			sideDist := SideRadius
			if move == BackwardRight || move == BackwardLeft {
				frontDist = -frontDist
				sideDist = -sideDist
			}
			next = next.translate(frontDist)
			next.Direction = next.Direction.Rotate(move)
			next = next.translate(sideDist)
		}
	}

	return
}

// translate move the position according to current's direction
func (current Position) translate(dist int) (next Position) {
	next = current

	switch current.Direction {
	case North:
		next.Cell.Ycoord -= dist
	case South:
		next.Cell.Ycoord += dist
	case East:
		next.Cell.Xcoord += dist
	case West:
		next.Cell.Xcoord -= dist
	}
	return next
}

// Footprint returns the cells affected when make a move
func (current Position) Footprint(move Move) (footprint []Cell) {
	next := current.Transition(move)
	if move == Forward || move == Backward {
		return []Cell{current.Cell, next.Cell}
	}

	// Turn move
	translateMove := Forward
	if move == BackwardRight || move == BackwardLeft {
		translateMove = Backward
	}

	mid := current
	for mid.Cell.Xcoord != next.Cell.Xcoord && mid.Cell.Ycoord != next.Cell.Ycoord {
		mid = mid.Transition(translateMove)
	}
	mid.Direction = next.Direction

	lastTmp := current
	for lastTmp.Cell != mid.Cell {
		footprint = append(footprint, lastTmp.Cell)
		tmp := lastTmp
		for tmp.Cell != mid.Cell {
			tmp = tmp.Transition(translateMove)
			footprint = append(footprint, tmp.Cell)
		}
		mid = mid.Transition(translateMove)
		lastTmp.Direction = next.Direction
		lastTmp = lastTmp.Transition(translateMove)
		lastTmp.Direction = current.Direction
		lastTmp = lastTmp.Transition(translateMove)
	}

	for mid.Cell != next.Cell {
		footprint = append(footprint, mid.Cell)
		mid = mid.Transition(translateMove)
	}
	footprint = append(footprint, mid.Cell)

	return
}

// IsValidMove checks if the move footprint is moveable
func (current Position) IsValidMove(move Move, moveableTest func(Cell) bool) bool {
	// Check if footprint is valid
	for _, cell := range current.Footprint(move) {
		if !moveableTest(cell) {
			return false
		}
	}

	return true
}

// Difference return a move from current to next position
func (current Position) Difference(next Position) (move Move) {
	if current.Direction == next.Direction {
		if current.Transition(Forward) == next {
			return Forward
		} else if current.Transition(Backward) == next {
			return Backward
		} else {
			panic("No possible move")
		}
	}

	// Rotation Move
	for _, turnMove := range AllTurnMoves {
		if current.Transition(turnMove) == next {
			return turnMove
		}
	}
	panic("No possible move")
}

func (t Transition) ToStringArray() (arr []string) {
	for _, move := range t {
		arr = append(arr, MoveName[move])
	}
	return
}

func (t Transition) Log() {
	for _, move := range t {
		fmt.Println(MoveName[move])
	}
}

func (path Path) ToTransition() (transition Transition) {
	for i := 0; i < len(path)-1; i++ {
		current := path[i]
		next := path[i+1]
		move := current.Difference(next)
		transition = append(transition, move)
	}
	return
}

func (path Path) Log() {
	for _, pos := range path {
		fmt.Printf("(%v, %v, %v)\n", pos.Cell.Xcoord, pos.Cell.Ycoord, pos.Direction)
	}
}
