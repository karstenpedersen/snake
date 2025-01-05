package core

const (
	TITLE   = "Snake"
	VERSION = "0.0.1"
)

const (
	// CellSnake > 0
	CellWall   = 0
	CellEmpty  = -1
	CellApple  = -2
	CellCherry = -3
	CellSpawn  = -4 // TODO: Think about this
)

type DirectionType int

const (
	DirectionNone DirectionType = iota
	DirectionUp
	DirectionLeft
	DirectionDown
	DirectionRight
)
