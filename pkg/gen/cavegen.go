package gen

import (
	"math/rand"

	"github.com/yoowhi/yarg/pkg/h"
)

func GenRandom(size h.Vector, percentFill int) [][]bool {
	arr := GenEmpty(size, true)

	f := func(coord h.Vector, arr [][]bool) {
		if rand.Intn(101) <= percentFill {
			arr[coord.X][coord.Y] = false
		} else {
			arr[coord.X][coord.Y] = true
		}
	}

	ForEach(arr, f)
	return arr
}

func Smooth(collisions [][]bool, minNeighbors int) [][]bool {
	smoothed := GenEmpty(Size(collisions), true)

	f := func(coord h.Vector, arr [][]bool) {
		neighbors, isEdge := CountFloorNeighbors(collisions, coord)
		if isEdge {
			smoothed[coord.X][coord.Y] = true
		} else if neighbors >= minNeighbors {
			smoothed[coord.X][coord.Y] = false
		} else {
			smoothed[coord.X][coord.Y] = true
		}
	}

	ForEach(collisions, f)
	return smoothed
}

func SmoothRoughness(collisions [][]bool) [][]bool {
	maxFloor := 5

	f := func(coord h.Vector, arr [][]bool) {
		neighbors, isEdge := CountFloorNeighbors(collisions, coord)
		if isEdge {
			collisions[coord.X][coord.Y] = true
		} else if neighbors >= maxFloor {
			collisions[coord.X][coord.Y] = false
		}
	}

	ForEach(collisions, f)
	return collisions
}
