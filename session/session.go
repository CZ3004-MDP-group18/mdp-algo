package session

import "mdp_algo/common"

type Session interface {
	HamiltonPath() Plan
	FastestPath() (cost int, path common.Path)

	LoadArena(arena [][]string)
	Reset()
}

type Plan struct {
	cost int
	path common.Path
}
