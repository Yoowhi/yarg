package engine

import (
	"github.com/gdamore/tcell"
	"github.com/yoowhi/yarg/pkg/gen"
	"github.com/yoowhi/yarg/pkg/h"
)

type Level struct {
	Size       h.Vector
	Collisions [][]bool
	Visuals    [][]Cell
	Actors     Pool[Actor]
	Rooms      []gen.Area
	RoomsMap   [][]int
	DepthMap   [][]int
	MaxDepth   int
}

func (lvl *Level) IsCollision(coord h.Vector) bool {
	return lvl.Collisions[coord.X][coord.Y]
}

func GenLevel(size h.Vector) *Level {
	collisions := genCollisions(size)
	rooms, roomsMap := gen.GetAreas(collisions)
	depthMap, maxDepth := gen.GetDepthMap(collisions)

	lvl := Level{
		Size:       size,
		Collisions: collisions,
		Visuals:    genCells(collisions, depthMap, maxDepth),
		Actors:     Pool[Actor]{},
		Rooms:      rooms,
		RoomsMap:   roomsMap,
		DepthMap:   depthMap,
		MaxDepth:   maxDepth,
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

func genCells(collisions [][]bool, depthMap [][]int, maxDepth int) [][]Cell {
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
				color := getColor(depthMap[x][y], maxDepth)
				subarr[y].Style = floorStyle.Background(color)
				subarr[y].Symbol = ' '
			}

		}
		cells[x] = subarr
	}

	return cells
}

func getColor(depth int, maxDepth int) tcell.Color {
	depth = maxDepth - depth
	coeff := float64(depth) / float64(maxDepth)
	coeff *= coeff * coeff * coeff
	coeff += 1
	coeff /= 2
	return tcell.NewRGBColor(int32(46*coeff), int32(41*coeff), int32(59*coeff))
}
