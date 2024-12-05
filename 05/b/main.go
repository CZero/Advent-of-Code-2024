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
	// fmt.Printf("Before rules:\n")
	// manualUpdateGuide.before.printRules()
	// fmt.Printf("\nAfter rules:\n")
	// manualUpdateGuide.after.printRules()
	// fmt.Printf("\nPages: %v\n", manualUpdateGuide.updates)

	// Get the wrong updates
	wrongUpdates := manualUpdateGuide.getWrongUpdates()
	corrected := fixWrongUpdates(wrongUpdates, manualUpdateGuide)
	// fmt.Println(corrected)
	fmt.Printf("The corrected sum: %d\n", sumMiddles(corrected))

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

func findCorrectOrder(update Update, manualUpdateGuide ManualUpdate) Update {
	check, one, two := manualUpdateGuide.validateUpdateVerbose(update)
	var newUpdate Update
	if check {
		// fmt.Println(update)
		return update
	}
	newUpdate = append(newUpdate, update[:one]...)
	newUpdate = append(newUpdate, update[one+1:two+1]...)
	newUpdate = append(newUpdate, update[one])
	newUpdate = append(newUpdate, update[two+1:]...)
	// fmt.Printf("This is better now. From: %v To %v -- We put %d behind %d", update, newUpdate, one, two)
	return (findCorrectOrder(newUpdate, manualUpdateGuide))
}

func fixWrongUpdates(updates []Update, manualUpdateGuide ManualUpdate) (fixed []Update) {
	for _, update := range updates {
		fixed = append(fixed, findCorrectOrder(update, manualUpdateGuide))
	}
	return fixed
}
