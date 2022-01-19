package controller

import "mdp_algo/common"

type Controller interface {
	Move(move common.Move)
	Transition(transition common.Transition)
}
