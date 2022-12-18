package actor

import "github.com/yoowhi/yarg/pkg/h"

type MeleeCharacter struct {
	Health   int
	Breed    Breed
	Position h.Vector
}

func (c *MeleeCharacter) GetHealth() int {
	return c.Health
}

func (c *MeleeCharacter) SetHealth(value int) {
	c.Health = value
}

func (c *MeleeCharacter) GetBreed() Breed {
	return c.Breed
}

func (c *MeleeCharacter) SetBreed(value Breed) {
	c.Breed = value
}

func (c *MeleeCharacter) GetPosition() h.Vector {
	return c.Position
}

func (c *MeleeCharacter) SetPosition(value h.Vector) {
	c.Position = value
}

func (c *MeleeCharacter) TakeTurn() {

}
