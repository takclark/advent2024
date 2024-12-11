package main

import (
	"fmt"
	"strings"

	"github.com/takclark/advent2024/parsing"
)

func main() {
	// input := parsing.MustReadFile("input")
	input := parsing.MustReadFile("test")
	part1, part2 := solve(input)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}

type disk struct {
	data      []int64
	firstFree int   // first index of available free space
	freeSize  int   // size of above space
	file      int   // last index of file under consideration
	fileID    int64 // ID of file
	fileSize  int   // size of above
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

	for d.firstFree < d.file {
		d.data[d.firstFree], d.data[d.file] = d.data[d.file], d.data[d.firstFree]

		d.setFirstFree()
		d.setLastFile()
	}
}

func (d *disk) condenseDefrag() {
	// d.setFirstFree()
	// d.setLastFile()
	//
	// for d.fileID > 0 { // i.e. try all the files once, we know 0 can't move
	// 	fmt.Printf("%+v\n", d)
	// 	if d.fileSize <= d.freeSize {
	// 		for i := 0; i < d.fileSize; i++ {
	// 			d.data[d.firstFree+i], d.data[d.file-i] = d.data[d.file-i], d.data[d.firstFree+1]
	// 		}
	//
	// 		d.setFirstFree()
	// 		d.setLastFile()
	// 	} else {
	// 		d.skipFile()
	// 	}
	// }
	//
	fmt.Printf("%+v\n", d)
}

func (d *disk) setFirstFree() {
	var i, j int
	for i = d.firstFree; d.data[i] != -1; i++ {
	}
	for j = i; j < len(d.data) && d.data[j] == -1; j++ {
	}
	d.firstFree = i
	d.freeSize = j - i
}

func (d *disk) setLastFile() {
	var i, j int
	for i = d.file; d.data[i] == -1; i-- {
	}
	d.fileID = d.data[i]
	for j = i; j >= 0 && d.data[j] == d.fileID; j-- {
	}
	d.file = i
	d.fileSize = i - j
}

func (d *disk) skipFile() {
	d.file = d.file - d.fileSize
	d.setLastFile()
}

func solve(input string) (int64, int64) {
	var total1, total2 int64

	d := parseInput(input)
	d.condense()
	total1 = d.checksum()

	d = parseInput(input)
	d.condenseDefrag()
	total2 = d.checksum()

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
		data: data,
		file: len(data) - 1,
	}
}
