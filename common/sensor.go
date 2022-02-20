package common

type SensorPayload struct {
	FrontDist int `json:"front"`
	LeftDist  int `json:"left"`
}

type Sensor struct {
	FrontObstacle Transition
	LeftObstacle  Transition
}

func (p SensorPayload) Parse() Sensor {
	sensor := Sensor{
		FrontObstacle: nil,
		LeftObstacle:  nil,
	}

	// Front
	if 0 < p.FrontDist && p.FrontDist <= FrontSensorRange {
		// Need +1 because the front camera is at the front of the robot, but
		// calculations are done using robot center
		for i := 0; i < p.FrontDist+1; i++ {
			sensor.FrontObstacle = append(sensor.FrontObstacle, Forward)
		}
	}

	// Left
	if 0 < p.LeftDist && p.LeftDist <= LeftSensorRange {
		sensor.LeftObstacle = append(sensor.LeftObstacle, LeftRotation)
		// Same reason as above why we need to +1
		for i := 0; i < p.LeftDist+1; i++ {
			sensor.LeftObstacle = append(sensor.LeftObstacle, Forward)
		}
	}

	return sensor
}

// GetObstacleCells get the cells of front and left obstacle. If return value
// is nil, then can assume cannot detect
func (s Sensor) GetObstacleCells(current Position) (front *Cell, left *Cell) {
	if s.FrontObstacle != nil {
		next := current.Transition(s.FrontObstacle...)
		front = &next.Cell
	}
	if s.LeftObstacle != nil {
		next := current.Transition(s.LeftObstacle...)
		left = &next.Cell
	}
	return
}
