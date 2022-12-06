package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
	"github.com/yoowhi/yarg/pkg/helpers"
	"github.com/yoowhi/yarg/pkg/level"
	"github.com/yoowhi/yarg/pkg/render"
)

func main() {

	screen := initScreen()
	width, height := screen.Size()

	lvl := level.GenMap(helpers.Vector{X: width, Y: height})
	cells := genCells(lvl)

	for {
		render.Draw(screen, cells)
		ev := screen.PollEvent()
		//temp switch
		switch ev := ev.(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEsc {
				screen.Fini()
				os.Exit(0)
			}
			if ev.Key() == tcell.KeyEnter {
				width, height = screen.Size()
				lvl = level.GenMap(helpers.Vector{X: width, Y: height})
				cells = genCells(lvl)
			}
		}
	}

}

func initScreen() tcell.Screen {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	defStyle := tcell.StyleDefault.Background((tcell.ColorBlack)).Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)
	return screen
}

// temp func
func genCells(lvl [][]int) [][]helpers.Cell {
	wallStyle := tcell.StyleDefault.Background((tcell.ColorDarkSlateGrey)).Foreground(tcell.ColorWhite)
	floorStyle := tcell.StyleDefault.Background((tcell.ColorBlack)).Foreground(tcell.ColorWhite)
	cells := make([][]helpers.Cell, len(lvl))
	for x := 0; x < len(lvl); x++ {
		subarr := make([]helpers.Cell, len(lvl[x]))
		for y := 0; y < len(lvl[x]); y++ {
			subarr[y].Symbol = ' '
			if lvl[x][y] == 1 {
				subarr[y].Style = floorStyle
			} else {
				subarr[y].Style = wallStyle
			}

		}
		cells[x] = subarr
	}

	return cells
}
