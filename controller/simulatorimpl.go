package controller

import "mdp_algo/common"

type simulatorImpl struct {
	physicalArena [][]bool
	current       common.Position
}

var turns = 0

func NewSimulatorController(arena [][]bool, start common.Position) Controller {
	return &simulatorImpl{
		physicalArena: arena,
		current:       start,
	}
}

func (s *simulatorImpl) Move(move string) {
	println(move)
}

func (s *simulatorImpl) Sensor() common.SensorPayload {
	var payload = common.SensorPayload{}
	var obstacledist int
	for i := 1; i < 20; i++ {
		if s.physicalArena[i][s.current.Cell.Xcoord] == true {
			obstacledist = i
		}
	}

	payload.FrontDist = s.current.Cell.Ycoord - obstacledist
	payload.LeftDist = s.current.Cell.Ycoord - obstacledist

	return payload
}
