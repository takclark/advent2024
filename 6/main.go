package main

import (
	"fmt"
	"strings"

	"github.com/takclark/advent2024/parsing"
)

func main() {
	input := parsing.MustReadFile("input")
	// input := parsing.MustReadFile("test")
	part1, part2 := solve(input)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}

func solve(input string) (int, int) {
	var total1, total2 int

	m := parseInput(input)
	fmt.Println(m)

	for m.tick() {
		fmt.Println(m.grid)
	}

	for i := range len(m.grid) {
		for j := range len(m.grid[i]) {
			if m.grid[i][j] == "X" {
				total1++
			}
		}
	}

	return total1, total2
}

func parseInput(input string) *labMap {
	var grid [][]string
	var startingGuardPos tuple
	lines := strings.Split(input, "\n")
	for y, l := range lines {
		if l == "" {
			continue
		}
		grid = append(grid, strings.Split(l, ""))
		x := strings.Index(l, "^") // she starts -y in test/real input
		if x != -1 {
			startingGuardPos = tuple{x, y}
		}
		// guard may re-cross her position
	}

	return &labMap{
		grid:     grid,
		guardPos: startingGuardPos,
		guardDir: tuple{x: 0, y: -1},
	}
}

func (m *labMap) tick() bool {
	m.grid[m.guardPos.y][m.guardPos.x] = "X"

	next := tuple{
		x: m.guardPos.x + m.guardDir.x,
		y: m.guardPos.y + m.guardDir.y,
	}

	if next.x < 0 || next.x >= len(m.grid[0]) {
		return false
	}

	if next.y < 0 || next.y >= len(m.grid) {
		return false
	}

	if m.grid[next.y][next.x] == "#" {
		m.guardDir = turn(m.guardDir)
	} else {
		m.guardPos = next
	}

	return true
}

func turn(t tuple) tuple {
	if t.x == 0 && t.y == -1 {
		return tuple{x: 1, y: 0}
	}
	if t.x == 1 && t.y == 0 {
		return tuple{x: 0, y: 1}
	}
	if t.x == 0 && t.y == 1 {
		return tuple{x: -1, y: 0}
	}

	return tuple{x: 0, y: -1}
}

type tuple struct{ x, y int }

type labMap struct {
	grid     [][]string
	guardPos tuple
	guardDir tuple // dx, dy
}
