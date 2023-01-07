package main

/*

 */

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
	"github.com/yoowhi/yarg/pkg/engine"
	"github.com/yoowhi/yarg/pkg/h"
)

var hero Hero
var currentLvl *engine.Level
var screen tcell.Screen

func main() {

	screen = initScreen()
	width, height := screen.Size()

	currentLvl = engine.GenLevel(h.Vector{X: width, Y: height})

	style := tcell.StyleDefault.Background((tcell.ColorBlack)).Foreground(tcell.ColorPurple)
	hero = Hero{
		Health:    10,
		Position:  h.Vector{X: 15, Y: 15},
		MaxHealth: 10,
		Cell:      engine.Cell{Symbol: '@', Style: style},
	}

	currentLvl.Actors.Add(&hero)

	for {
		engine.Draw(screen, currentLvl)
		WaitInput()
		for _, actor := range currentLvl.Actors.List {
			actor.TakeTurn()
		}
		//temp
		switch KeyPressed.Key() {
		case tcell.KeyEsc:
			screen.Fini()
			os.Exit(0)
		case tcell.KeyEnter:
			currentLvl = engine.GenLevel(h.Vector{X: width, Y: height})
			style := tcell.StyleDefault.Background((tcell.ColorBlack)).Foreground(tcell.ColorPurple)
			hero = Hero{
				Health:    10,
				Position:  h.Vector{X: 15, Y: 15},
				MaxHealth: 10,
				Cell:      engine.Cell{Symbol: '@', Style: style},
			}
			currentLvl.Actors.Add(&hero)
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
