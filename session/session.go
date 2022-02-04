package session

import "mdp_algo/common"

type Session interface {
	HamiltonPath() Plan
	FastestPath() Plan

	LoadArena(arena [][]string)
	Reset()
}

type Plan struct {
	Cost int
	Path common.Path
}
