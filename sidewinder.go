package main

import (
	"math/rand"
	"time"
)

func sidewinder(g *Grid) {
	g.forEachRowDo(func(r []*Cell) {
		run := []*Cell{}

		for _, c := range r {
			run = append(run, c)

			atEasternBounary := c.e == nil
			atNorthernBounary := c.n == nil

			s := rand.NewSource(time.Now().UnixNano())
			r := rand.New(s)

			shouldCloseOut := atEasternBounary || (!atNorthernBounary && r.Intn(2) == 0)

			if shouldCloseOut {
				m := run[r.Intn(len(run))]
				if m.n != nil {
					m.Link(m.n, true)
					run = []*Cell{}
				}
			} else {
				c.Link(c.e, true)
			}
		}
	})
}
