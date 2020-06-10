package main

import (
	"math/rand"
	"time"
)

func binaryTree(g *Grid) {
	g.forEachCellDo(func(c *Cell) {
		neighbors := []*Cell{}
		if c.n != nil {
			neighbors = append(neighbors, c.n)
		}
		if c.e != nil {
			neighbors = append(neighbors, c.e)
		}

		if len(neighbors) > 0 {
			s := rand.NewSource(time.Now().UnixNano())
			r := rand.New(s)
			rn := r.Intn(len(neighbors))
			c.Link(neighbors[rn], true)
		}
	})
}
