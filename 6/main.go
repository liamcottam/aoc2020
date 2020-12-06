package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		m      = make(map[rune]int)
		n      int
		result int
		unique int
	)

	parse("./6/input.txt", func(str string, reset bool) {
		if reset {
			for _, v := range m {
				if v == n {
					unique++
				}
			}
			m = make(map[rune]int)
			n = 0
			return
		}
		for _, r := range str {
			if _, ok := m[r]; ok {
				m[r]++
			} else {
				m[r] = 1
				result++
			}
		}
		n++
	})

	fmt.Println(result, unique)
}

func parse(filename string, fu func(string, bool)) {
	f, err := os.Open("./6/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	reset := false
	for scanner.Scan() {
		if reset {
			reset = false
		}
		str := scanner.Text()

		if len(str) == 0 {
			reset = true
		}
		fu(str, reset)
	}

	fu("", true)
}
