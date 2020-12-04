package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("./d2/input.txt")
	defer f.Close()

	bts, _ := ioutil.ReadFile("./d3/input.txt")
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
		if (input[i][posX%nLen]) == '#' {
			result++
		}
		posX += dX
	}

	return result
}

func part2(input []string) int {
	var results []int
	for _, slope := range [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}} {
		results = append(results, part1(input, slope[0], slope[1]))
	}

	answer := results[0]
	for i := 1; i < len(results); i++ {
		answer *= results[i]
	}

	return answer
}
