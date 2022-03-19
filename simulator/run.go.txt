package simulator

import (
	"mdp_algo/session"
	"sync"
)

var sess = session.NewSession(height, width, startPosition)

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
