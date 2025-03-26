package graph

import "testing"

func TestTaskScheduling(t *testing.T) {
	t.Log(TaskScheduling([]string{"a", "b", "c", "d"}, []Requirement{{"a", "b"}, {"c", "b"}, {"b", "d"}}))
}
