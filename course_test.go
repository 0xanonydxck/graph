package graph

import "testing"

func TestValidCourse(t *testing.T) {
	t.Log(ValidCourse(2, [][]int{{0, 1}}))
	t.Log(ValidCourse(2, [][]int{{0, 1}, {1, 0}}))
}
