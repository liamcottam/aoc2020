package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, part1("FBFBBFFRLR"), 357)
	assert.Equal(t, part1("BFFFBBFRRR"), 567)
	assert.Equal(t, part1("FFFBBBFRRR"), 119)
	assert.Equal(t, part1("BBFFBBFRLL"), 820)
}
