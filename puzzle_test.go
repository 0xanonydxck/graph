package graph

import "testing"

func TestSlidingPuzzle(t *testing.T) {
	t.Log(SlidingPuzzle(Puzzle{{1, 5, 2}, {4, 0, 3}}))
	t.Log(SlidingPuzzle(Puzzle{{4, 1, 3}, {2, 0, 5}}))
}
