package graph

import "github.com/zyedidia/generic/queue"

type Puzzle [2][3]int

var Moves = []Coordinate{
	{0, 1},
	{1, 0},
	{-1, 0},
	{0, -1},
}

func SlidingPuzzle(grid Puzzle) int {
	rowSize, colSize := 2, 3
	target := Puzzle{{1, 2, 3}, {4, 5, 0}}

	getNeighbors := func(grid Puzzle) []Puzzle {
		var pos Coordinate
		for row := range rowSize {
			for col := range colSize {
				if grid[row][col] == 0 {
					pos = NewCoordinate(row, col)
				}
			}
		}

		result := make([]Puzzle, 0)
		for _, move := range Moves {
			x, y := pos.X()+move.X(), pos.Y()+move.Y()
			if (x >= 0 && x < rowSize) && (y >= 0 && y < colSize) {
				temp := grid
				temp[pos.X()][pos.Y()], temp[x][y] = temp[x][y], temp[pos.X()][pos.Y()]
				result = append(result, temp)
			}
		}

		return result
	}

	solvePuzzle := func(start Puzzle) int {
		visited := map[Puzzle]bool{start: true}
		q := queue.New[Puzzle]()
		q.Enqueue(start)

		level := 0
		for !q.Empty() {
			newQ := queue.New[Puzzle]()
			for !q.Empty() {
				current := q.Dequeue()
				if current == target {
					return level
				}

				for _, moved := range getNeighbors(current) {
					if !visited[moved] {
						newQ.Enqueue(moved)
						visited[moved] = true
					}
				}
			}

			q = newQ
			level++
		}

		return -1
	}

	return solvePuzzle(grid)
}
