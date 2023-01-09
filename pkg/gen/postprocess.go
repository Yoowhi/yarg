package gen

import (
	"github.com/yoowhi/yarg/pkg/h"
)

func GetAreas(collisions [][]bool) ([]Area, [][]int) {
	size := Size(collisions)
	areasMap := GenEmpty(size, -1)
	checked := GenEmpty(size, false)
	areas := []Area{}
	areaCounter := 0

	f := func(coord h.Vector, arr [][]bool) {
		if !collisions[coord.X][coord.Y] && !checked[coord.X][coord.Y] {
			room := Area{Id: areaCounter}
			queue := h.Queue[h.Vector]{}
			queue.Enqueue(coord)
			for !queue.IsEmpty() {
				checkCoord, _ := queue.Dequeue()
				if checked[checkCoord.X][checkCoord.Y] {
					continue
				}
				room.NumOfCells++
				areasMap[checkCoord.X][checkCoord.Y] = areaCounter
				for _, neighbor := range GetNeighborCoords(*checkCoord) {
					if !collisions[neighbor.X][neighbor.Y] && !checked[neighbor.X][neighbor.Y] {
						queue.Enqueue(neighbor)
					}
				}
				checked[checkCoord.X][checkCoord.Y] = true
			}
			areas = append(areas, room)
			areaCounter++
		}

	}

	ForEach(collisions, f)
	return areas, areasMap
}

func GetDepthMap(collisions [][]bool) ([][]int, int) {
	collisions = Copy(collisions)
	size := Size(collisions)
	depthMap := GenEmpty(size, 0)
	edges := GetFloorNearCollision(collisions)
	depthCounter := 1
	for len(edges) > 0 {
		for _, coord := range edges {
			depthMap[coord.X][coord.Y] = depthCounter
			collisions[coord.X][coord.Y] = true
		}
		edges = GetFloorNearCollision(collisions)
		depthCounter++
	}

	return depthMap, depthCounter
}

func GetFloorNearCollision(collisions [][]bool) []h.Vector {
	coords := []h.Vector{}

	f := func(coord h.Vector, arr [][]bool) {
		walls, _ := CountCollisionNeighbors(collisions, coord)
		if !collisions[coord.X][coord.Y] && walls > 0 {
			coords = append(coords, coord)
		}
	}

	ForEach(collisions, f)
	return coords
}
