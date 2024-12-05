package parsing

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

func MustReadFile(filename string) string {
	inBytes, _ := os.ReadFile(filename)
	return string(inBytes)
}

func MustParse(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func SeparatedStringToIntSlice(s, sep string) []int {
	nums := strings.Split(s, sep)
	nums = slices.DeleteFunc(nums, func(t string) bool {
		return t == ""
	})

	a := []int{}
	for _, v := range nums {
		a = append(a, MustParse(v))
	}

	return a
}
