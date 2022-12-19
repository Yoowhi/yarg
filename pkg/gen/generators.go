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
	arr := make([][]bool, size.X)
	for x := 0; x < size.X; x++ {
		subarr := make([]bool, size.Y)
		for y := 0; y < size.Y; y++ {
			if rand.Intn(101) <= percentFill {
				subarr[y] = false
			} else {
				subarr[y] = true
			}
		}
		arr[x] = subarr
	}
	return arr
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

func SmoothAutomata(lvl [][]bool, minNeighbors int) [][]bool {
	smoothed := GenEmpty(h.Vector{X: len(lvl), Y: len(lvl[0])})
	for x := range lvl {
		for y := range lvl[x] {
			neighbors, isEdge := CountFloorNeighbors(lvl, x, y)
			if isEdge {
				smoothed[x][y] = true
			} else if neighbors >= minNeighbors {
				smoothed[x][y] = false
			} else {
				smoothed[x][y] = true
			}
		}
	}
	return smoothed
}

func SmoothRoughness(lvl [][]bool) [][]bool {
	maxFloor := 5
	for x := range lvl {
		for y := range lvl[x] {
			neighbors, isEdge := CountFloorNeighbors(lvl, x, y)
			if isEdge {
				lvl[x][y] = true
			} else if neighbors >= maxFloor {
				lvl[x][y] = false
			}
		}
	}
	return lvl
}

func CountFloorNeighbors(lvl [][]bool, x, y int) (int, bool) {
	lastX := len(lvl) - 1
	lastY := len(lvl[0]) - 1
	neighbors := GetNeighborCoords(x, y)
	counter := 0
	for _, neighbor := range neighbors {
		if neighbor.X < 0 || neighbor.X > lastX || neighbor.Y < 0 || neighbor.Y > lastY {
			return counter, true
		}
		if !lvl[neighbor.X][neighbor.Y] {
			counter += 1
		}
	}
	return counter, false
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
