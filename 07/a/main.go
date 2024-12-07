// https://adventofcode.com/2024/day/7
// Day 7: Bridge Repair
package main

import (
	"aoc/libaoc"
	"fmt"
	"log"
	"strings"
	"time"
)

type MissingOp struct {
	testvalue int
	parts     []int
}

func main() {
	start := time.Now()
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	missingops := readTests(input)
	// fmt.Printf("Missingops: %v\n", missingops)
	fmt.Printf("Sum of the possible missingops: %d\n", findSumPossibleOps(missingops))
	// Time notation
	elapsed := time.Since(start)
	log.Printf("\nProgram took %s\n", elapsed)
}

// readTests reads the input to the MissingOps.
func readTests(input []string) (missingops []MissingOp) {
	for _, line := range input {
		var missingop MissingOp
		step1 := strings.Split(line, ":")
		missingop.testvalue = libaoc.SilentAtoi(step1[0])
		step2 := strings.Fields(step1[1])
		for _, num := range step2 {
			missingop.parts = append(missingop.parts, libaoc.SilentAtoi(num))
		}
		missingops = append(missingops, missingop)
	}
	return missingops
}

func findSumPossibleOps(missingops []MissingOp) (sum int) {
	for _, missingop := range missingops {
		if findOps(missingop, true, 0) {
			sum += missingop.testvalue
		}
	}
	return sum
}

// findOps finds if a MissingOp is possible. You pass the missingop, start and 0 for the first.
func findOps(missingop MissingOp, start bool, calc int) (found bool) {
	if len(missingop.parts) == 1 { // Last value to op with
		if calc*missingop.parts[0] == missingop.testvalue {
			return true
		}
		if calc+missingop.parts[0] == missingop.testvalue {
			return true
		}
		return false
	}
	switch start {
	case true: // Init round, we haven't calc yet
		parts := missingop.parts[1:]
		return (findOps(MissingOp{missingop.testvalue, parts}, false, missingop.parts[0]))
	case false: // Recursive round, we have calc, even if 0.
		parts := missingop.parts[1:]
		if findOps(MissingOp{missingop.testvalue, parts}, false, calc+missingop.parts[0]) {
			return true
		} else {
			if findOps(MissingOp{missingop.testvalue, parts}, false, calc*missingop.parts[0]) {
				return true
			}
		}
	}
	return false
}
