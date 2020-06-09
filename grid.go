package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Grid struct {
	rows    int
	columns int
	cells   [][]*Cell
}

func newGrid(r, c int) *Grid {

	g := &Grid{
		rows:    r,
		columns: c,
	}

	g.prepareGrid()
	g.configureCells()

	return g
}

func (g *Grid) prepareGrid() {
	cells := [][]*Cell{}
	for i := 0; i < g.rows; i++ {
		cells = append(cells, []*Cell{})
		for j := 0; j < g.columns; j++ {
			cells[i] = append(cells[i], newCell(i, j))
		}
	}
	g.cells = cells
}

func (g *Grid) configureCells() {
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.columns; j++ {
			if i > 0 {
				g.cells[i][j].n = g.cells[i-1][j]
			}
			if i < g.rows-1 {
				g.cells[i][j].s = g.cells[i+1][j]
			}
			if j > 0 {
				g.cells[i][j].w = g.cells[i][j-1]
			}
			if j < g.columns-1 {
				g.cells[i][j].e = g.cells[i][j+1]
			}
		}
	}
}

func (g *Grid) getRandomCell() *Cell {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	row := r.Intn(g.rows)
	col := r.Intn(g.columns)

	return g.cells[row][col]
}

func (g *Grid) size() int {
	return g.rows * g.columns
}

func (g *Grid) forEachCellDo(f func(c *Cell) bool) {
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.columns; j++ {
			f(g.cells[i][j])
		}
	}
}

func (g *Grid) String() string {
	str := fmt.Sprint("+")
	for j := 0; j < g.columns; j++ {
		str += fmt.Sprint("---+")
	}
	str += fmt.Sprintln()

	for i := 0; i < g.rows; i++ {
		str += fmt.Sprint("|")
		for j := 0; j < g.columns; j++ {
			str += fmt.Sprint("   |")
		}
		str += fmt.Sprintln()
		str += fmt.Sprint("+")
		for j := 0; j < g.columns; j++ {
			str += fmt.Sprint("---+")
		}
		str += fmt.Sprintln()
	}

	return str
}
