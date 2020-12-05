package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	bts, _ := ioutil.ReadFile("./3/input.txt")
	lines := strings.Split(string(bts), "\n")
	fmt.Println(part1(lines, 3, 1))
	fmt.Println(part2(lines))
}

func part1(input []string, dX, dY int) int {
	var (
		nLen   = len(input[0])
		nRows  = len(input)
		posX   = dX
		result int
	)

	for i := dY; i < nRows; i += dY {
		if input[i][posX%nLen] == '#' {
			result++
		}
		posX += dX
	}

	return result
}

func part2(input []string) int {
	result := part1(input, 1, 1)
	for _, slope := range [][]int{{3, 1}, {5, 1}, {7, 1}, {1, 2}} {
		result *= part1(input, slope[0], slope[1])
	}

	return result
}
