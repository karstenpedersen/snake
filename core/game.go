package core

type Game struct {
	Score int
	Heads []*Head
	Board *Board
}

func NewGame(columns, rows int) *Game {
	length := 3
	lives := 3
	cx := columns / 2
	cy := rows / 2

	player := NewHead(lives, length, cx, cy)

	return &Game{
		Heads: []*Head{player},
		Board: NewBoard(columns, rows),
	}
}

func (g *Game) Update() {
	// Update board
	g.Board.Update()

	// Update headers
	for _, head := range g.Heads {
		g.UpdateHead(head)
	}
}

func (g *Game) UpdateHead(head *Head) {
	// Check if head can move
	x, y := head.GetStepPosition()
	cellType, isValid := g.Board.Get(x, y)
	if isValid {
		// Apply item to player
		switch cellType {
		case CellApple:
			head.Length += 1
			g.SpawnItem(CellApple)
		case CellCherry:
			head.Lives += 1
		}

		// Update head position
		head.Step()
		g.Board.Put(head.Length, head.X, head.Y)
	} else {
		head.Lives -= 1
		// TODO: Respawn snake
	}
}

func (g *Game) SpawnItem(item int) {
	x := 1
	y := 1
	g.Board.Put(item, x, y)
	// TODO: Somehow make sure item spawns
}
