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
	memo map[pair]int64
}

func (s solver) solve(ps []int64, ticks int64) int64 {
	var total int64
	for _, n := range ps {
		total += s.lenAfter(n, ticks)
	}

	return total
}

func (s *solver) lenAfter(n, ticks int64) int64 {
	if ticks == 0 {
		return 1
	}

	if v, ok := s.memo[pair{n, ticks}]; ok {
		return v
	}

	res := next(n)
	var total int64
	for _, v := range res {
		lenHere := s.lenAfter(v, ticks-1)
		s.memo[pair{v, ticks - 1}] = lenHere
		total += lenHere
	}

	return total
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

	s := solver{memo: make(map[pair]int64)}
	total1 = s.solve(ps, 25)
	total2 = s.solve(ps, 75)

	return total1, total2
}
