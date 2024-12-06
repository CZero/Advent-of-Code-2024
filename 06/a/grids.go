package main

import "fmt"

// Coord holds {Column, Row} (I always get confuzzled with x,y)
type Coord struct {
	r int // row
	c int // col
}

// Grid is the map containing the strings
type Grid map[Coord]string

// StringsMatrix is the object containing a Grid and documents the height and width
type StringsMatrix struct {
	grid      Grid
	height    int
	width     int
	objects   map[Coord]bool
	guard     Coord
	next      Coord
	direction string
	visited   map[Coord]bool
}

// buildMatrix is a method to build the matrix.
func (m *StringsMatrix) BuildMatrix(input []string) {
	m.grid = make(Grid)
	m.objects = make(map[Coord]bool)
	m.visited = make(map[Coord]bool)
	(*m).height = len(input) - 1
	for r, line := range input {
		if (*m).width == 0 {
			(*m).width = len(line) - 1
		}
		for c, char := range line {
			m.grid[Coord{r, c}] = string(char)
			switch string(char) {
			case "#":
				m.objects[Coord{r, c}] = true
			case "^":
				m.guard = Coord{r, c}
				m.direction = "n"
			case ">":
				m.guard = Coord{r, c}
				m.direction = "e"
			case "<":
				m.guard = Coord{r, c}
				m.direction = "w"
			case "v":
				m.guard = Coord{r, c}
				m.direction = "s"
			}
		}
	}
	m.WalkTheGuard()
}

// printMatrix is a method to visually validate the matrix.
func (m *StringsMatrix) PrintMatrix() {
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
	fmt.Printf("Objects: %v\n", m.objects)
}

// inMatrix is a method to validate if a point is in the matrix.
func (m *StringsMatrix) InMatrix(coord Coord) bool {
	_, ok := m.grid[coord]
	return ok
}

// WalkTheGuard makes the guard do it's round
func (m *StringsMatrix) WalkTheGuard() {
	for {
		switch m.direction {
		case "n":
			m.next = Coord{m.guard.r - 1, m.guard.c}
		case "s":
			m.next = Coord{m.guard.r + 1, m.guard.c}
		case "e":
			m.next = Coord{m.guard.r, m.guard.c + 1}
		case "w":
			m.next = Coord{m.guard.r, m.guard.c - 1}
		}
		switch {
		case !m.InMatrix(m.next):
			return
		case m.IsObject(m.next):
			m.TurnTheGuard()
		default:
			m.visited[m.next] = true
			m.guard = m.next
		}
	}
}

// TurnTheGuard turns the guard 90 degrees to the right
func (m *StringsMatrix) TurnTheGuard() {
	switch m.direction {
	case "n":
		m.direction = "e"
	case "e":
		m.direction = "s"
	case "s":
		m.direction = "w"
	case "w":
		m.direction = "n"
	}
}

// IsObject checks if the coord is in the object list
func (m *StringsMatrix) IsObject(coord Coord) bool {
	_, ok := m.objects[coord]
	return ok
}
