// https://adventofcode.com/2024/day/3
// Day 3: Mull It Over
package main

import (
	"aoc/libaoc"
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	a int
	b int
}

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	instructions := getInstructions(input)
	fmt.Printf("We'll multiply the pairs:\n%v\n\nThe total product is: %d\n", instructions, getTotalProduct(instructions))
}

func getInstructions(input []string) []Instruction {
	var instructions []Instruction
	for _, line := range input {
		step1 := strings.Split(line, "mul(") // We break by "mul("'s, so each field should be the start
		for _, part1 := range step1 {
			step2 := strings.Split(part1, ")") // We break the pieces by ")". Now they SHOULD contain x,x (at least the first field)
			// fmt.Printf("%#v\n", step2)
			step3 := strings.Split(step2[0], ",") // We break the pieces by ","
			if len(step3) == 2 {                  // More or less would have been garbage
				// Now let's validate if they're numbers
				a, err := strconv.Atoi(step3[0])
				if err == nil {
					b, err := strconv.Atoi(step3[1])
					if err == nil {
						instructions = append(instructions, Instruction{a: a, b: b})
					}
				}
			}
		}
	}
	// fmt.Printf("%#v\n", instructions)
	return instructions
}

func getTotalProduct(instructions []Instruction) (product int) {
	for _, instruction := range instructions {
		product += instruction.a * instruction.b
	}
	return product
}
