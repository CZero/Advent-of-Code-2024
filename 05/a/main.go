// https://adventofcode.com/2024/day/5
// Day 5: Print Queue
package main

import (
	"aoc/libaoc"
	"fmt"
	"strings"
)

func main() {
	// input, err := libaoc.ReadLines("shortexample.txt")
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	manualUpdateGuide := loadManuals(input)

	// Some visual checking:
	fmt.Printf("\nAfter rules:\n")
	manualUpdateGuide.after.printRules()
	fmt.Printf("\nPages: %v\n", manualUpdateGuide.updates)

	// Now the solution
	correctUpdates := manualUpdateGuide.getCorrectUpdates()
	fmt.Printf("\nCorrect updates: %v\n", correctUpdates)
	fmt.Printf("\nThe sum: %d\n", sumMiddles(correctUpdates))
}

func loadManuals(input []string) ManualUpdate {
	var (
		manualUpdate ManualUpdate
		lineNr       int
		updates      []Update
	)
	before := make(Instructions)
	after := make(Instructions)

	// First we do the instructions
	for nr, line := range input {
		if line == "" {
			lineNr = nr + 1
			break
		}
		fields := strings.Split(line, "|")
		values := []int{libaoc.SilentAtoi(fields[0]), libaoc.SilentAtoi(fields[1])}

		// after 0 comes 1
		if !after.hasKey(values[0]) {
			after[values[0]] = make(map[int]bool)
		}
		after[values[0]][values[1]] = true
		// before 1 comes 0
		if !before.hasKey(values[1]) {
			before[values[1]] = make(map[int]bool)
		}
		before[values[1]][values[0]] = true

	}

	// Then we load the updates
	for _, line := range input[lineNr:] {
		fields := strings.Split(line, ",")
		var update Update
		for _, field := range fields {
			update = append(update, libaoc.SilentAtoi(field))
		}
		updates = append(updates, update)

	}
	manualUpdate = ManualUpdate{
		after:   after,
		updates: updates,
	}
	return manualUpdate
}

func sumMiddles(updates []Update) (sum int) {
	for _, update := range updates {
		var half int
		half = len(update) / 2
		// fmt.Printf("Update %v is %d long, so we take pos %d: %d\n", update, len(update), half, update[half])
		sum += update[half]
	}
	return sum
}
