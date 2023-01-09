package engine

import (
	"github.com/gdamore/tcell"
)

func Draw(screen tcell.Screen, lvl *Level) {
	screen.Clear()
	for x := 0; x < len(lvl.Visuals); x++ {
		for y := 0; y < len(lvl.Visuals[x]); y++ {
			rn := lvl.Visuals[x][y].Symbol
			// if lvl.RoomsMap[x][y] >= 0 {
			// 	rn = rune(lvl.RoomsMap[x][y] + 48)
			// }
			screen.SetContent(x, y, rn, nil, lvl.Visuals[x][y].Style)
		}
	}
	for _, actor := range lvl.Actors.List {
		pos := actor.GetPosition()
		cell := actor.GetCell()
		screen.SetContent(pos.X, pos.Y, cell.Symbol, nil, cell.Style.Background(tcell.Color(lvl.Visuals[pos.X][pos.Y].Style)))
	}
	screen.Show()
}
