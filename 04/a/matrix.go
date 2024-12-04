package main

import "fmt"

type Coord struct {
	c int // col
	r int // row
}
type Grid map[Coord]string

type Matrix struct {
	grid   Grid
	height int
	width  int
}

// buildMatrix is a method to build the matrix.
func (m *Matrix) buildMatrix(input []string) {
	grid := m.grid
	(*m).height = len(input)
	for i, line := range input {
		if (*m).width == 0 {
			(*m).width = len(line)
		}
		for j, char := range line {
			grid[Coord{j, i}] = string(char)
		}
	}
	m.grid = grid
}

// printMatrix is a method to visually validate the matrix.
func (m *Matrix) printMatrix() {
	fmt.Println()
	var line string
	for j := 0; j < (*m).height; j++ {
		for i := 0; i < (*m).width; i++ {
			if i == 0 {
				line = ""
			}
			line = line + (*m).grid[Coord{i, j}]
		}
		fmt.Printf("%s\n", line)
	}
	fmt.Println()
}

// inMatrix is a method to validate if a point is in the matrix.
func (m *Matrix) inmatrix(coord Coord) bool {
	_, ok := m.grid[coord] // This works because the matrix is a map. This checks if an entry is in the map.
	return ok
}

// wordSearch is a method to find the number of occurrences for a word.
func (m *Matrix) wordSearch(word string) (occurrences int) {
	firstLetter := string(word[0])
	for j := 0; j < (*m).height; j++ {
		for i := 0; i < (*m).width; i++ {
			switch m.grid[Coord{i, j}] {
			case firstLetter:
				fmt.Printf("Found X: i %d, j %d\n", i, j)
				if m.searchDirection("north", word, i, j) {
					occurrences++
				}
				if m.searchDirection("northeast", word, i, j) {
					occurrences++
				}
				if m.searchDirection("east", word, i, j) {
					occurrences++
				}
				if m.searchDirection("southeast", word, i, j) {
					occurrences++
				}
				if m.searchDirection("south", word, i, j) {
					occurrences++
				}
				if m.searchDirection("southwest", word, i, j) {
					occurrences++
				}
				if m.searchDirection("west", word, i, j) {
					occurrences++
				}
				if m.searchDirection("northwest", word, i, j) {
					occurrences++
				}
			}
		}
	}
	return occurrences
}

// searchDirection is a method that searches a "word" in a "direction", starting at i,j
func (m *Matrix) searchDirection(direction, word string, i, j int) (found bool) {
	// fmt.Println("north")
	for _, letter := range word {
		if !(*m).inmatrix(Coord{i, j}) {
			return false
		}
		if (*m).grid[Coord{i, j}] == string(letter) { // If the letter is correct, prep for the next
			switch direction {
			case "north":
				j--
			case "east":
				i++
			case "south":
				j++
			case "west":
				i--
			case "northeast":
				i++
				j--
			case "northwest":
				i--
				j--
			case "southeast":
				i++
				j++
			case "southwest":
				i--
				j++
			}
		} else { // The found letter was different
			return false
		}
	}
	return true // We never had a miss
}
