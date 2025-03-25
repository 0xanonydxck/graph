package graph

import "github.com/zyedidia/generic/queue"

func WordLadder(start, end string, list []string) int {
	if len(start) != len(end) {
		return -1
	}

	size := len(start)
	getNeighbors := func(word string) []string {
		result := make([]string, 0)
		for _, item := range list {
			count := 0
			for i := range size {
				if item[i] == word[i] {
					count++
				}
			}
			if count == size-1 {
				result = append(result, item)
			}
		}
		return result
	}

	calculate := func(start, end string) int {
		visited := map[string]bool{start: true}
		q := queue.New[string]()
		q.Enqueue(start)
		level := 0

		for !q.Empty() {
			newQ := queue.New[string]()
			for !q.Empty() {
				current := q.Dequeue()
				if current == end {
					return level
				}

				for _, word := range getNeighbors(current) {
					if !visited[word] {
						visited[word] = true
						newQ.Enqueue(word)
					}
				}
			}

			q = newQ
			level++
		}

		return -1
	}

	return calculate(start, end)
}
