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

type disk struct {
	data      []int64
	firstFree int
	lastFile  int
}

func (d *disk) checksum() int64 {
	var sum int64
	for i, v := range d.data {
		if v == -1 {
			break
		}

		sum += int64(i) * v
	}

	return sum
}

func (d *disk) condense() {
	d.setFirstFree()
	d.setLastFile()

	for d.firstFree < d.lastFile {
		d.data[d.firstFree], d.data[d.lastFile] = d.data[d.lastFile], d.data[d.firstFree]

		d.setFirstFree()
		d.setLastFile()
	}
}

func (d *disk) condenseDefrag() {
}

func (d *disk) setFirstFree() {
	var i int
	for i = d.firstFree; d.data[i] != -1; i++ {
	}
	d.firstFree = i
}

func (d *disk) setLastFile() {
	var i int
	for i = d.lastFile; d.data[i] == -1; i-- {
	}
	d.lastFile = i
}

func solve(input string) (int64, int64) {
	var total1, total2 int64

	d := parseInput(input)
	d.condense()
	total1 = d.checksum()

	return total1, total2
}

func parseInput(input string) disk {
	stringDiskMap := strings.Split(input, "")
	var data []int64
	var fileNdx int64
	for i := 0; i < len(stringDiskMap); i += 2 {
		file := parsing.MustParseInt64(stringDiskMap[i])

		var free int64
		if i+1 < len(stringDiskMap) {
			free = parsing.MustParseInt64(stringDiskMap[i+1])
		} else {
			free = 0
		}

		for range file {
			data = append(data, fileNdx)
		}
		for range free {
			data = append(data, -1) // we'll use -1 to represent free space
		}

		fileNdx++
	}

	return disk{
		data:     data,
		lastFile: len(data) - 1,
	}
}
