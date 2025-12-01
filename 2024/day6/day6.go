package main

import (
	"bufio"
	"fmt"
	"log"
	"maps"
	"os"
	"time"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Position struct {
	Y int
	X int
}
type ParsedMap struct {
	Direction Direction
	WholeMap  map[Position]string
	Tiles     map[Position]string
	Obstacles map[Position]string
	Visited   map[Position]bool
	Start     Position
	States    map[State]bool
}
type State struct {
	Pos       Position
	Direction Direction
}

var count int

func main() {
	start := time.Now()
	parsedMap := parseFile("day6.txt")
	part1(parsedMap)
	part1time := time.Since(start)
	fmt.Println("part1: ", part1time)
	part2(parsedMap)
	elapsed := time.Since(start)

	fmt.Println("part2: ", elapsed)
}

func part1(parsedMap *ParsedMap) {
	current := parsedMap.Start
	move1(parsedMap, current)

}
func part2(parsedMap *ParsedMap) {
	for pos := range parsedMap.WholeMap {

		if parsedMap.Obstacles[pos] == "obstacle" {
			continue
		}
		m := cloneMaps(parsedMap)
		m.Obstacles[pos] = "obstacle"
		a := move(m, m.Start)
		if a {
			count++
		}
	}
	fmt.Println(count)

}
func move1(m *ParsedMap, current Position) {
	var next Position

	for {
		switch m.Direction {
		case Up:
			next = Position{Y: current.Y - 1, X: current.X}
		case Right:
			next = Position{Y: current.Y, X: current.X + 1}
		case Down:
			next = Position{Y: current.Y + 1, X: current.X}
		case Left:
			next = Position{Y: current.Y, X: current.X - 1}
		}

		// Check if next position is inside the map
		if _, ok := m.WholeMap[next]; !ok {
			fmt.Printf("guard has exited the map, count: %d\n", len(m.Visited))
			return
		}

		// Check for obstacle
		if m.Obstacles[next] == "obstacle" {
			// change direction clockwise
			m.Direction = (m.Direction + 1) % 4
			continue // recalc next with new direction
		}

		// move forward
		current = next
		m.Visited[current] = true
	}
}
func move(m *ParsedMap, current Position) bool {
	var next Position

	for {
		state := State{Pos: current, Direction: m.Direction}

		if m.States[state] {
			return true
		}
		m.States[state] = true

		switch m.Direction {
		case Up:
			next = Position{Y: current.Y - 1, X: current.X}
		case Right:
			next = Position{Y: current.Y, X: current.X + 1}
		case Down:
			next = Position{Y: current.Y + 1, X: current.X}
		case Left:
			next = Position{Y: current.Y, X: current.X - 1}
		}

		if _, ok := m.WholeMap[next]; !ok {
			return false
		}
		if m.Obstacles[next] == "obstacle" {
			m.Direction = (m.Direction + 1) % 4
			continue
		}

		current = next
		m.Visited[next] = true
	}
}

func cloneMaps(orig *ParsedMap) *ParsedMap {
	return &ParsedMap{
		Direction: orig.Direction,
		Start:     orig.Start,
		WholeMap:  maps.Clone(orig.WholeMap),
		Tiles:     maps.Clone(orig.Tiles),
		Obstacles: maps.Clone(orig.Obstacles),
		Visited:   maps.Clone(orig.Visited),
		States:    make(map[State]bool),
	}
}

func parseFile(fileName string) *ParsedMap {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := &ParsedMap{
		WholeMap:  make(map[Position]string),
		Tiles:     make(map[Position]string),
		Obstacles: make(map[Position]string),
		Visited:   make(map[Position]bool),
		States:    make(map[State]bool),
	}
	var index int = 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, char := range line {
			result.WholeMap[Position{Y: index, X: i}] = "tile"
			switch char {
			case '.':
				result.Tiles[Position{Y: index, X: i}] = "open"
			case '#':
				result.Obstacles[Position{Y: index, X: i}] = "obstacle"
			case '^':
				result.Visited[Position{Y: index, X: i}] = true
				result.Start = Position{Y: index, X: i}
				result.Direction = Up
				result.States[State{Pos: Position{Y: index, X: i}, Direction: result.Direction}] = true
			}
		}
		index++
	}
	return result
}
