package main

import (
	"fmt"
	"strings"

	"github.com/takclark/advent2024/parsing"
)

const (
	opPlus = iota
	opMul
	opCat
)

func main() {
	input := parsing.MustReadFile("input")
	// input := parsing.MustReadFile("test")
	part1, part2 := solve(input)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}

func solve(input string) (int, int) {
	m := parseInput(input)

	var total1, total2 int
	affected1 := map[cord]struct{}{}
	affected2 := map[cord]struct{}{}
	for y := range m.grid {
		for x := range m.grid[y] {
			here := cord([2]int{x, y})
			for freq := range m.antennae {
				locs := m.antennae[freq]

				for i := range locs {
					// the forbidden fifth nested for loop
					for j := i + 1; j < len(locs); j++ {
						if !here.colinear(locs[i], locs[j]) {
							continue
						}

						affected2[here] = struct{}{}

						if (here.dist(locs[i]) == here.dist(locs[j])*2) || (here.dist(locs[j]) == here.dist(locs[i])*2) {
							affected1[here] = struct{}{}
						}
					}
				}
			}
		}
	}

	total1 = len(affected1)
	total2 = len(affected2)

	return total1, total2
}

type cord [2]int

func (c cord) dist(d cord) int {
	// looks like we're doing "city block" and not euclidean
	return abs(c[0]-d[0]) + abs(c[1]-d[1])
}

func (c cord) colinear(d, e cord) bool {
	// i.e. area of triangle formed by three points is 0.
	return c[0]*(d[1]-e[1])+d[0]*(e[1]-c[1])+e[0]*(c[1]-d[1]) == 0
}

type antennaMap struct {
	grid     [][]string
	antennae map[string][]cord
}

func parseInput(input string) antennaMap {
	var grid [][]string
	antennae := make(map[string][]cord)
	lines := strings.Split(input, "\n")
	for i, l := range lines {
		if l == "" {
			continue
		}

		row := strings.Split(l, "")
		for j, v := range row {
			if v == "." {
				continue
			}

			loc := cord{}
			loc[0] = j
			loc[1] = i

			if _, ok := antennae[v]; ok {
				antennae[v] = append(antennae[v], loc)
				continue
			}
			antennae[v] = []cord{loc}
		}
		grid = append(grid, row)
	}

	return antennaMap{
		grid:     grid,
		antennae: antennae,
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
