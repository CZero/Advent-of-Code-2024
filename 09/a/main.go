// https://adventofcode.com/2024/day/9
// Day 9: Disk Fragmenter
package main

import (
	"aoc/libaoc"
	"log"
	"time"
)

func main() {
	start := time.Now()

	// input, err := libaoc.ReadLines("shortexample.txt")
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}

	var disk Disk
	disk.build(input)
	disk.defrag()
	disk.calcChecksum()

	// Time
	elapsed := time.Since(start)
	log.Printf("\nProgram took %s\n", elapsed)
}
