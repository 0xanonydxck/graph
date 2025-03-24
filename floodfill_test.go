package graph

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	image := [][]int{
		{0, 1, 3, 4, 1},
		{3, 8, 8, 3, 3},
		{6, 7, 8, 8, 3},
		{12, 2, 8, 9, 1},
		{12, 3, 1, 3, 2},
	}
	row, col := 2, 2
	replacement := 9

	expected := [][]int{
		{0, 1, 3, 4, 1},
		{3, 9, 9, 3, 3},
		{6, 7, 9, 9, 3},
		{12, 2, 9, 9, 1},
		{12, 3, 1, 3, 2},
	}

	result := FloodFill(row, col, replacement, image)
	t.Log("IMAGE    :", image)
	t.Log("EXPECTED :", expected)
	t.Log("RESULT   :", result)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	} else {
		t.Log("PASS!!!!")
	}
}
