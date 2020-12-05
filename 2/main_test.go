package main

import "testing"

func TestPart2(t *testing.T) {
	var tests = []struct {
		input  string
		expect bool
	}{
		{"1-3 a: abcde", true},
		{"1-3 b: cdefg", false},
		{"2-9 c: ccccccccc", false},
	}

	for _, test := range tests {
		parts := parseLine(test.input)
		res := part2(parts)
		if res != test.expect {
			t.Errorf("expected %v for input '%s' got %v", test.expect, test.input, res)
		}
	}
}
