package graph

import (
	"fmt"

	"github.com/zyedidia/generic/queue"
)

func TraversalBFS(node *Node) {
	visited := map[*Node]bool{}
	q := queue.New[*Node]()
	q.Enqueue(node)

	for !q.Empty() {
		newQ := queue.New[*Node]()
		for !q.Empty() {
			current := q.Dequeue()
			visited[current] = true
			fmt.Println(current.Value)

			for _, node := range current.Neighbors {
				if _, exist := visited[node]; !exist {
					newQ.Enqueue(node)
					visited[node] = true
				}
			}
		}

		q = newQ
	}
}
