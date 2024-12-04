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
	lines := strings.Split(input, "\n")

	content := [][]string{}
	for _, v := range lines {
		if v == "" {
			continue
		}

		content = append(content, strings.Split(v, ""))
	}

	g := &grid{
		chars: content,
		word:  []string{"X", "M", "A", "S"},
	}

	var total1, total2 int
	for i, row := range g.chars {
		for j := range row {
			total1 += g.baseSearch(i, j) // count XMAS's starting here
		}
	}

	for i := range len(g.chars) - 2 {
		for j := range len(g.chars[i]) - 2 {
			total2 += g.findXShapesMASes(i, j)
		}
	}

	return total1, total2
}

type grid struct {
	chars [][]string
	word  []string
}

func (g *grid) baseSearch(i, j int) int {
	var sum int
	for di := -1; di < 2; di++ {
		for dj := -1; dj < 2; dj++ {
			sum += g.search(i, j, di, dj, 0)
		}
	}

	return sum
}

func (g *grid) search(i, j, di, dj, n int) int {
	if n >= len(g.word) {
		return 1
	}

	if i < 0 || i >= len(g.chars) || j < 0 || j >= len(g.chars) {
		return 0
	}

	if g.chars[i][j] != g.word[n] {
		return 0
	}

	return g.search(i+di, j+dj, di, dj, n+1)
}

func (g *grid) findXShapesMASes(i, j int) int {
	if g.chars[i+1][j+1] != "A" {
		return 0
	}

	if !((g.chars[i][j] == "M" && g.chars[i+2][j+2] == "S") || (g.chars[i][j] == "S" && g.chars[i+2][j+2] == "M")) {
		return 0
	}

	if !((g.chars[i+2][j] == "M" && g.chars[i][j+2] == "S") || (g.chars[i+2][j] == "S" && g.chars[i][j+2] == "M")) {
		return 0
	}

	return 1
}
