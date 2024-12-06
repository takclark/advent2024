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
	initialGuardDir := m.guardDir
	initialGuardPos := m.guardPos
	m.traverse()
	total1 = len(m.visited)

	// It's only possible to affect the path by placing an obstacle in a position that was visited. We're going to brute force but at lease narrow it down to those.
	candidates := []tuple{}
	for k := range m.visited {
		if k.x == initialGuardPos.x && k.y == initialGuardPos.y {
			continue
		}
		candidates = append(candidates, k)
	}

	for _, c := range candidates {
		m.wipe(initialGuardDir, initialGuardPos)
		m.grid[c.y][c.x] = "O"
		total2 += m.traverse()
		m.grid[c.y][c.x] = "."
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
	}

	return &labMap{
		grid:     grid,
		guardPos: startingGuardPos,
		guardDir: tuple{x: 0, y: -1},
		visited:  map[tuple]map[tuple]struct{}{},
	}
}

func (m *labMap) traverse() int {
	for {
		if dirs, ok := m.visited[m.guardPos]; ok {
			if _, dok := dirs[m.guardDir]; dok {
				// cycle, exit
				return 1
			}
		} else {
			m.visited[m.guardPos] = map[tuple]struct{}{}
		}
		m.visited[m.guardPos][m.guardDir] = struct{}{}

		next := tuple{
			x: m.guardPos.x + m.guardDir.x,
			y: m.guardPos.y + m.guardDir.y,
		}

		if next.x < 0 || next.x >= len(m.grid[0]) {
			return 0
		}

		if next.y < 0 || next.y >= len(m.grid) {
			return 0
		}

		if m.grid[next.y][next.x] == "#" || m.grid[next.y][next.x] == "O" {
			m.guardDir = turn(m.guardDir)
		} else {
			m.guardPos = next
		}

	}
}

func (m *labMap) wipe(dir, pos tuple) {
	m.visited = map[tuple]map[tuple]struct{}{}
	m.guardDir = dir
	m.guardPos = pos
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
	visited  map[tuple]map[tuple]struct{}
}
