package graph

import (
	"math"
	"testing"
)

const Inf = math.MaxInt

func TestGateDistance(t *testing.T) {
	dungeonMap := [][]int{
		{Inf, -1, 0, Inf},
		{Inf, Inf, Inf, -1},
		{Inf, -1, Inf, -1},
		{0, -1, Inf, Inf},
	}

	t.Log(GateDistance(dungeonMap))
}
