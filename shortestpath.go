package graph

import (
	"github.com/zyedidia/generic/queue"
)

func ShortestPath(graph [][]int, origin, destination int) int {
	getNeighbors := func(node int) []int {
		if node < len(graph) {
			return graph[node]
		}
		return nil
	}

	bfs := func(origin, destination int) int {
		visited := map[int]bool{origin: true}
		q := queue.New[int]()
		q.Enqueue(origin)
		level := 0

		for !q.Empty() {
			newQ := queue.New[int]()
			for !q.Empty() {
				current := q.Dequeue()
				if current == destination {
					return level
				}

				for _, node := range getNeighbors(current) {
					if !visited[node] {
						visited[node] = true
						newQ.Enqueue(node)
					}
				}
			}

			q = newQ
			level += 1
		}

		return 0
	}

	return bfs(origin, destination)
}
