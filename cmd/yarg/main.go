package main

/**
Игровой луп должен включать в себя поведение NPC в игровой среде.
Все игровые объекты можно разделить на 3 главные составляющие:
	- игровая карта
	- статичные объекты со взаимодейстием (рычаги, сундуки...)
	- NPC

Первая нужна для отрисовки местности и рассчёта столкновений.
Вторую и третью можно объединить в одну большую структуру GameObject.
GameObject должен включать в себя возможность хранить компоненты
объекта и логику взаимодействия между ними.
Компонент включает в себя свойства определенной игровой фичи. Напр-
имер инвентарь, координату, здоровье, визуальное отображение...
Логика взаимодействия должна иметь доступ к этим свойствам и изменять
их.

Предположим, GameObject будет реализовывать интерфейсы, которые
гарантируют доступ к определенным компонентам GameObject.
Желательно бы иметь возможность добавлять новые компоненты на лету.

Как тогда можно узнать какой объект и какого типа будет находиться
в искомой точке игровой карты? lvl.getFrom(x, y)

////////////// ХУЙНЯ, В ГОЛАНГЕ НЕТ НАСЛЕДИЯ //////////////


Все GameObject можно разделить по категориям:
	- предметы
	- персонажи
	- снаряды
	- сундуки
	- двери
	- рычаги
	- ...
Каждый такой объект будет иметь в себе компоненты. Интерфейсы будут
обеспечивать доступ к наборам компонентов. Это понадобится для того
чтобы производить действия с этими оъектами.
В каждом файле который описывает интерфейс должен находиться слайс,
который хранит ссылки на все инстансы данного интерфейса. Набор функций
для управления этими слайсами (для поиска, создания, удаления) позволит
получать доступ к любому GameObject в нужном для исполнения аспекте.
**/

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
	"github.com/yoowhi/yarg/pkg/engine"
	"github.com/yoowhi/yarg/pkg/h"
)

var hero Hero
var currentLvl engine.Level
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
		engine.Draw(screen, &currentLvl)
		HandleInput()
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
