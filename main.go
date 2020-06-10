package main

import "fmt"

func main() {
	grid := newGrid(8, 8)

	binaryTree(grid)
	fmt.Println(grid)

	// grid.cells[0][0].Link(grid.cells[0][1], true)
	// grid.cells[0][1].Link(grid.cells[1][1], true)
	// grid.cells[1][1].Link(grid.cells[1][2], true)
	// grid.cells[1][2].Link(grid.cells[1][3], true)
	// grid.cells[1][3].Link(grid.cells[2][3], true)
	// grid.cells[2][3].Link(grid.cells[3][3], true)
	// grid.cells[3][3].Link(grid.cells[3][2], true)
	// grid.cells[3][2].Link(grid.cells[4][2], true)
	// grid.cells[4][2].Link(grid.cells[5][2], true)
	// grid.cells[5][2].Link(grid.cells[6][2], true)
	// grid.cells[6][2].Link(grid.cells[7][2], true)
	// grid.cells[7][2].Link(grid.cells[7][3], true)
	// grid.cells[7][3].Link(grid.cells[7][4], true)
	// grid.cells[7][4].Link(grid.cells[7][5], true)
	// grid.cells[7][5].Link(grid.cells[7][6], true)
	// grid.cells[7][6].Link(grid.cells[7][7], true)
	// grid.cells[7][7].Link(grid.cells[6][7], true)
	// grid.cells[6][7].Link(grid.cells[6][6], true)
	// grid.cells[6][6].Link(grid.cells[6][5], true)
	// grid.cells[6][5].Link(grid.cells[5][5], true)
	// grid.cells[5][5].Link(grid.cells[4][5], true)
	// grid.cells[4][5].Link(grid.cells[3][5], true)
	// grid.cells[3][5].Link(grid.cells[3][4], true)
	// grid.cells[3][4].Link(grid.cells[3][3], true)
	// grid.cells[3][5].Link(grid.cells[2][5], true)
	// grid.cells[2][5].Link(grid.cells[1][5], true)
	// grid.cells[1][5].Link(grid.cells[0][5], true)
	// grid.cells[0][5].Link(grid.cells[0][6], true)
	// grid.cells[0][6].Link(grid.cells[0][7], true)
	// fmt.Println(grid)

	// printVals := func(c *Cell) bool {
	// 	fmt.Printf("[%v,%v]\n", c.row, c.col)
	// 	return true
	// }

	// grid.forEachCellDo(printVals)

	// r := grid.getRandomCell()
	// fmt.Printf("[%v,%v]\n", r.row, r.col)
	// if r.n != nil {
	// 	fmt.Printf("n:[%v,%v]\n", r.n.row, r.n.col)
	// } else {
	// 	fmt.Printf("n:[%v,%v]\n", "-", "-")
	// }
	// if r.e != nil {
	// 	fmt.Printf("e:[%v,%v]\n", r.e.row, r.e.col)
	// } else {
	// 	fmt.Printf("e:[%v,%v]\n", "-", "-")
	// }
	// if r.s != nil {
	// 	fmt.Printf("s:[%v,%v]\n", r.s.row, r.s.col)
	// } else {
	// 	fmt.Printf("s:[%v,%v]\n", "-", "-")
	// }
	// if r.w != nil {
	// 	fmt.Printf("w:[%v,%v]\n", r.w.row, r.w.col)
	// } else {
	// 	fmt.Printf("w:[%v,%v]\n", "-", "-")
	// }

	// for _, x := range grid.cells {
	// 	for _, y := range x {
	// 		fmt.Printf("[%v,%v] ", y.row, y.col)
	// 		if y.n != nil {
	// 			fmt.Printf("- n:[%v,%v] ", y.n.row, y.n.col)
	// 		} else {
	// 			fmt.Printf("- n:[%v,%v] ", "-", "-")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
}
