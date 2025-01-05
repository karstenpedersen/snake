package loader

import (
	"bufio"
	"io"
	"os"

	"github.com/karstenpedersen/snake/core"
)

func LoadBoardFromFile(filePath string) (*core.Board, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	columns := 0
	cells := make([]int, 0)

	reader := bufio.NewReader(file)
	for {
		// Read bytes from file
		data, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}

		// Process data
		oneLine := false
		for _, c := range data {
			cell, skip := scanCell(c)
			if !skip {
				cells = append(cells, cell)
			}

			// Count how wide the first column is
			if !oneLine {
				columns += 1
				if c == '\n' {
					oneLine = true
				}
			}
		}

		// EOF
		if err != nil {
			break
		}
	}

	return core.NewBoardFromCells(cells, columns), nil
}

func scanCell(c byte) (cellType int, skip bool) {
	switch c {
	case 'a':
		return core.CellApple, false
	case 'c':
		return core.CellCherry, false
	case 's':
		return core.CellSpawn, false
	case '#':
		return core.CellWall, false
	case ' ':
		return core.CellEmpty, false
	default:
		return 0, true
	}
}
