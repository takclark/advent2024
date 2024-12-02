package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/takclark/advent2024/parsing"
)

func main() {
	input := parsing.MustReadFile("input")
	// input := parsing.MustReadFile("test")
	part1, part2 := solve(input)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}

type report []int

const (
	DirIncreasing = 1
	DirDecreasing = -1
)

func (r report) isSafe() int {
	var dir int
	for i := range (len(r)) - 1 {
		if dir == 0 {
			if r[i+1] < r[i] {
				dir = DirDecreasing
			}
			if r[i+1] > r[i] {
				dir = DirIncreasing
			}
		}

		if r[i+1] == r[i] {
			return 0
		}

		if abs(r[i+1]-r[i]) > 3 {
			return 0
		}

		if (dir == DirIncreasing) && (r[i+1] < r[i]) {
			return 0
		}

		if (dir == DirDecreasing) && (r[i+1] > r[i]) {
			return 0
		}
	}

	return 1
}

func (r report) isSafeWithDampener() int {
	// brute force? they're not that big
	if r.isSafe() > 0 {
		return 1
	}

	for i := range len(r) {
		s := make([]int, len(r))
		copy(s, r)
		s = slices.Delete(s, i, i+1)
		if report(s).isSafe() > 0 {
			return 1
		}
	}

	return 0
}

func solve(input string) (int, int) {
	lines := strings.Split(input, "\n")

	var total1, total2 int
	for _, l := range lines {
		if l == "" {
			continue
		}
		r := report(parsing.SpaceSeparatedStringToIntSlice(l))
		total1 += r.isSafe()
		total2 += r.isSafeWithDampener()
	}

	return total1, total2
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
