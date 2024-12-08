package libaoc

import "fmt"

// These are based on AoC2024 D4

// Coord holds {Row, Column} (I always get confuzzled with x,y)
type Coord struct {
	r int // row
	c int // col
}

// Grid is the map containing the strings
type Grid map[Coord]string

// StringsMatrix is the object containing a Grid and documents the height and width
type StringsMatrix struct {
	grid   Grid
	height int
	width  int
}

// buildMatrix is a method to build the matrix.
func (m *StringsMatrix) buildMatrix(input []string) {
	m.grid = make(Grid)
	(*m).height = len(input) - 1
	for r, line := range input {
		if (*m).width == 0 {
			(*m).width = len(line) - 1
		}
		for c, char := range line {
			m.grid[Coord{r, c}] = string(char)
		}
	}
}

// printMatrix is a method to visually validate the matrix.
func (m *StringsMatrix) printMatrix() {
	fmt.Println()
	var line string
	for r := 0; r < (*m).height; r++ {
		for c := 0; c < (*m).width; c++ {
			if c == 0 {
				line = ""
			}
			line = line + (*m).grid[Coord{r, c}]
		}
		fmt.Printf("%s\n", line)
	}
	fmt.Println()
}

// inMatrix is a method to validate if a point is in the matrix.
func (m *StringsMatrix) inmatrix(coord Coord) bool {
	_, ok := m.grid[coord]
	return ok
}

// wordSearch is a method to find the number of occurrences for a word in all directions, horizontal, vertical, diagonal
func (m *StringsMatrix) wordSearch(word string) (occurrences int) {
	firstLetter := string(word[0])
	directions := []string{"n", "ne", "e", "se", "s", "sw", "w", "nw"}
	for r := 0; r < (*m).height; r++ {
		for c := 0; c < (*m).width; c++ {
			switch m.grid[Coord{r, c}] {
			case firstLetter:
				// fmt.Printf("Found X: i %d, j %d\n", i, j)
				for _, direction := range directions {
					if m.searchDirection(direction, word, r, c) {
						occurrences++
					}
				}
			}
		}
	}
	return occurrences
}

// searchDirection is a method that searches a "word" in a "direction", starting at i,j
func (m *StringsMatrix) searchDirection(direction, word string, r, c int) (found bool) {
	for _, letter := range word {
		if !(*m).inmatrix(Coord{r, c}) {
			return false
		}
		if (*m).grid[Coord{r, c}] == string(letter) { // If the letter is correct, prep for the next
			switch direction {
			case "n":
				r--
			case "e":
				c++
			case "s":
				r++
			case "w":
				c--
			case "ne":
				c++
				r--
			case "nw":
				c--
				r--
			case "se":
				c++
				r++
			case "sw":
				c--
				r++
			}
		} else { // The found letter was different
			return false
		}
	}
	return true // We never had a miss
}
