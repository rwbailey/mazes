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

// Populate the cells in the grid
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

// Set the n,s,e, & w fields of each cell in the grid
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

// Return a randoom cell
func (g *Grid) getRandomCell() *Cell {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	row := r.Intn(g.rows)
	col := r.Intn(g.columns)

	return g.cells[row][col]
}

// Return the number of cells in the grid
func (g *Grid) size() int {
	return g.rows * g.columns
}

// Pass each row to the function f()
func (g *Grid) forEashRowDo(f func(r []*Cell) bool) {
	for i := 0; i < g.rows; i++ {
		f(g.cells[i])
	}
}

// Pass each cell to the function f()
func (g *Grid) forEachCellDo(f func(c *Cell) bool) {
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.columns; j++ {
			f(g.cells[i][j])
		}
	}
}

// Custom format for printing the grid
func (g *Grid) String() string {
	str := fmt.Sprint("+")
	for j := 0; j < g.columns; j++ {
		str += fmt.Sprint("---+")
	}
	str += fmt.Sprintln()

	for i := 0; i < g.rows; i++ {
		str += fmt.Sprint("|")
		for j := 0; j < g.columns; j++ {
			str += fmt.Sprint("   ")
			if g.cells[i][j].isLinked(g.cells[i][j].e) {
				str += fmt.Sprint(" ")
			} else {
				str += fmt.Sprint("|")
			}
		}
		str += fmt.Sprintln()
		str += fmt.Sprint("+")
		for j := 0; j < g.columns; j++ {
			if g.cells[i][j].isLinked(g.cells[i][j].s) {
				str += fmt.Sprint("   +")
			} else {
				str += fmt.Sprint("---+")
			}
		}
		str += fmt.Sprintln()
	}
	return str
}
