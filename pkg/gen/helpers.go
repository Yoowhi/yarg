package gen

import "github.com/yoowhi/yarg/pkg/h"

func ForEach[T any](field [][]T, f func(coord h.Vector, field [][]T)) {
	for x, col := range field {
		for y := range col {
			f(h.Vector{X: x, Y: y}, field)
		}
	}
}

func GenEmpty[T any](size h.Vector, initVal T) [][]T {
	arr := make([][]T, size.X)
	for x := 0; x < size.X; x++ {
		subarr := make([]T, size.Y)
		for y := 0; y < size.Y; y++ {
			subarr[y] = initVal
		}
		arr[x] = subarr
	}
	return arr
}

func CountFloorNeighbors(lvl [][]bool, coord h.Vector) (int, bool) {
	lastX := len(lvl) - 1
	lastY := len(lvl[0]) - 1
	neighbors := GetNeighborCoords(coord)
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

func CountCollisionNeighbors(lvl [][]bool, coord h.Vector) (int, bool) {
	lastX := len(lvl) - 1
	lastY := len(lvl[0]) - 1
	neighbors := GetNeighborCoords(coord)
	isWall := false
	counter := 0
	for _, neighbor := range neighbors {
		if neighbor.X < 0 || neighbor.X > lastX || neighbor.Y < 0 || neighbor.Y > lastY {
			isWall = true
		} else if lvl[neighbor.X][neighbor.Y] {
			counter += 1
		}
	}
	return counter, isWall
}

func GetNeighborCoords(coord h.Vector) [8]h.Vector {
	arr := [8]h.Vector{}

	xs := [3]int{-1, 0, 1}
	ys := [3]int{1, 0, -1}

	i := 0
	for _, xdif := range xs {
		for _, ydif := range ys {
			if xdif == 0 && ydif == 0 {
				continue
			}
			arr[i].X = coord.X + xdif
			arr[i].Y = coord.Y + ydif
			i++
		}
	}
	return arr
}

func Size[T any](plane [][]T) h.Vector {
	return h.Vector{
		X: len(plane),
		Y: len(plane[0]),
	}
}

func Copy[T any](plane [][]T) [][]T {
	size := Size(plane)
	duplicate := make([][]T, size.X)
	for x := range plane {
		duplicate[x] = make([]T, size.Y)
		copy(duplicate[x], plane[x])
	}
	return duplicate
}
