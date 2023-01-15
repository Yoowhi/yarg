package engine

import (
	"github.com/gdamore/tcell"
	"github.com/yoowhi/yarg/pkg/h"
)

func Draw(screen tcell.Screen, lvl *Level, viewPoint h.Vector) {
	width, height := screen.Size()
	viewSize := h.Vector{X: width, Y: height}
	mapSize := lvl.Size
	drawFrom, drawTo := getDrawBorders(viewSize, viewPoint)

	screen.Clear()
	for x := drawFrom.X; x < drawTo.X; x++ {
		for y := drawFrom.Y; y < drawTo.Y; y++ {
			insideMap := inBounds(h.Vector{X: x, Y: y}, h.Vector{}, mapSize)
			mapCoord := h.Vector{X: x, Y: y}
			rp := mapCoord.Substract(drawFrom)
			if insideMap {
				rn := lvl.Visuals.Get(mapCoord).Symbol
				screen.SetContent(rp.X, rp.Y, rn, nil, lvl.Visuals.Get(mapCoord).Style)
			} else {
				screen.SetContent(rp.X, rp.Y, ' ', nil, lvl.Visuals.Get(h.Vector{}).Style)
			}
		}
	}
	for _, actor := range lvl.Actors.List {
		pos := actor.GetPosition()
		if inBounds(pos, drawFrom, drawTo) {
			rp := pos.Substract(drawFrom)
			cell := actor.GetCell()
			screen.SetContent(rp.X, rp.Y, cell.Symbol, nil, cell.Style.Background(tcell.Color(lvl.Visuals.Get(pos).Style)))
		}
	}
	screen.Show()
}

func getDrawBorders(size h.Vector, viewPoint h.Vector) (h.Vector, h.Vector) {
	halfSize := size.Divide(h.Vector{X: 2, Y: 2})
	drawFrom := viewPoint.Substract(halfSize)
	drawTo := viewPoint.Add(halfSize)
	return drawFrom, drawTo
}

func inBounds(point, boundFrom, boundTo h.Vector) bool {
	inX := point.X >= boundFrom.X && point.X < boundTo.X
	inY := point.Y >= boundFrom.Y && point.Y < boundTo.Y
	return inX && inY
}
