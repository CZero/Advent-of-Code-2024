// https://adventofcode.com/2023/day/
// Day x:
package main

import (
	"aoc/libaoc"
	"fmt"
	"sort"
	"strings"
)

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	listA, listB := getLists(input)
	sort.Ints(listA)
	sort.Ints(listB)
	distance := sumDistances(listA, listB)
	fmt.Printf("Distance: %d\n", distance)
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

func sumDistances(listA, listB []int) (distance int) {
	for i := 0; i < len(listA); i++ {
		switch {
		case listA[i] > listB[i]:
			distance += listA[i] - listB[i]
		default:
			distance += listB[i] - listA[i]
		}
	}
	return distance
}

func removeSpaces(input string) string {
	return strings.Join(strings.Fields(input), " ")
}
