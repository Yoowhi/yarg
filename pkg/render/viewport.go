package render

import (
	"github.com/gdamore/tcell"
	"github.com/yoowhi/yarg/pkg/level"
)

func Draw(screen tcell.Screen, cells [][]level.Cell) {
	screen.Clear()
	for x := 0; x < len(cells); x++ {
		for y := 0; y < len(cells[x]); y++ {
			screen.SetContent(x, y, cells[x][y].Symbol, nil, cells[x][y].Style)
		}
	}
	screen.Show()
}
