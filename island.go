package graph

import (
	"github.com/zyedidia/generic/queue"
)

func NumIsland(grid [][]int) int {
	deltaRows, deltaCols := []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}
	getNeighbors := func(pos Coordinate) []Coordinate {
		result := make([]Coordinate, 0)
		for i := range 4 {
			subRow, subCol := pos[0]-deltaRows[i], pos[1]-deltaCols[i]
			if subRow >= 0 && subCol >= 0 && subRow < len(grid) && subCol < len(grid[0]) {
				result = append(result, Coordinate{subRow, subCol})
			}
		}
		return result
	}

	fill := func(start Coordinate) {
		q := queue.New[Coordinate]()
		q.Enqueue(start)

		for !q.Empty() {
			newQ := queue.New[Coordinate]()
			for !q.Empty() {
				curPos := q.Dequeue()
				if grid[curPos[0]][curPos[1]] == 1 {
					grid[curPos[0]][curPos[1]] = 0
				}

				for _, pos := range getNeighbors(curPos) {
					if grid[pos[0]][pos[1]] == 0 {
						continue
					}

					newQ.Enqueue(pos)
				}
			}

			q = newQ
		}
	}

	countIsland := func() int {
		count := 0
		for row := range len(grid) {
			for col := range len(grid[0]) {
				if grid[row][col] == 0 {
					continue
				}

				fill(Coordinate{row, col})
				count += 1
			}
		}
		return count
	}

	return countIsland()
}
