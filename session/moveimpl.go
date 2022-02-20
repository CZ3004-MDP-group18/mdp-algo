package session

import "mdp_algo/common"

func (s *sessionImpl) MakeMove(sensor common.Sensor) (transition common.Transition) {
	// TODO(Zhi Ying): Implement this
	return
}

func (s *sessionImpl) Move(move common.Move) {
	s.current = s.current.Transition(move)
}
