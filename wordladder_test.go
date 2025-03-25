package graph

import "testing"

func TestWordLadder(t *testing.T) {
	result := WordLadder("COLD", "WARM", []string{"COLD", "GOLD", "CORD", "SOLD", "CARD", "WARD", "WARM", "TARD"})
	t.Log(result)
}
