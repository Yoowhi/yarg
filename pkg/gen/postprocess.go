package gen

import (
	"github.com/yoowhi/yarg/pkg/h"
)

func GetAreas(collisions h.Field[bool]) ([]Area, h.Field[int]) {
	size := collisions.Size
	areasMap := h.NewField(size, -1)
	checked := h.NewField(size, false)
	areas := []Area{}
	areaCounter := 0

	f := func(coord h.Vector, field *h.Field[bool]) {
		if !collisions.Get(coord) && !checked.Get(coord) {
			room := Area{Id: areaCounter}
			queue := h.Queue[h.Vector]{}
			queue.Enqueue(coord)
			for !queue.IsEmpty() {
				checkCoord, _ := queue.Dequeue()
				if checked.Get(*checkCoord) {
					continue
				}
				room.NumOfCells++
				areasMap.Set(*checkCoord, areaCounter)
				for _, neighbor := range GetNeighborCoords(*checkCoord) {
					if !collisions.Get(neighbor) && !checked.Get(neighbor) {
						queue.Enqueue(neighbor)
					}
				}
				checked.Set(*checkCoord, true)
			}
			areas = append(areas, room)
			areaCounter++
		}

	}

	collisions.ForEach(f)
	return areas, areasMap
}

func GetDepthMap(collisions h.Field[bool]) (h.Field[int], int) {
	collisions = collisions.Duplicate()
	size := collisions.Size
	depthMap := h.NewField(size, 0)
	edges := GetFloorNearCollision(collisions)
	depthCounter := 1
	for len(edges) > 0 {
		for _, coord := range edges {
			depthMap.Set(coord, depthCounter)
			collisions.Set(coord, true)
		}
		edges = GetFloorNearCollision(collisions)
		depthCounter++
	}

	return depthMap, depthCounter
}

func GetFloorNearCollision(collisions h.Field[bool]) []h.Vector {
	coords := []h.Vector{}

	f := func(coord h.Vector, field *h.Field[bool]) {
		walls, _ := CountCollisionNeighbors(collisions, coord)
		if !collisions.Get(coord) && walls > 0 {
			coords = append(coords, coord)
		}
	}

	collisions.ForEach(f)
	return coords
}
