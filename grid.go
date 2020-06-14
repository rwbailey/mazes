package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
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

// Return a random cell
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
func (g *Grid) forEachRowDo(f func(r []*Cell)) {
	for i := 0; i < g.rows; i++ {
		f(g.cells[i])
	}
}

// Pass each cell to the function f()
func (g *Grid) forEachCellDo(f func(c *Cell)) {
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.columns; j++ {
			f(g.cells[i][j])
		}
	}
}

// Custom format for printing the grid to string
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

// Print the grid to a png
func (g *Grid) toPNG(cellSize int, trans bool) {

	width := g.columns*cellSize + 1
	height := g.rows*cellSize + 1

	origin := image.Point{0, 0}
	end := image.Point{width, height}

	background := color.White
	wall := color.Black

	img := image.NewRGBA(image.Rectangle{origin, end})
	if !trans {
		draw.Draw(img, img.Bounds(), &image.Uniform{background}, image.ZP, draw.Src)
	}

	g.forEachCellDo(func(c *Cell) {
		x1 := c.col * cellSize
		y1 := c.row * cellSize
		x2 := (c.col + 1) * cellSize
		y2 := (c.row + 1) * cellSize

		if c.n == nil {
			for i := x1; i <= x2; i++ {
				img.Set(i, y1, wall)
			}
		}
		if c.w == nil {
			for j := y1; j <= y2; j++ {
				img.Set(x1, j, wall)
			}
		}

		if !c.isLinked(c.s) {
			for i := x1; i <= x2; i++ {
				img.Set(i, y2, wall)
			}
		}
		if !c.isLinked(c.e) {
			for j := y1; j <= y2; j++ {
				img.Set(x2, j, wall)
			}
		}
	})

	f, err := os.Create("draw.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
