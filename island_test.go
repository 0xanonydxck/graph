package graph

import (
	"fmt"
	"testing"
)

func TestNumIsland(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0},
	}

	fmt.Println(NumIsland(grid))
}
