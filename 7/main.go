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

func solve(input string) (int64, int64) {
	eqs := parseInput(input)

	var total1, total2 int64
	s1 := solver{includeConcat: false}
	s2 := solver{includeConcat: true}
	for _, e := range eqs {
		if s1.couldWork(e) > 0 {
			total1 += e.result
		}
		if s2.couldWork(e) > 0 {
			total2 += e.result
		}
	}

	return total1, total2
}

type solver struct {
	includeConcat bool
}

type eq struct {
	result   int64
	operands []int64
}

// return the number of ways in which eq could be solved
func (s *solver) couldWork(e eq) int {
	ways := s.waysToSolve(e, 0, opPlus, 0) + s.waysToSolve(e, 1, opMul, 0)
	if s.includeConcat {
		ways += s.waysToSolve(e, 0, opCat, 0)
	}

	return ways
}

func (s *solver) waysToSolve(e eq, running int64, op, i int) int {
	// fmt.Println("ways to solve for:", e)
	// fmt.Printf("with running: %d, op: %d, i: %d\n", running, op, i)
	tally := do(running, e.operands[i], op)
	if i == len(e.operands)-1 {
		// base case - see if this solves it
		if tally == e.result {
			return 1
		}

		return 0
	}

	ways := s.waysToSolve(e, tally, opPlus, i+1) + s.waysToSolve(e, tally, opMul, i+1)
	if s.includeConcat {
		ways += s.waysToSolve(e, tally, opCat, i+1)
	}

	return ways
}

func do(m, n int64, op int) int64 {
	if op == opPlus {
		return m + n
	}
	if op == opMul {
		return m * n
	}

	// concat
	if m == 0 {
		return n
	}

	return parsing.MustParseInt64(fmt.Sprintf("%d%d", m, n))
}

func parseInput(input string) []eq {
	var parsed []eq
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		if l == "" {
			continue
		}
		split := strings.Split(l, ":")
		this := eq{}
		this.result = parsing.MustParseInt64(split[0])
		this.operands = parsing.SeparatedStringToInt64Slice(split[1], " ")
		parsed = append(parsed, this)

	}

	return parsed
}
