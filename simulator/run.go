package simulator

import (
	"mdp_algo/controller"
	"mdp_algo/session"
	"sync"
)

var sess = session.NewSession(height, width, startPosition)
var contr = controller.NewController()

// 1 thread running
var mutex sync.Mutex

func setArena(arena [][]string) {
	mutex.Lock()
	defer mutex.Unlock()

	sess.LoadArena(arena)
}

func getHamiltonPath() session.HamiltonPlan {
	mutex.Lock()
	defer mutex.Unlock()

	return sess.HamiltonPath()
}

func runFastestPath() {
	mutex.Lock()
	defer mutex.Unlock()

	for {
		sensor := contr.Sensor()
		turns := sess.MakeMove(sensor)
		for _, turnStr := range turns.ToStringArray(true) {
			contr.Move(turnStr)
		}
		for _, turn := range turns {
			sess.Move(turn)
		}
	}
}
