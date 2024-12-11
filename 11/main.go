package main

import (
	"fmt"
	"math"

	"github.com/takclark/advent2024/parsing"
)

const (
	testInput = "125 17"
	realInput = "773 79858 0 71 213357 2937 1 3998391"
)

func main() {
	input := realInput
	part1, part2 := solve(input)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}

type pebbles []int64

func (p pebbles) tick() pebbles {
	result := []int64{}
	for i := 0; i < len(p); i++ {
		new := next(p[i])
		result = append(result, new...)
	}

	return result
}

// map key for memoization
type pair struct {
	n int64
	t int64
}

type solver struct {
	ps   pebbles
	memo map[pair]int64
}

func (s solver) lenAfter(n, ticks int64) int64 {
	return 0
}

func next(n int64) []int64 {
	if n == 0 {
		return []int64{1}
	}

	l := math.Log10(float64(n))
	digits := int64(math.Floor(l + 1))
	if digits != 0 && digits%2 == 0 {
		pow := int64(math.Pow(10, float64(digits/2)))
		left := n / pow
		right := n - left*pow
		return []int64{left, right}
	}

	return []int64{n * 2024}
}

func solve(input string) (int64, int64) {
	ps := pebbles(parsing.SeparatedStringToInt64Slice(input, " "))
	var total1, total2 int64

	fmt.Println(ps)
	for range 25 {
		ps = ps.tick()
	}

	total1 = int64(len(ps))

	return total1, total2
}
