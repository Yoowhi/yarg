package main

import "github.com/gdamore/tcell"

var KeyPressed *tcell.EventKey

func WaitInput() {
	KeyPressed = nil
	for KeyPressed == nil {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			KeyPressed = ev
		}
	}
}
