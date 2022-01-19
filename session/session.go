package session

import "mdp_algo/common"

type Session interface {
	HamiltonPath() (cost int, path common.Path)
	FastestPath() (cost int, path common.Path)

	LoadArena(arena [][]string)
}
