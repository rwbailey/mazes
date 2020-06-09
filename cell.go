package main

type Cell struct {
	row   int
	col   int
	links map[*Cell]bool
	n     *Cell
	e     *Cell
	s     *Cell
	w     *Cell
}

func newCell(r, c int) *Cell {
	return &Cell{
		row:   r,
		col:   c,
		links: map[*Cell]bool{},
	}
}

func (c *Cell) Link(cell *Cell, b bool) {
	c.links[cell] = true
	if b {
		cell.Link(c, false)
	}
}

func (c *Cell) isLinked(cell *Cell) bool {
	_, ok := c.links[cell]
	return ok
}

func (c *Cell) getNeighbors() []*Cell {
	neighbors := []*Cell{}
	if c.n != nil {
		neighbors = append(neighbors, c.n)
	}
	if c.e != nil {
		neighbors = append(neighbors, c.e)
	}
	if c.s != nil {
		neighbors = append(neighbors, c.s)
	}
	if c.w != nil {
		neighbors = append(neighbors, c.w)
	}
	return neighbors
}
