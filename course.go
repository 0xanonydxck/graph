package graph

const (
	ToVisit = iota
	Visiting
	Visited
)

func ValidCourse(n int, prerequires [][]int) bool {
	buildGraph := func() map[int][]int {
		graph := make(map[int][]int)
		for _, preq := range prerequires {
			if _, ok := graph[preq[0]]; !ok {
				graph[preq[0]] = make([]int, 0)
			}
			graph[preq[0]] = append(graph[preq[0]], preq[1])
		}
		return graph
	}

	graph := buildGraph()
	state := make([]int, n)
	for i := range state {
		state[i] = ToVisit
	}

	var dfs func(start int) bool
	dfs = func(start int) bool {
		state[start] = Visiting
		for _, node := range graph[start] {
			if state[node] == Visited {
				continue
			}

			if state[node] == Visiting {
				return false
			}

			if !dfs(node) {
				return false
			}
		}

		state[start] = Visited
		return true
	}

	for i := range n {
		if dfs(i) {
			continue
		}
		return false
	}

	return true
}
