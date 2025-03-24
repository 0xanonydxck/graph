package graph

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
