// https://adventofcode.com/2024/day/6
// Day 6: Guard Gallivant
package main

import (
	"aoc/libaoc"
	"fmt"
)

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	var matrix StringsMatrix
	matrix.BuildMatrix(input)
	matrix.PrintMatrix()
	fmt.Printf("Places visited: %d\n", len(matrix.visited))
}
