package graph

import "testing"

func TestReconstructed(t *testing.T) {
	t.Log(Reconstructed([]int{1, 2, 3}, [][]int{{1, 2}, {1, 3}, {2, 3}}))
	t.Log(Reconstructed([]int{1, 2, 3}, [][]int{{1, 2}, {1, 3}}))
}
