package graph

type Coordinate [2]int

func NewCoordinate(x, y int) Coordinate {
	return Coordinate{x, y}
}

func (c *Coordinate) X() int {
	return c[0]
}

func (c *Coordinate) Y() int {
	return c[1]
}
