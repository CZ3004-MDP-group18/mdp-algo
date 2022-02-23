package session

//var distanceMoved int

/*func (s *sessionImpl) FastestPathInti(
	dist int,
) FastestPlan {
	fmt.Println("Going towards obstacle")
	var d []int
	var movement []string

	d = append(d, dist/10-2)
	d = append(d, 1)

	movement = append(movement, common.MoveName[common.Forward])
	movement = append(movement, common.MoveName[common.ForwardRight])

	return FastestPlan{
		Distance: d,
		Moves:    movement,
		Action:   nil,
	}
}

func (s *sessionImpl) CheckIR(IRdist int, front bool) FastestPlan {
	fmt.Println("Turning araound the obstacle")
	fmt.Println(distanceMoved)

	var step = 30
	var d []int
	var movement []string
	if front {
		if IRdist < 15 {
			fmt.Println("Cannot turn")
			distanceMoved += 30
			d = append(d, step/10)
			movement = append(movement, common.MoveName[common.Forward])
		} else {
			d = append(d, 0)
			movement = append(movement, uTurn("Left")...)
			if distanceMoved > 40 {
				distanceMoved -= 40
				d = append(d, distanceMoved/10)
				movement = append(movement, common.MoveName[common.Forward])
			}
		}
		return FastestPlan{
			Distance: d,
			Moves:    movement,
			Action:   nil,
		}
	} else {
		if IRdist < 55 {
			fmt.Println("Cannot turn")
			distanceMoved += 30
			d = append(d, step/10)
			movement = append(movement, common.MoveName[common.Forward])
		} else {
			d = append(d, 0)
			movement = append(movement, uTurn("Left")...)
			if distanceMoved > 40 {
				d = append(d, distanceMoved/10)
				movement = append(movement, common.MoveName[common.Forward])
			}
		}
		return FastestPlan{
			Distance: d,
			Moves:    movement,
			Action:   nil,
		}
	}
}

func (s *sessionImpl) FastestPathEnd() FastestPlan {
	fmt.Println("Going back to carpark")
	var d []int
	var movement []string

	//Turn back to face carpark
	d = append(d, 0)
	movement = append(movement, common.MoveName[common.ForwardRight])

	// move forward
	d = append(d, 0)
	movement = append(movement, common.MoveName[common.Forward])

	return FastestPlan{
		Distance: d,
		Moves:    movement,
		Action:   nil,
	}
}

func uTurn(dir string) []string {
	var movement []string
	if dir == "Left" {
		movement = append(movement, common.MoveName[common.ForwardLeft])
		movement = append(movement, common.MoveName[common.ForwardLeft])
	} else {
		movement = append(movement, common.MoveName[common.ForwardRight])
		movement = append(movement, common.MoveName[common.ForwardRight])
	}
	return movement
}*/
