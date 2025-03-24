package graph

import "errors"

type Node struct {
	Value     int
	Neighbors []*Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

func (n *Node) AddNeighbors(nodes ...*Node) *Node {
	n.Neighbors = append(n.Neighbors, nodes...)
	return n
}

func BuildGraph(graph [][]int) (*Node, error) {
	if len(graph) == 0 {
		return nil, errors.New("graph is empty")
	}

	nodes := make([]*Node, len(graph))
	for i := range graph {
		nodes[i] = NewNode(i)
	}

	for i := range graph {
		for _, j := range graph[i] {
			nodes[i].AddNeighbors(nodes[j])
		}
	}

	return nodes[0], nil
}
