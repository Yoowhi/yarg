package main

import (
	"github.com/yoowhi/yarg/pkg/engine"
	"github.com/yoowhi/yarg/pkg/engine/component"
	"github.com/yoowhi/yarg/pkg/h"
)

type Hero struct {
	Health    int
	MaxHealth int
	Position  h.Vector
	Cell      engine.Cell

	Inventory component.Inventory
}

func (c *Hero) GetHealth() int {
	return c.Health
}

func (c *Hero) SetHealth(value int) {
	c.Health = value
}

func (c *Hero) GetMaxHealth() int {
	return c.MaxHealth
}

func (c *Hero) SetMaxHealth(value int) {
	c.MaxHealth = value
}

func (c *Hero) GetPosition() h.Vector {
	return c.Position
}

func (c *Hero) SetPosition(value h.Vector) {
	c.Position = value
}

func (c *Hero) TakeTurn() {

}

func (c *Hero) GetCell() engine.Cell {
	return c.Cell
}
