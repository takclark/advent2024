package main

import (
	"fmt"
	"regexp"

	"github.com/takclark/advent2024/parsing"
)

func main() {
	input := parsing.MustReadFile("input")
	// input := parsing.MustReadFile("test2")
	part1, part2 := solve(input)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}

func solve(input string) (int64, int64) {
	exp1 := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	muls := exp1.FindAllString(input, -1)

	var m, n, total1, total2 int64
	for _, v := range muls {
		fmt.Sscanf(v, "mul(%d,%d)", &m, &n)
		total1 += m * n
	}

	exp2 := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don't\(\))`)
	instructions := exp2.FindAllString(input, -1)

	do := true
	for _, v := range instructions {
		if v == "do()" {
			do = true
			continue
		}

		if v == "don't()" {
			do = false
			continue
		}

		if !do {
			continue
		}

		fmt.Sscanf(v, "mul(%d,%d)", &m, &n)
		total2 += m * n
	}

	return total1, total2
}
