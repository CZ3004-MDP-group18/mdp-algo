package session

import "mdp_algo/common"

type Session interface {
	HamiltonPath() HamiltonPlan
	FastestPath() Plan
	FastestPathInti(dist int) FastestPlan
	CheckIR(IRdist int, front bool) FastestPlan
	FastestPathEnd() FastestPlan

	// MakeMove make the decision to move
	MakeMove(sensor common.Sensor) (transition common.Transition)
	Move(move common.Move)

	LoadArena(arena [][]string)
	Reset()
}

type Plan struct {
	Cost int
	Path common.Path
}

type HamiltonPlan struct {
	Cost               int
	Path               common.Path
	Order              []common.Cell
	DetectImageIndices []int
}

type FastestPlan struct {
	Distance []int
	Moves    []string
	Action   []string
}
