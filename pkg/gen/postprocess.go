package gen

import (
	"github.com/yoowhi/yarg/pkg/h"
)

func FindRooms(collisions [][]bool) ([]Room, [][]int) {
	size := Size(collisions)
	roomsMap := GenEmpty(size, -1)
	checked := GenEmpty(size, false)
	rooms := []Room{}
	roomsCounter := 0

	f := func(coord h.Vector, arr [][]bool) {
		if !collisions[coord.X][coord.Y] && !checked[coord.X][coord.Y] {
			room := Room{Id: roomsCounter}
			queue := h.Queue[h.Vector]{}
			queue.Enqueue(coord)
			for !queue.IsEmpty() {
				checkCoord, _ := queue.Dequeue()
				if checked[checkCoord.X][checkCoord.Y] {
					continue
				}
				room.NumOfCells++
				roomsMap[checkCoord.X][checkCoord.Y] = roomsCounter
				for _, neighbor := range GetNeighborCoords(*checkCoord) {
					if !collisions[neighbor.X][neighbor.Y] && !checked[neighbor.X][neighbor.Y] {
						queue.Enqueue(neighbor)
					}
				}
				checked[checkCoord.X][checkCoord.Y] = true
			}
			rooms = append(rooms, room)
			roomsCounter++
		}

	}

	ForEach(collisions, f)
	return rooms, roomsMap
}
