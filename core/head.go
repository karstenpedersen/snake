package core

type Head struct {
	X           int
	Y           int
	Length      int
	StartLength int
	Direction   DirectionType
	Lives       int
}

func NewHead(lives, length, x, y int) *Head {
	return &Head{
		X:           x,
		Y:           y,
		Length:      length,
		StartLength: length,
		Direction:   DirectionRight,
		Lives:       lives,
	}
}

func (h *Head) GetStepPosition() (x int, y int) {
	dx, dy := DirectionToPosition(h.Direction)
	return h.X + dx, h.Y + dy
}

func (h *Head) Step() {
	x, y := h.GetStepPosition()
	h.X = x
	h.Y = y
}

func (h *Head) Score() int {
	return h.Length - h.StartLength
}
