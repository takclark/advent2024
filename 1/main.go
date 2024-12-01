package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	inBytes, _ := os.ReadFile("./input")
	input := string(inBytes)
	fmt.Println(solve01_01(input))
	fmt.Println(solve01_02(input))
}

func solve01_01(input string) int64 {
	var a, b []int64

	lines := strings.Split(input, "\n")
	for _, l := range lines {
		var m, n int64
		fmt.Sscanf(l, "%d  %d", &m, &n)
		a = append(a, m)
		b = append(b, n)
	}

	slices.Sort(a)
	slices.Sort(b)

	var total int64
	for i := range a {
		diff := b[i] - a[i]

		// pls add int abs Go
		if diff < 0 {
			diff = -diff
		}

		total += diff
	}

	return total
}

func solve01_02(input string) int64 {
	var a, b []int64

	lines := strings.Split(input, "\n")
	for _, l := range lines {
		var m, n int64
		fmt.Sscanf(l, "%d  %d", &m, &n)
		a = append(a, m)
		b = append(b, n)
	}

	m := map[int64]int64{}
	for _, v := range b {
		m[v]++
	}

	var score int64
	for _, v := range a {
		score += v * m[v]
	}

	return score
}
