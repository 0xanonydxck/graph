package graph

import "github.com/zyedidia/generic/queue"

func TopoSort(graph map[int][]int) []int {
	findInDegree := func(graph map[int][]int) map[int]int {
		inDegree := make(map[int]int)
		for key := range graph {
			inDegree[key] = 0
		}

		for _, node := range graph {
			for _, neighbor := range node {
				inDegree[neighbor] += 1
			}
		}

		return inDegree
	}

	topoSort := func(graph map[int][]int) []int {
		result := make([]int, 0)
		q := queue.New[int]()
		inDegree := findInDegree(graph)

		// enqueue 0 in-degree node
		for key, value := range inDegree {
			if value == 0 {
				q.Enqueue(key)
			}
		}

		for !q.Empty() {
			current := q.Dequeue()
			result = append(result, current)

			for _, neighbor := range graph[current] {
				inDegree[neighbor] -= 1
				if inDegree[neighbor] == 0 {
					q.Enqueue(neighbor)
				}
			}
		}

		if len(result) <= len(graph) {
			return nil
		}

		return result
	}

	return topoSort(graph)
}
