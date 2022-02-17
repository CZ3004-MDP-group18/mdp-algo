package session

import "mdp_algo/common"

type Session interface {
	HamiltonPath() HamiltonPlan
	FastestPath() Plan

	LoadArena(arena [][]string)
	Reset()
}

type Plan struct {
	Cost int
	Path common.Path
	//Distance []int
	//Moves    []string
	//Action   []string
}

type HamiltonPlan struct {
	Cost               int
	Path               common.Path
	Order              []common.Cell
	DetectImageIndices []int
}
