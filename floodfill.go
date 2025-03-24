package graph

import (
	"github.com/zyedidia/generic/queue"
)

func FloodFill(row, col, replacement int, image [][]int) [][]int {
	deltaRows, deltaCols := []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}

	getNeighbors := func(row, col int) (positions []Coordinate) {
		for i := range 4 {
			subRow, subCol := row-deltaRows[i], col-deltaCols[i]
			if subRow >= 0 && subCol >= 0 && subRow < len(image) && subCol < len(image[0]) {
				positions = append(positions, Coordinate{subRow, subCol})
			}
		}

		return
	}

	bfs := func(row, col, replacement int, image [][]int) [][]int {
		target := image[row][col]
		visited := map[Coordinate]bool{}     // arrange visited store
		visited[Coordinate{row, col}] = true // assign the selected position as visited
		q := queue.New[Coordinate]()         // arrange queue for store BFS path walk
		q.Enqueue(Coordinate{row, col})      // enqueue the root

		for !q.Empty() {
			newQ := queue.New[Coordinate]()
			for !q.Empty() {
				pos := q.Dequeue() // pos[0] = row, pos[1] = col
				if image[pos[0]][pos[1]] == target {
					image[pos[0]][pos[1]] = replacement
				}

				for _, pos := range getNeighbors(pos[0], pos[1]) {
					if !visited[pos] {
						newQ.Enqueue(pos)
						visited[pos] = true
					}
				}
			}

			q = newQ
		}

		return image
	}
	return bfs(row, col, replacement, image)
}
