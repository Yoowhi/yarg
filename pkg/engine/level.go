package engine

import (
	"github.com/gdamore/tcell"
	"github.com/yoowhi/yarg/pkg/gen"
	"github.com/yoowhi/yarg/pkg/h"
)

type Level struct {
	Collisions [][]bool
	Visuals    [][]Cell
	Actors     Pool[Actor]
}

func (lvl *Level) IsCollision(coord h.Vector) bool {
	return lvl.Collisions[coord.X][coord.Y]
}

func GenLevel(size h.Vector) Level {
	collisions := gen.GenCollisions(size)

	lvl := Level{
		Collisions: collisions,
		Visuals:    GenCells(collisions),
		Actors:     Pool[Actor]{},
	}

	return lvl
}

func GenCells(collisions [][]bool) [][]Cell {
	wallStyle := tcell.StyleDefault.Background((tcell.ColorDarkGray)).Foreground(tcell.ColorWhite)
	floorStyle := tcell.StyleDefault.Background((tcell.ColorBlack)).Foreground(tcell.ColorDarkGray)
	cells := make([][]Cell, len(collisions))
	for x := 0; x < len(collisions); x++ {
		subarr := make([]Cell, len(collisions[x]))
		for y := 0; y < len(collisions[x]); y++ {
			if collisions[x][y] {
				subarr[y].Style = wallStyle
				subarr[y].Symbol = ' '
			} else {
				subarr[y].Style = floorStyle
				subarr[y].Symbol = '.'
			}

		}
		cells[x] = subarr
	}

	return cells
}
