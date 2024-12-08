// https://adventofcode.com/2024/day/8
// Day 8: Resonant Collinearity
package main

import (
	"aoc/libaoc"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	// input, err := libaoc.ReadLines("example2.txt")
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}

	var matrix StringsMatrix
	matrix.buildMatrix(input)
	matrix.printMatrix()
	// matrix.printAntennas()
	fmt.Printf("%d Antinodes\n", len(matrix.antinodes))
	matrix.printAntinodes()
	// Time
	elapsed := time.Since(start)
	log.Printf("\nProgram took %s\n", elapsed)
}
