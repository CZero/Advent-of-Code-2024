package main

import "fmt"

// Coord holds {Row, Column} (I always get confuzzled with x,y)
type Coord struct {
	r int // row
	c int // col
}

// Grid is the map containing the strings
type Grid map[Coord]string

type Antennas map[string][]Coord
type AntennaPair []Coord

type Antinodes map[Coord]bool

// StringsMatrix is the object containing a Grid and documents the height and width
type StringsMatrix struct {
	grid      Grid
	height    int
	width     int
	antennas  Antennas
	antinodes Antinodes
}

// buildMatrix is a method to build the matrix.
func (m *StringsMatrix) buildMatrix(input []string) {
	m.grid = make(Grid)
	m.antennas = make(Antennas)
	m.antinodes = make(Antinodes)
	(*m).height = len(input) - 1
	for r, line := range input {
		if (*m).width == 0 {
			(*m).width = len(line) - 1
		}
		for c, char := range line {
			m.grid[Coord{r, c}] = string(char)
			switch {
			case string(char) != ".":
				m.antennas[string(char)] = append(m.antennas[string(char)], Coord{r, c})
			}
		}
	}
	m.calcAntinodes()
}

// printMatrix is a method to visually validate the matrix.
func (m *StringsMatrix) printMatrix() {
	fmt.Println()
	var line string
	for r := 0; r <= (*m).height; r++ {
		for c := 0; c <= (*m).width; c++ {
			if c == 0 {
				line = ""
			}
			line = line + (*m).grid[Coord{r, c}]
		}
		fmt.Printf("%s\n", line)
	}
	fmt.Println()
}

func (m *StringsMatrix) printAntinodes() {
	fmt.Println()
	var line string
	for r := 0; r <= (*m).height; r++ {
		for c := 0; c <= (*m).width; c++ {
			if c == 0 {
				line = ""
			}
			_, an := m.antinodes[Coord{r, c}]
			if an {
				line = line + "#"
			} else {
				line = line + "."
			}
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

func (m *StringsMatrix) printAntennas() {
	for antenna, coords := range m.antennas {
		fmt.Printf("%v: %v\n", antenna, coords)
	}
}

func (m *StringsMatrix) calcAntinodes() {
	for _, coords := range m.antennas {
		if len(coords) > 1 {
			combis := m.getPossiblePairs(coords)
			for _, combi := range combis {
				// First we add the antennas (there's a combi, so they count per antenna)
				m.antinodes[combi[0]] = true
				m.antinodes[combi[1]] = true

				// Calculate the vector
				rVector := combi[0].r - combi[1].r
				cVector := combi[0].c - combi[1].c

				// fmt.Printf("For combi %v: %v and %v\n", combi, Coord{combi[0].r + rVector, combi[0].c + cVector}, Coord{combi[1].r - rVector, combi[1].c - cVector})
				bounce := combi[1] // First bounce is the antenna itself
				for {
					if m.inmatrix(Coord{bounce.r - rVector, bounce.c - cVector}) { // Still exists?
						m.antinodes[Coord{bounce.r - rVector, bounce.c - cVector}] = true // Antinode
						bounce = Coord{bounce.r - rVector, bounce.c - cVector}            // We BOUNCE
					} else { // Awh, no more bounces :(
						break
					}
				}
				bounce = combi[0] // First bounce is the antenna itself
				for {
					if m.inmatrix(Coord{bounce.r + rVector, bounce.c + cVector}) { // Still exists?
						m.antinodes[Coord{bounce.r + rVector, bounce.c + cVector}] = true // Antinode
						bounce = Coord{bounce.r + rVector, bounce.c + cVector}            // We BOUNCE
					} else { // Awh, no more bounces :(
						break
					}
				}
			}
		}
	}
}

// getPossiblePairs(a b c) returns a slice of all possible pairs:
// ab
// ac
// bc
func (m *StringsMatrix) getPossiblePairs(coords []Coord) (pairs []AntennaPair) {
	for i := 0; i < len(coords)-1; i++ {
		for j := i + 1; j < len(coords); j++ {
			pairs = append(pairs, []Coord{coords[i], coords[j]})
		}
	}
	fmt.Printf("Pairs: %v\n", pairs)
	return pairs
}
