package controller

import "mdp_algo/common"

type Controller interface {
	Move(move string)
	Sensor() common.SensorPayload
}
