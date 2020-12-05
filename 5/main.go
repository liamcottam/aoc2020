package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	f, _ := os.Open("./5/input.txt")
	defer f.Close()

	var highest = -1

	scanner := bufio.NewScanner(f)

	var wat []int
	for scanner.Scan() {
		res := part1(scanner.Text())
		wat = append(wat, res)
		if res > highest || highest == -1 {
			highest = res
		}
	}
	fmt.Println(highest)

	sort.Ints(wat)
	var prev = wat[0]

	for _, v := range wat[1:] {
		if v != prev+1 {
			fmt.Println(v - 1)
			break
		}
		prev = v
	}
}

func part1(input string) int {
	var (
		l   = 0
		r   = 127
		row = 0
		col = 0
	)

loop:
	for _, c := range input {
		d := (r - l)
		switch c {
		case 'F':
			r -= (d / 2) + 1
		case 'B':
			l += (d / 2) + 1
		case 'L':
			row = l
			break loop
		case 'R':
			row = r
			break loop
		}
	}

	l = 0
	r = 7
	n := len(input[len(input)-3:]) - 1

	for idx, c := range input[len(input)-3:] {
		d := (r - l)
		switch c {
		case 'F', 'L':
			if idx == n {
				col = l
				break
			}
			r -= (d / 2) + 1
		case 'B', 'R':
			if idx == n {
				col = r
				break
			}
			l += (d / 2) + 1
		}
	}

	return row*8 + col
}

func part2() {}
