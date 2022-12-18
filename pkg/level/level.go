package level

import (
	"github.com/yoowhi/yarg/pkg/actor"
	"github.com/yoowhi/yarg/pkg/h"
)

type Level struct {
	Collisions [][]bool
	Visuals    [][]Cell
	Actors     h.Pool[actor.IActor]
}
