package libaoc

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ReadLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// SilentAtoi returns an int or panics
func SilentAtoi(input string) int {
	val, err := strconv.Atoi(input)
	if err != nil {
		panic(fmt.Sprintf("SilentAtoi failed with %v as input", val))
	}
	return val
}

// SumSlice returns the sum of a slice
func SumSlice(inOrder []int) (sum int) {
	for _, item := range inOrder {
		sum += item
	}
	return sum
}

// ProductSlice returns the product of a slice
func ProductSlice(inOrder []int) (sum int) {
	for _, item := range inOrder {
		sum *= item
	}
	return sum
}

// ReverseString returns a reversed string
func ReverseString(word string) string {
	var reversed string
	for i := len(word) - 1; i >= 0; i-- {
		reversed = reversed + string(word[i])
	}
	return reversed
}

// Paniccheck panics on error
func Paniccheck(e error) {
	if e != nil {
		panic(e)
	}
}

// RemoveSpaces removes all extra whitespace from a string
func RemoveSpaces(input string) string {
	return strings.Join(strings.Fields(input), " ")
}

// SortIntSlice sorts an intslice
func sortIntSlice(input []int) []int {
	sort.Ints(input)
	return input
}
