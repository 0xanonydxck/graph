package graph

import "testing"

func TestOpenLock(t *testing.T) {
	result := OpenLock("0202", []string{"0201", "0101", "0102", "1212", "2002"})
	t.Log(result)
}
