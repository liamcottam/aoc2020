package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	bts, _ := ioutil.ReadFile("./d4/input.txt")
	lines := strings.Split(string(bts), "\n\n")

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

var expect = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func part1(lines []string) int {
	var result int
main:
	for _, line := range lines {
		line = strings.ReplaceAll(line, "\n", " ")
		parts := strings.Split(line, " ")
		var m = make(map[string]bool)

		for _, part := range parts {
			parts2 := strings.Split(part, ":")
			m[parts2[0]] = true
		}

		for _, expect := range expect {
			if _, ok := m[expect]; !ok {
				continue main
			}
		}
		result++
	}

	return result
}

func part2(lines []string) int {
	var result int

	for _, line := range lines {
		line = strings.ReplaceAll(line, "\n", " ")
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")

		if isValid(parts) {
			result++
		}
	}

	return result
}

func isValid(lines []string) (result bool) {

	m := map[string]bool{}

	result = false
	for _, line := range lines {
		parts := strings.Split(line, ":")

		switch parts[0] {
		case "byr":
			if len(parts[1]) != 4 {
				return
			}
			i := mustParseInt(parts[1])
			if i < 1920 || i > 2002 {
				return
			}
		case "iyr":
			if len(parts[1]) != 4 {
				return
			}
			i := mustParseInt(parts[1])
			if i < 2010 || i > 2020 {
				return
			}
		case "eyr":
			if len(parts[1]) != 4 {
				return
			}
			i := mustParseInt(parts[1])
			if i < 2020 || i > 2030 {
				return
			}
		case "hgt":
			if strings.HasSuffix(parts[1], "cm") {
				i := mustParseInt(parts[1][:len(parts[1])-2])
				if i < 150 || i > 193 {
					return
				}
			} else if strings.HasSuffix(parts[1], "in") {
				i := mustParseInt(parts[1][:len(parts[1])-2])
				if i < 59 || i > 76 {
					return
				}
			} else {
				return
			}
		case "hcl":
			re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
			if !re.MatchString(parts[1]) {
				return
			}
		case "ecl":
			switch parts[1] {
			case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			default:
				return
			}
		case "pid":
			re := regexp.MustCompile(`^[0-9]{9}$`)
			if !re.MatchString(parts[1]) {
				return
			}
		}
		m[parts[0]] = true
	}

	for _, expect := range expect {
		if _, ok := m[expect]; !ok {
			return
		}
	}

	return true
}

func mustParseInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return out
}
