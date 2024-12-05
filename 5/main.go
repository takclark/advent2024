package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/takclark/advent2024/parsing"
)

func main() {
	input := parsing.MustReadFile("input")
	// input := parsing.MustReadFile("test")
	part1, part2 := solve(input)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}

type rule struct {
	before int
	after  int
}

type updater struct {
	rules []rule
}

func solve(input string) (int, int) {
	rules, prints := parseInput(input)

	u := &updater{rules: rules}
	var total1, total2 int
	for _, p := range prints {
		eval := u.evaluate(p)
		total1 += eval

		if eval == 0 {
			u.sort(p)
			total2 += p[(len(p) / 2)]
		}

	}

	return total1, total2
}

// return the middle integer if print is valid according to rules. otherwise return 0.
func (u *updater) evaluate(print []int) int {
	for i := 1; i < len(print); i++ {
		for j := 0; j < i; j++ {
			if !u.allow(print[j], print[i]) {
				return 0
			}
		}
	}

	return print[(len(print) / 2)]
}

func (u *updater) allow(b, a int) bool {
	for _, r := range u.rules {
		if r.after == b && r.before == a {
			return false
		}
	}

	return true
}

func (u *updater) sort(print []int) {
	for range len(print) { // bless me elves for I have sinned
		sort.Slice(print, func(a, b int) bool {
			return u.allow(print[a], print[b])
		})
	}
}

func parseInput(input string) ([]rule, [][]int) {
	lines := strings.Split(input, "\n")

	var rules []rule
	var prints [][]int
	var part int
	for _, l := range lines {
		if l == "" {
			part++
			continue
		}

		if part == 0 {
			r := rule{}
			fmt.Sscanf(l, "%d|%d", &r.before, &r.after)
			rules = append(rules, r)
			continue
		}

		prints = append(prints, parsing.SeparatedStringToIntSlice(l, ","))
	}

	return rules, prints
}
