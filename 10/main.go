package main

import (
	"fmt"
	"strings"

	"github.com/takclark/advent2024/parsing"
)

func main() {
	input := parsing.MustReadFile("input")
	part1, part2 := solve(input)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}

type trailMap struct {
	grid           [][]int
	reachablePeaks map[cord]map[cord]struct{}
}

type cord struct {
	x, y int
}

func (m trailMap) score(prev, x, y int, trailhead cord) int64 {
	// fmt.Println("score", prev, x, y)
	if x < 0 || x >= len(m.grid[0]) {
		return 0
	}
	if y < 0 || y >= len(m.grid) {
		return 0
	}

	if m.grid[y][x] != prev+1 {
		return 0
	}

	if m.grid[y][x] == 9 {
		if _, ok := m.reachablePeaks[trailhead]; ok {
			m.reachablePeaks[trailhead][cord{x, y}] = struct{}{}
		} else {
			m.reachablePeaks[trailhead] = map[cord]struct{}{{x, y}: {}}
		}
		return 1
	}

	return m.score(m.grid[y][x], x-1, y, trailhead) +
		m.score(m.grid[y][x], x+1, y, trailhead) +
		m.score(m.grid[y][x], x, y-1, trailhead) +
		m.score(m.grid[y][x], x, y+1, trailhead)
}

func solve(input string) (int64, int64) {
	var total1, total2 int64

	m := parseInput(input)
	for y, row := range m.grid {
		for x, l := range row {
			if l != 0 {
				continue
			}
			total2 += m.score(-1, x, y, cord{x, y})
		}
	}

	for _, peaks := range m.reachablePeaks {
		total1 += int64(len(peaks))
	}

	return total1, total2
}

func parseInput(input string) trailMap {
	m := trailMap{
		grid:           [][]int{},
		reachablePeaks: make(map[cord]map[cord]struct{}),
	}
	lines := strings.Split(input, "\n")

	for _, l := range lines {
		m.grid = append(m.grid, parsing.SeparatedStringToIntSlice(l, ""))
	}

	return m
}
