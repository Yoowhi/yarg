package gen

import (
	"math/rand"

	"github.com/yoowhi/yarg/pkg/h"
)

func GenCollisions(size h.Vector) h.Field[bool] {
	rand := GenRandom(size, 43)
	smoothed := Smooth(rand, 4)
	for i := 0; i < 4; i++ {
		smoothed = Smooth(smoothed, 4)
	}
	smoothed = SmoothRoughness(smoothed)
	return smoothed
}

func GenRandom(size h.Vector, percentFill int) h.Field[bool] {
	arr := h.NewField(size, true)

	f := func(coord h.Vector, field *h.Field[bool]) {
		if rand.Intn(101) <= percentFill {
			field.Set(coord, false)
		} else {
			field.Set(coord, true)
		}
	}

	return arr.ForEach(f)
}

func Smooth(collisions h.Field[bool], minNeighbors int) h.Field[bool] {
	smoothed := h.NewField(collisions.Size, true)

	f := func(coord h.Vector, field *h.Field[bool]) {
		neighbors, isEdge := CountFloorNeighbors(collisions, coord)
		if isEdge {
			field.Set(coord, true)
		} else if neighbors >= minNeighbors {
			field.Set(coord, false)
		} else {
			field.Set(coord, true)
		}
	}

	return smoothed.ForEach(f)
}

func SmoothRoughness(collisions h.Field[bool]) h.Field[bool] {
	maxFloor := 5

	f := func(coord h.Vector, field *h.Field[bool]) {
		neighbors, isEdge := CountFloorNeighbors(collisions, coord)
		if isEdge {
			field.Set(coord, true)
		} else if neighbors >= maxFloor {
			field.Set(coord, false)
		}
	}

	return collisions.ForEach(f)
}
