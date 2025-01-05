package core

func DirectionToPosition(direction DirectionType) (x int, y int) {
	x = 0
	y = 0

	switch direction {
	case DirectionUp:
		y -= 1
	case DirectionLeft:
		x -= 1
	case DirectionDown:
		y += 1
	case DirectionRight:
		x += 1
	}

	return x, y
}
