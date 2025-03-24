package graph

import "github.com/zyedidia/generic/queue"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GateDistance(grid [][]int) [][]int {
	deltaRows, deltaCols := []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}
	getNeighbors := func(pos Coordinate) []Coordinate {
		result := make([]Coordinate, 0)
		for i := range 4 {
			subRow, subCol := pos[0]-deltaRows[i], pos[1]-deltaCols[i]
			if subRow >= 0 && subCol >= 0 && subRow < len(grid) && subCol < len(grid[0]) {
				result = append(result, NewCoordinate(subRow, subCol))
			}
		}
		return result
	}

	bfs := func(start Coordinate) {
		visited := map[Coordinate]bool{start: true}
		q := queue.New[Coordinate]()
		q.Enqueue(start)

		level := 0
		for !q.Empty() {
			newQ := queue.New[Coordinate]()
			for !q.Empty() {
				current := q.Dequeue()
				if grid[current.X()][current.Y()] > 0 {
					grid[current.X()][current.Y()] = Min(grid[current.X()][current.Y()], level)
				}

				for _, pos := range getNeighbors(current) {
					if grid[pos.X()][pos.Y()] > 0 && !visited[pos] {
						visited[pos] = true
						newQ.Enqueue(pos)
					}
				}
			}

			q = newQ
			level++
		}
	}

	calculateDistance := func() {
		for row := range grid {
			for col := range grid[0] {
				if grid[row][col] == 0 {
					bfs(NewCoordinate(row, col))
				}
			}
		}
	}

	calculateDistance()
	return grid
}
