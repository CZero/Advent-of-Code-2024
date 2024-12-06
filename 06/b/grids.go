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
	grid                Grid
	height              int
	width               int
	objects             map[Coord]bool
	guardStart          Coord
	guardStartDirection string
	guard               Coord
	next                Coord
	direction           string
	visited             map[Coord][]string
	loopingSpots        map[Coord]bool
}

// buildMatrix is a method to build the matrix.
func (m *StringsMatrix) BuildMatrix(input []string) {
	m.grid = make(Grid)
	m.objects = make(map[Coord]bool)
	m.visited = make(map[Coord][]string)
	m.loopingSpots = make(map[Coord]bool)
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
				m.guardStart = Coord{r, c}
				m.guardStartDirection = "n"
			case ">":
				m.guardStart = Coord{r, c}
				m.guardStartDirection = "e"
			case "<":
				m.guardStart = Coord{r, c}
				m.guardStartDirection = "w"
			case "v":
				m.guardStart = Coord{r, c}
				m.guardStartDirection = "s"
			}
		}
	}
	m.WalkTheGuard()
	m.FindLoopingSpots()
}

// FindLoopingSpots replaces all empty spaces with an object to see if we get a loop walking the guard. It adds the found spot to loopingspots
func (m *StringsMatrix) FindLoopingSpots() {
	for r := 0; r <= m.height; r++ {
		for c := 0; c <= m.width; c++ {
			switch m.grid[Coord{r, c}] {
			case ".":
				m.objects[Coord{r, c}] = true // We introduce an object
				if m.WalkTheGuard() {
					m.loopingSpots[Coord{r, c}] = true
				}
				delete(m.objects, Coord{r, c}) // And we remove it.
			}
		}
	}
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

// LoopDetected checks wether the guard was in this spot GOING THE SAME DIRECTION before, then we have a loop
func (m *StringsMatrix) LoopDetected(coord Coord, direction string) bool {
	_, ok := m.visited[coord]
	// fmt.Printf("We've seen these coords before: %v\n", coord)
	if ok {
		for _, directionSeen := range m.visited[coord] {
			// fmt.Printf("DirectionSeen: %v\n", directionSeen)
			if direction == directionSeen {
				// fmt.Printf("We've seen these coords and this direction before! %v %s", coord, direction)
				return true
			}
		}
	}
	return false
}

// WalkTheGuard makes the guard do his rounds. If we detect a loop, we stop and return true. If the guard walks out, we return false.
func (m *StringsMatrix) WalkTheGuard() (loop bool) {
	m.visited = make(map[Coord][]string) // We wipe previous runs
	m.guard = m.guardStart
	m.direction = m.guardStartDirection
	// fmt.Printf("This should be empty: %v\n", m.visited)
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
			// fmt.Println("We've reached the end")
			return false
		case m.IsObject(m.next):
			// fmt.Printf("Turn the guard, facing: %s\n", m.direction)
			m.TurnTheGuard()
			// fmt.Printf("Turned the guard, now facing: %s\n", m.direction)
		default:
			if !m.LoopDetected(m.next, m.direction) {
				m.visited[m.next] = append(m.visited[m.next], m.direction)
			} else {
				fmt.Println("Loop detected!")
				return true
			}
			m.guard = m.next
			// fmt.Printf("Guard is now at %v\n", m.guard)
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
