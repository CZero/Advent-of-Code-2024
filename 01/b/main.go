// https://adventofcode.com/2023/day/1
// Day 1: Historian Hysteria
package main

import (
	"fmt"
	"strings"

	"aoc/libaoc"
)

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	listA, listB := getLists(input)
	sum := sumSimiliarity(listA, listB)
	fmt.Printf("Sum: %d\n", sum)
}

func getLists(input []string) (listA, listB []int) {
	for _, line := range input {
		line := removeSpaces(line)
		values := strings.Split(line, " ")
		listA = append(listA, libaoc.SilentAtoi(values[0]))
		listB = append(listB, libaoc.SilentAtoi(values[1]))
	}
	return listA, listB
}

func sumSimiliarity(listA, listB []int) (sum int) {
	for _, numA := range listA {
		var count int
		for _, numB := range listB {
			if numA == numB {
				count++
			}
		}
		sum += numA * count
	}
	return sum
}

func removeSpaces(input string) string {
	return strings.Join(strings.Fields(input), " ")
}
