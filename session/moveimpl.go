package session

import "mdp_algo/common"

var stepbystep = 0
var finalObstacleDist = 0
var forwardDist = 0
var returnDist = 0

func (s *sessionImpl) MakeMove(sensor common.SensorPayload) (transition common.Transition) {
	// TODO(Zhi Ying): Implement this
	if sensor.FrontDist != 0 {
		forwardDist = sensor.FrontDist
		for i := 0; i < ((forwardDist - 20) / 10); i++ {
			transition = append(transition, common.Forward)
		}
		transition = append(transition, common.ForwardRight)
	} else if sensor.LeftDist < 20 {
		transition = append(transition, common.Forward)
		stepbystep += 1
	} else {
		if stepbystep == 0 {
			finalObstacleDist = common.ObstacleDistance + common.ObstacleBuffer
			returnDist = finalObstacleDist - common.RemoveBuffer
			return getFinalSprint(finalObstacleDist, returnDist)
		} else {
			finalObstacleDist = common.ObstacleDistance - common.RemoveBuffer
			returnDist = (stepbystep + 2) * 10
			return getFinalSprint(finalObstacleDist, returnDist)
		}

	}

	return
}

func getFinalSprint(obDist int, returnDist int) (transition common.Transition) {
	// Uturn
	transition = append(transition, uTurn("Left")...)

	// At back of the obstacle
	for i := 0; i < (obDist-30)/10; i++ {
		transition = append(transition, common.Forward)
	}

	// Uturn
	transition = append(transition, uTurn("Left")...)

	// At Front of obstacle
	for i := 0; i < returnDist/10; i++ {
		if stepbystep == 0 {
			transition = append(transition, common.Forward)
		} else {
			transition = append(transition, common.Backward)
		}
	}

	//Turn to face parking lot
	transition = append(transition, common.ForwardRight)

	//Going back to carpark
	for i := 0; i < (forwardDist-40)/10; i++ {
		transition = append(transition, common.Forward)
	}

	return
}

func uTurn(dir string) (moves []common.Move) {
	if dir == "Left" {
		moves = append(moves, common.ForwardLeft)
		moves = append(moves, common.ForwardLeft)
	} else {
		moves = append(moves, common.ForwardRight)
		moves = append(moves, common.ForwardRight)
	}
	return
}

func (s *sessionImpl) Move(move common.Move) {
	s.current = s.current.Transition(move)
}
