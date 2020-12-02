package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`([0-9]+)-([0-9]+)\s(\w):\s(\w+)`)

func parseLine(input string) [][]string {
	res := re.FindAllStringSubmatch(input, -1)
	if len(res) != 1 || len(res[0]) != 5 {
		panic("invalid spec")
	}
	return res
}

func main() {
	f, _ := os.Open("./d2/input.txt")
	defer f.Close()

	var (
		p1valid int
		p2valid int
	)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		res := parseLine(scanner.Text())
		if part1(res) {
			p1valid++
		}
		if part2(res) {
			p2valid++
		}
	}

	fmt.Println(p1valid, p2valid)
}

func part1(matches [][]string) bool {
	var (
		min   int
		max   int
		found int
		c     rune
	)

	min, _ = strconv.Atoi(matches[0][1])
	max, _ = strconv.Atoi(matches[0][2])
	c = rune(matches[0][3][0])

	for _, r := range matches[0][4] {
		if r == c {
			found++
		}
	}
	return found >= min && found <= max
}

func part2(matches [][]string) bool {
	var (
		pos1 int
		pos2 int
		c    rune
	)

	pos1, _ = strconv.Atoi(matches[0][1])
	pos2, _ = strconv.Atoi(matches[0][2])
	c = rune(matches[0][3][0])
	pos1--
	pos2--

	var (
		p1match int
		p2match int
	)

	if rune(matches[0][4][pos1]) == c {
		p1match = 1
	}
	if rune(matches[0][4][pos2]) == c {
		p2match = 1
	}
	return (p1match ^ p2match) != 0
}
