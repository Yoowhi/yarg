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
	Rooms      []gen.Room
	RoomsMap   [][]int
}

func (lvl *Level) IsCollision(coord h.Vector) bool {
	return lvl.Collisions[coord.X][coord.Y]
}

func GenLevel(size h.Vector) *Level {
	collisions := genCollisions(size)
	rooms, roomsMap := gen.FindRooms(collisions)

	lvl := Level{
		Collisions: collisions,
		Visuals:    genCells(collisions),
		Actors:     Pool[Actor]{},
		Rooms:      rooms,
		RoomsMap:   roomsMap,
	}

	return &lvl
}

func genCollisions(size h.Vector) [][]bool {
	rand := gen.GenRandom(size, 43)
	smoothed := gen.Smooth(rand, 4)
	for i := 0; i < 4; i++ {
		smoothed = gen.Smooth(smoothed, 4)
	}
	smoothed = gen.SmoothRoughness(smoothed)
	return smoothed
}

func genCells(collisions [][]bool) [][]Cell {
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
