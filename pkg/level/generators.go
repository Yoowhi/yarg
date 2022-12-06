package level

import (
	"math/rand"

	"github.com/yoowhi/yarg/pkg/helpers"
)

func GenMap(size helpers.Vector) [][]int {
	lvl := genRandom(size, 43)
	// smoothed := smoothAutomata(lvl, 3)
	// for i := 0; i < 5; i++ {
	// 	smoothed = smoothAutomata(smoothed, 3)
	// }
	smoothed := smoothAutomata(lvl, 4)
	for i := 0; i < 4; i++ {
		smoothed = smoothAutomata(smoothed, 4)
	}
	return smoothed
}

func genRandom(size helpers.Vector, percentFill int) [][]int {
	arr := make([][]int, size.X)
	for x := 0; x < size.X; x++ {
		subarr := make([]int, size.Y)
		for y := 0; y < size.Y; y++ {
			if rand.Intn(101) <= percentFill {
				subarr[y] = 1
			} else {
				subarr[y] = 0
			}
		}
		arr[x] = subarr
	}
	return arr
}

func genEmpty(size helpers.Vector) [][]int {
	arr := make([][]int, size.X)
	for x := 0; x < size.X; x++ {
		subarr := make([]int, size.Y)
		for y := 0; y < size.Y; y++ {
			subarr[y] = 0
		}
		arr[x] = subarr
	}
	return arr
}

func smoothAutomata(lvl [][]int, minNeighbors int) [][]int {
	smoothed := genEmpty(helpers.Vector{X: len(lvl), Y: len(lvl[0])})
	for x := range lvl {
		for y := range lvl[x] {
			neighbors, isEdge := countNeighbors(lvl, x, y)
			if isEdge {
				smoothed[x][y] = 0
			} else if neighbors >= minNeighbors {
				smoothed[x][y] = 1
			} else {
				smoothed[x][y] = 0
			}
		}
	}
	return smoothed
}

func countNeighbors(lvl [][]int, x, y int) (int, bool) {
	lastX := len(lvl) - 1
	lastY := len(lvl[0]) - 1
	neighbors := getNeighborCoords(x, y)
	counter := 0
	for _, neighbor := range neighbors {
		if neighbor.X < 0 || neighbor.X > lastX || neighbor.Y < 0 || neighbor.Y > lastY {
			return counter, true
		}
		counter += lvl[neighbor.X][neighbor.Y]
	}
	return counter, false
}

func getNeighborCoords(x, y int) [8]helpers.Vector {
	arr := [8]helpers.Vector{}

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
