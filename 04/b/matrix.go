package main

import "fmt"

// Coord holds {Column, Row} (I always get confuzzled with x,y)
type Coord struct {
	c int // col
	r int // row
}

// Grid is the map containing the strings
type Grid map[Coord]string

// Matrix is the object containing a Grid and documents the height and width
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
	_, ok := m.grid[coord]
	return ok
}

// wordSearch is a method to find the number of occurrences for a word in all directions, horizontal, vertical, diagonal
func (m *Matrix) wordSearch(word string) (occurrences int) {
	firstLetter := string(word[0])
	directions := []string{"n", "ne", "e", "se", "s", "sw", "w", "nw"}
	for j := 0; j < (*m).height; j++ {
		for i := 0; i < (*m).width; i++ {
			switch m.grid[Coord{i, j}] {
			case firstLetter:
				// fmt.Printf("Found X: i %d, j %d\n", i, j)
				for _, direction := range directions {
					if m.searchDirection(direction, word, i, j) {
						occurrences++
					}
				}
			}
		}
	}
	return occurrences
}

// xmasSearch is a method to find the "x-mas" issue
func (m *Matrix) xmasSearch() (occurrences int) {
	for j := 0; j < (*m).height; j++ {
		for i := 0; i < (*m).width; i++ {
			if m.grid[Coord{i, j}] == "A" {
				fmt.Printf("A found: i %d, j%d\n", i, j)
				var legs int
				if m.searchDirection("se", "MAS", i-1, j-1) {
					legs++
					if legs == 2 {
						occurrences++
					}
				}
				if m.searchDirection("ne", "MAS", i-1, j+1) {
					legs++
					if legs == 2 {
						occurrences++
					}
				}
				if m.searchDirection("sw", "MAS", i+1, j-1) {
					legs++
					if legs == 2 {
						occurrences++
					}
				}
				if m.searchDirection("nw", "MAS", i+1, j+1) {
					legs++
					if legs == 2 {
						occurrences++
					}
				}
			}
		}
	}
	return occurrences
}

// searchDirection is a method that searches a "word" in a "direction", starting at i,j
func (m *Matrix) searchDirection(direction, word string, i, j int) (found bool) {
	for _, letter := range word {
		if !(*m).inmatrix(Coord{i, j}) {
			return false
		}
		if (*m).grid[Coord{i, j}] == string(letter) { // If the letter is correct, prep for the next
			switch direction {
			case "n":
				j--
			case "e":
				i++
			case "s":
				j++
			case "w":
				i--
			case "ne":
				i++
				j--
			case "nw":
				i--
				j--
			case "se":
				i++
				j++
			case "sw":
				i--
				j++
			}
		} else { // The found letter was different
			return false
		}
	}
	return true // We never had a miss
}
