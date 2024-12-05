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

type rule struct {
	before int
	after  int
}

type updater struct {
	rules      []rule
	priorities map[int]int
}

func solve(input string) (int, int) {
	rules, prints := parseInput(input)

	u := &updater{rules: rules}
	u.populatePriorities()
	fmt.Printf("%+v\n", u)
	var total1, total2, n int
	for _, p := range prints {
		eval := u.evaluate(p)
		total1 += u.evaluate(p)
		n += u.naiveEvaluate(p)

		if eval == 0 {
			ordered := u.sort(p)
			total2 += ordered[(len(ordered) / 2)]
		}

	}

	fmt.Println("N:", n)

	return total1, total2
}

// return the middle integer if print is valid according to rules. otherwise return 0.
func (u *updater) evaluate(print []int) int {
	for i := 1; i < len(print); i++ {
		for j := 0; j < i; j++ {
			if u.priorities[print[i]] < u.priorities[print[j]] {
				fmt.Printf("not allowed to print %d before %d! %v is invalid.\n", print[j], print[i], print)
				return 0
			}
		}
	}

	return print[(len(print) / 2)]
}

// return the middle integer if print is valid according to rules. otherwise return 0.
func (u *updater) naiveEvaluate(print []int) int {
	for i := 1; i < len(print); i++ {
		for j := 0; j < i; j++ {
			if !u.allow(print[j], print[i]) {
				fmt.Printf("NAIVE: not allowed to print %d before %d! %v is invalid.\n", print[j], print[i], print)
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

func (u *updater) populatePriorities() {
	p := make(map[int]int)

	for _, r := range u.rules {
		if _, ok := p[r.before]; !ok {
			p[r.before] = 0
		}

		if p[r.after] > p[r.before] {
			continue
		}

		p[r.after] = p[r.before] + 1
	}

	u.priorities = p
}

func (u *updater) sort(print []int) []int {
	slices.SortFunc(print, func(a, b int) int {
		return -1
	})

	return print
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
			// still parsing rules
			r := rule{}
			fmt.Sscanf(l, "%d|%d", &r.before, &r.after)
			rules = append(rules, r)
			continue
		}

		prints = append(prints, parsing.SeparatedStringToIntSlice(l, ","))
	}

	return rules, prints
}
