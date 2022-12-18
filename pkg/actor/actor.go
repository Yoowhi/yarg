package actor

import "github.com/yoowhi/yarg/pkg/h"

type IActor interface {
	GetHealth() int
	SetHealth(value int)

	GetBreed() Breed
	SetBreed(value Breed)

	GetPosition() h.Vector
	SetPosition(value h.Vector)

	TakeTurn()
}

type Breed struct {
	MaxHealth int
	Name      string
}
