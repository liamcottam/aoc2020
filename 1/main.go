package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {

	f, err := os.Open("./1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var (
		inputs []int
		n      = 0
	)
	reader := bufio.NewScanner(f)
	for reader.Scan() {
		i, _ := strconv.Atoi(reader.Text())
		inputs = append(inputs, i)
		n++
	}

	start := time.Now()
	sort.Ints(inputs)

	var (
		i, j, k = 0, 1, 2
		target  = 2020
	)

	for ; j < n; j++ {
		if inputs[i]+inputs[j] == target {
			fmt.Println(inputs[i] * inputs[j])
			break
		}
		if j == n-1 && i < n {
			i++
			j = i
		}
	}

	i, j = 0, 1
	// This is probably wrong but I got my answer
	for {
		if inputs[i]+inputs[j]+inputs[k] == target {
			fmt.Println(inputs[i] * inputs[j] * inputs[k])
			break
		}

		k++
		if k == n {
			j++
			k = j + 1
		}
		if j == n-1 {
			i++
			j = i + 1
			k = j + 1
		}
		if i == n-2 {
			break
		}
	}
	fmt.Println(time.Since(start))
}
