package graph

import (
	"reflect"

	"github.com/zyedidia/generic/queue"
)

func Reconstructed(original []int, seqs [][]int) bool {
	buildGraph := func(original []int, seqs [][]int) map[int][]int {
		result := make(map[int][]int)
		for _, key := range original {
			result[key] = make([]int, 0)
		}

		for _, seq := range seqs {
			result[seq[0]] = append(result[seq[0]], seq[1])
		}

		return result
	}

	findInDegree := func(graph map[int][]int) map[int]int {
		inDegree := make(map[int]int)
		for key := range graph {
			inDegree[key] = 0
		}

		for _, neighbors := range graph {
			for _, neighbor := range neighbors {
				inDegree[neighbor]++
			}
		}

		return inDegree
	}

	findQueueLen := func(q *queue.Queue[int]) int {
		count := 0
		q.Each(func(t int) {
			count++
		})
		return count
	}

	reconstructed := func(original []int, seqs [][]int) bool {
		graph := buildGraph(original, seqs)
		inDegree := findInDegree(graph)
		q := queue.New[int]()
		result := make([]int, 0)

		for key, value := range inDegree {
			if value == 0 {
				q.Enqueue(key)
			}
		}

		for !q.Empty() {
			size := findQueueLen(q)
			if size > 1 {
				return false
			}

			current := q.Dequeue()
			result = append(result, current)

			for _, neighbor := range graph[current] {
				inDegree[neighbor]--
				if inDegree[neighbor] == 0 {
					q.Enqueue(neighbor)
				}
			}
		}

		return reflect.DeepEqual(result, original)
	}

	return reconstructed(original, seqs)
}
