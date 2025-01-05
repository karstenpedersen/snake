package core

import (
	"slices"
)

type Board struct {
	Columns   int
	Rows      int
	Cells     []int
	FreeCells []int
}

func NewBoard(columns, rows int) *Board {
	cells := slices.Repeat([]int{CellEmpty}, columns*rows)
	return NewBoardFromCells(cells, columns)
}

func NewBoardFromCells(cells []int, columns int) *Board {
	rows := len(cells) / columns
	return &Board{
		Columns: columns,
		Rows:    rows,
		Cells:   cells,
	}
}

func (b *Board) IsPositionOutOfBounds(x, y int) bool {
	return x < 0 || x > b.Columns-1 || y < 0 || y > b.Rows-1
}

func (b *Board) IsCellFree(x, y int) bool {
	return x >= 0 && x < b.Columns && y > 0 && y < b.Rows && b.Cells[x*y] < 0
}

func (b *Board) Get(x, y int) (cellType int, isValidPutPosition bool) {
	if b.IsPositionOutOfBounds(x, y) {
		return 0, false
	}

	// Check if cell contains snake or wall
	cell := b.Cells[x*y]
	if cell >= 0 {
		return cell, false
	}

	return cell, true
}

func (b *Board) Put(value, x, y int) (oldCellType int, success bool) {
	if b.IsPositionOutOfBounds(x, y) {
		return 0, false
	}

	// Check if cell contains snake or wall
	cell := b.Cells[x*y]
	if cell >= 0 {
		return cell, false
	}

	b.Cells[x*y] = value
	return cell, true
}

func (b *Board) Update() {
	for i := range len(b.Cells) {
		if b.Cells[i] > 0 {
			b.Cells[i] -= 1
		}
	}
}
