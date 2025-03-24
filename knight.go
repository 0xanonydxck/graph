package graph

import "github.com/zyedidia/generic/queue"

func MinKnightMove(x, y int) int {
	deltaRows := []int{-2, -2, -1, 1, 2, 2, 1, -1}
	deltaCols := []int{-1, 1, 2, 2, 1, -1, -2, -2}

	getNeighbors := func(current Coordinate) []Coordinate {
		result := make([]Coordinate, 0)
		for i := range 8 {
			result = append(result, Coordinate{current[0] + deltaRows[i], current[1] + deltaCols[i]})
		}
		return result
	}

	countMinMove := func(target Coordinate) int {
		start := Coordinate{0, 0}
		visited := map[Coordinate]bool{start: true}
		q := queue.New[Coordinate]()
		q.Enqueue(start)
		count := 0

		for !q.Empty() {
			newQ := queue.New[Coordinate]()
			for !q.Empty() {
				current := q.Dequeue()
				if current[0] == target[0] && current[1] == target[1] {
					return count
				}

				for _, move := range getNeighbors(current) {
					if !visited[move] {
						visited[move] = true
						newQ.Enqueue(move)
					}
				}
			}

			q = newQ
			count += 1
		}

		return count
	}

	return countMinMove(Coordinate{x, y})
}
