package gen

import "github.com/yoowhi/yarg/pkg/h"

func CountFloorNeighbors(field h.Field[bool], coord h.Vector) (int, bool) {
	lastX := field.Size.X - 1
	lastY := field.Size.Y - 1
	neighbors := GetNeighborCoords(coord)
	isWall := false
	counter := 0
	for _, neighbor := range neighbors {
		if neighbor.X < 0 || neighbor.X > lastX || neighbor.Y < 0 || neighbor.Y > lastY {
			isWall = true
		} else if !field.Get(neighbor) {
			counter += 1
		}
	}
	return counter, isWall
}

func CountCollisionNeighbors(field h.Field[bool], coord h.Vector) (int, bool) {
	lastX := field.Size.X - 1
	lastY := field.Size.Y - 1
	neighbors := GetNeighborCoords(coord)
	isWall := false
	counter := 0
	for _, neighbor := range neighbors {
		if neighbor.X < 0 || neighbor.X > lastX || neighbor.Y < 0 || neighbor.Y > lastY {
			isWall = true
		} else if field.Get(neighbor) {
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
