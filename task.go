package graph

import "github.com/zyedidia/generic/queue"

type Requirement [2]string

func (r *Requirement) Before() string { return r[0] }
func (r *Requirement) After() string  { return r[1] }

func TaskScheduling(tasks []string, requirements []Requirement) []string {
	buildGraph := func(tasks []string, requirements []Requirement) map[string][]string {
		result := make(map[string][]string)
		for _, task := range tasks {
			result[task] = make([]string, 0)
		}

		for _, req := range requirements {
			result[req.Before()] = append(result[req.Before()], req.After())
		}
		return result
	}

	findInDegree := func(graph map[string][]string) map[string]int {
		result := make(map[string]int)
		for key := range graph {
			result[key] = 0
		}

		for _, node := range graph {
			for _, neighbor := range node {
				result[neighbor]++
			}
		}

		return result
	}

	scheduling := func(tasks []string, requrements []Requirement) []string {
		graph := buildGraph(tasks, requrements)
		inDegree := findInDegree(graph)
		result := make([]string, 0)
		q := queue.New[string]()

		for key, value := range inDegree {
			if value == 0 {
				q.Enqueue(key)
			}
		}

		for !q.Empty() {
			current := q.Dequeue()
			result = append(result, current)

			for _, neighbor := range graph[current] {
				inDegree[neighbor]--
				if inDegree[neighbor] == 0 {
					q.Enqueue(neighbor)
				}
			}
		}

		if len(result) != len(graph) {
			return nil
		}

		return result
	}

	return scheduling(tasks, requirements)
}
