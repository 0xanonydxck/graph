package graph

import "testing"

func TestTaskScheduling2(t *testing.T) {
	t.Log(TaskScheduling2([]string{"a", "b", "c", "d"}, []int{1, 1, 2, 1}, []Requirement{{"a", "b"}, {"c", "b"}, {"b", "d"}}))
}
