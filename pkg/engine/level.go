package engine

import (
	"github.com/gdamore/tcell"
	"github.com/yoowhi/yarg/pkg/gen"
	"github.com/yoowhi/yarg/pkg/h"
)

type Level struct {
	Size       h.Vector
	Collisions h.Field[bool]
	Visuals    h.Field[Cell]
	Actors     Pool[Actor]
	Rooms      []gen.Area
	RoomsMap   h.Field[int]
	DepthMap   h.Field[int]
	MaxDepth   int
}

func (lvl *Level) IsCollision(coord h.Vector) bool {
	return lvl.Collisions.Get(coord)
}

func GenLevel(size h.Vector) *Level {
	collisions := gen.GenCollisions(size)
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

func genCells(collisions h.Field[bool], depthMap h.Field[int], maxDepth int) h.Field[Cell] {
	wallStyle := tcell.StyleDefault.Background((tcell.ColorDarkGray)).Foreground(tcell.ColorWhite)
	floorStyle := tcell.StyleDefault.Background((tcell.ColorBlack)).Foreground(tcell.ColorDarkGray)
	cells := h.EmptyField[Cell](collisions.Size)

	f := func(coord h.Vector, field *h.Field[Cell]) {
		if collisions.Get(coord) {
			field.Set(coord, Cell{
				Style:  wallStyle,
				Symbol: ' ',
			})
		} else {
			color := getColor(depthMap.Get(coord), maxDepth)
			field.Set(coord, Cell{
				Style:  floorStyle.Background(color),
				Symbol: ' ',
			})
		}
	}

	return cells.ForEach(f)
}

func getColor(depth int, maxDepth int) tcell.Color {
	depth = maxDepth - depth
	coeff := float64(depth) / float64(maxDepth)
	coeff *= coeff * coeff * coeff
	coeff += 1
	coeff /= 2
	return tcell.NewRGBColor(int32(46*coeff), int32(41*coeff), int32(59*coeff))
}
