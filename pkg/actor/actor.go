package actor

import "github.com/yoowhi/yarg/pkg/h"

type IActor interface {
	GetHealth() int
	SetHealth(value int)

	GetMaxHealth() int
	SetMaxHealth(value int)

	GetPosition() h.Vector
	SetPosition(value h.Vector)

	TakeTurn()
}
