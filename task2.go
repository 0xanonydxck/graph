package graph

import (
	"math"
	"slices"

	"github.com/zyedidia/generic/queue"
)

type Task struct {
	Name string
	Time int
}

func NewTask(name string, time int) *Task {
	return &Task{name, time}
}

type Tasks []Task

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t Tasks) Find(name string) *Task {
	index := slices.IndexFunc(t, func(task Task) bool {
		return task.Name == name
	})

	if index == -1 {
		return nil
	}

	return &(t)[index]
}

func (t Tasks) Len() int {
	return len(t)
}

func TaskScheduling2(tasks []string, times []int, requirements []Requirement) int {
	if len(tasks) != len(times) {
		panic("tasks != times")
	}

	buildTasks := func() Tasks {
		result := make(Tasks, 0)

		for i := range tasks {
			result = append(result, *NewTask(tasks[i], times[i]))
		}

		return result
	}

	buildGraph := func(tasks Tasks) map[Task]Tasks {
		result := make(map[Task]Tasks)

		for _, req := range requirements {
			before := tasks.Find(req.Before())
			after := tasks.Find(req.After())
			result[*before] = append(result[*before], *after)
		}

		return result
	}

	buildInDegree := func(tasks Tasks, graph map[Task]Tasks) map[Task]int {
		result := make(map[Task]int)
		for _, task := range tasks {
			result[task] = 0
		}

		for _, node := range graph {
			for _, neignode := range node {
				result[neignode]++
			}
		}

		return result
	}

	scheduling := func() int {
		tasks := buildTasks()
		graph := buildGraph(tasks)
		inDegree := buildInDegree(tasks, graph)

		result := 0
		q := queue.New[Task]()
		for key, value := range inDegree {
			if value == 0 {
				q.Enqueue(key)
			}
		}

		for !q.Empty() {
			newQ := queue.New[Task]()
			consumeTime := math.MinInt
			for !q.Empty() {
				current := q.Dequeue()
				consumeTime = Max(consumeTime, current.Time)

				for _, neighbor := range graph[current] {
					inDegree[neighbor]--
					if inDegree[neighbor] == 0 {
						newQ.Enqueue(neighbor)
					}
				}
			}

			result += consumeTime
			q = newQ
		}

		return result
	}

	return scheduling()
}
