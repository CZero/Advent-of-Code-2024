// https://adventofcode.com/2024/day/7
// Day 7: Bridge Repair
package main

import (
	"aoc/libaoc"
	"fmt"
	"log"
	"strconv"
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

	// for _, missingop := range missingops {
	// 	fmt.Printf("%v can be done: %t\n", missingop, findOps(missingop, true, 0))
	// }

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
	switch len(missingop.parts) { // Check for a last run
	case 0: // result of a merge
		return calc == missingop.testvalue
	case 1: // Last value to op with
		part := missingop.parts[0]
		first := strconv.Itoa(calc)
		second := strconv.Itoa(part)
		merged := libaoc.SilentAtoi(first + second)
		switch {
		case calc*missingop.parts[0] == missingop.testvalue:
			return true
		case calc+missingop.parts[0] == missingop.testvalue:
			return true
		case merged == missingop.testvalue:
			return true
		default:
			return false
		}
	}

	switch start { // First run?
	case true: // Init round, we haven't calc yet
		parts := missingop.parts[1:]
		switch {
		case findOps(MissingOp{missingop.testvalue, parts}, false, missingop.parts[0]):
			return true
		}
	case false: // Recursive round, we have calc, even if 0.
		first := strconv.Itoa(calc)
		second := strconv.Itoa(missingop.parts[0])
		merged := libaoc.SilentAtoi(first + second)
		switch {
		case findOps(MissingOp{missingop.testvalue, missingop.parts[1:]}, false, calc+missingop.parts[0]):
			return true
		case findOps(MissingOp{missingop.testvalue, missingop.parts[1:]}, false, calc*missingop.parts[0]):
			return true
		case findOps(MissingOp{missingop.testvalue, missingop.parts[1:]}, false, merged):
			return true
		default:
			return false
		}
	}
	return false
}
