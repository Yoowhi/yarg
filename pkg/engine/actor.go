package engine

import "github.com/yoowhi/yarg/pkg/h"

type Actor interface {
	GetHealth() int
	SetHealth(value int)

	GetMaxHealth() int
	SetMaxHealth(value int)

	GetPosition() h.Vector
	SetPosition(value h.Vector)

	GetCell() Cell

	TakeTurn()
}
