package main

import (
	"github.com/yoowhi/yarg/pkg/engine/component"
	"github.com/yoowhi/yarg/pkg/h"
)

type MeleeCharacter struct {
	Health    int
	MaxHealth int
	Position  h.Vector

	Inventory component.Inventory
}

func (c *MeleeCharacter) GetHealth() int {
	return c.Health
}

func (c *MeleeCharacter) SetHealth(value int) {
	c.Health = value
}

func (c *MeleeCharacter) GetMaxHealth() int {
	return c.MaxHealth
}

func (c *MeleeCharacter) SetMaxHealth(value int) {
	c.MaxHealth = value
}

func (c *MeleeCharacter) GetPosition() h.Vector {
	return c.Position
}

func (c *MeleeCharacter) SetPosition(value h.Vector) {
	c.Position = value
}

func (c *MeleeCharacter) TakeTurn() {

}
