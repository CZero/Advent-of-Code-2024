// https://adventofcode.com/2024/day/4
// Day 4: Ceres Search
package main

import (
	"aoc/libaoc"
	"fmt"
)

func main() {
	// input, err := libaoc.ReadLines("grid.txt")
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	var matrix Matrix
	matrix.grid = make(Grid)
	matrix.buildMatrix(input)
	matrix.printMatrix()
	fmt.Printf("Xmas found %d times\n", matrix.wordSearch("XMAS"))
}
