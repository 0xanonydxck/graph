package graph

import "testing"

func TestShortestPath(t *testing.T) {
	tests := []struct {
		name        string
		graph       [][]int
		origin      int
		destination int
		want        int
	}{
		{
			name:        "simple path",
			graph:       [][]int{{1, 2}, {0, 2, 3}, {0, 1}, {1}},
			origin:      0,
			destination: 3,
			want:        2,
		},
		{
			name:        "direct path",
			graph:       [][]int{{1}, {0, 2}, {1, 3}, {2}},
			origin:      0,
			destination: 3,
			want:        3,
		},
		{
			name:        "same node",
			graph:       [][]int{{1}, {0}},
			origin:      0,
			destination: 0,
			want:        0,
		},
		{
			name:        "no path exists",
			graph:       [][]int{{1}, {0}, {3}, {2}},
			origin:      0,
			destination: 2,
			want:        0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ShortestPath(tt.graph, tt.origin, tt.destination)
			if got != tt.want {
				t.Errorf("ShortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
