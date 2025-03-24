package graph

import "testing"

func TestTraversalDFS(t *testing.T) {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)
	node1.AddNeighbors(node2, node3)
	node2.AddNeighbors(node1, node3, node4)
	node3.AddNeighbors(node1, node2)
	node4.AddNeighbors(node2)

	TraversalDFS(node1)
}
