package graph

import "fmt"

func TraversalDFS(root *Node) {
	visited := make(map[*Node]bool)

	var dfs func(*Node)
	dfs = func(root *Node) {
		if visited[root] {
			return
		}

		fmt.Println(root.Value)
		visited[root] = true
		for _, node := range root.Neighbors {
			if !visited[node] {
				dfs(node)
			}
		}
	}

	dfs(root)
}
