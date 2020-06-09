package main

import "fmt"

func main() {
	grid := newGrid(4, 3)

	printVals := func(c *Cell) bool {
		fmt.Printf("[%v,%v]\n", c.row, c.col)
		return true
	}

	grid.forEachCellDo(printVals)

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
