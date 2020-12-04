package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

var lines = strings.Split(input, "\n")

func TestPart1(t *testing.T) {
	assert.Equal(t, 7, part1(lines, 3, 1))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 336, part2(lines))
}
