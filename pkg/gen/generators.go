package gen

import (
	"math/rand"

	"github.com/yoowhi/yarg/pkg/h"
)

func GenCollisions(size h.Vector) [][]bool {
	rand := GenRandom(size, 43)
	smoothed := SmoothAutomata(rand, 4)
	for i := 0; i < 4; i++ {
		smoothed = SmoothAutomata(smoothed, 4)
	}
	smoothed = SmoothRoughness(smoothed)
	return smoothed
}

func GenRandom(size h.Vector, percentFill int) [][]bool {
	arr := GenEmpty(size)

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

func SmoothAutomata(collisions [][]bool, minNeighbors int) [][]bool {
	smoothed := GenEmpty(h.Vector{X: len(collisions), Y: len(collisions[0])})

	f := func(coord h.Vector, arr [][]bool) {
		neighbors, isEdge := CountFloorNeighbors(collisions, coord.X, coord.Y)
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
		neighbors, isEdge := CountFloorNeighbors(collisions, coord.X, coord.Y)
		if isEdge {
			collisions[coord.X][coord.Y] = true
		} else if neighbors >= maxFloor {
			collisions[coord.X][coord.Y] = false
		}
	}

	ForEach(collisions, f)
	return collisions
}

func ForEach[T any](field [][]T, f func(coord h.Vector, field [][]T)) {
	for x, col := range field {
		for y := range col {
			f(h.Vector{X: x, Y: y}, field)
		}
	}
}

func GenEmpty(size h.Vector) [][]bool {
	arr := make([][]bool, size.X)
	for x := 0; x < size.X; x++ {
		subarr := make([]bool, size.Y)
		for y := 0; y < size.Y; y++ {
			subarr[y] = true
		}
		arr[x] = subarr
	}
	return arr
}

func CountFloorNeighbors(lvl [][]bool, x, y int) (int, bool) {
	lastX := len(lvl) - 1
	lastY := len(lvl[0]) - 1
	neighbors := GetNeighborCoords(x, y)
	isWall := false
	counter := 0
	for _, neighbor := range neighbors {
		if neighbor.X < 0 || neighbor.X > lastX || neighbor.Y < 0 || neighbor.Y > lastY {
			isWall = true
		} else if !lvl[neighbor.X][neighbor.Y] {
			counter += 1
		}
	}
	return counter, isWall
}

func GetNeighborCoords(x, y int) [8]h.Vector {
	arr := [8]h.Vector{}

	xs := [3]int{-1, 0, 1}
	ys := [3]int{1, 0, -1}

	i := 0
	for _, xdif := range xs {
		for _, ydif := range ys {
			if xdif == 0 && ydif == 0 {
				continue
			}
			arr[i].X = x + xdif
			arr[i].Y = y + ydif
			i++
		}
	}
	return arr
}
