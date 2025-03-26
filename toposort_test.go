package graph

import (
	"fmt"
	"testing"
)

func TestTopoSort(t *testing.T) {
	result := TopoSort(map[int][]int{
		4: {2},
		5: {2},
		6: {3},
		2: {1},
		3: {1},
	})

	fmt.Println(result)
}
