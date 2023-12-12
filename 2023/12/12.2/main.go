package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)
	var lines []string
	var groups [][]int
	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		split := strings.Split(string(b), " ")
		lines = append(lines, unfoldLine(split[0]))

		var g []int
		for i := 0; i < 5; i++ {
			for _, s := range strings.Split(split[1], ",") {
				n, _ := strconv.Atoi(s)
				g = append(g, n)
			}
		}

		groups = append(groups, g)
	}

	sum := 0
	for i := range lines {
		cache := initCache(len(lines[i]), len(groups[i]))
		sum += dp(0, 0, lines[i], groups[i], cache)
	}

	fmt.Println(sum)
}

func initCache(i, j int) [][]int {
	var c [][]int
	for n := 0; n < i; n++ {
		var d []int
		for m := 0; m < j+1; m++ {
			d = append(d, -1)
		}
		c = append(c, d)
	}
	return c
}

func unfoldLine(line string) string {
	b := []byte(line)
	for i := 0; i < 4; i++ {
		b = append(b, '?')
		b = append(b, line...)
	}

	return string(b)
}

func dp(i, j int, line string, groups []int, cache [][]int) int {
	if i >= len(line) {
		if j < len(groups) {
			return 0
		}
		return 1
	}

	if cache[i][j] != -1 {
		return cache[i][j]
	}

	var sum int

	if line[i] == '.' {
		sum = dp(i+1, j, line, groups, cache)
	} else {
		if line[i] == '?' {
			sum += dp(i+1, j, line, groups, cache) // Case where we don't put spring on ?
		}

		if j < len(groups) {
			// We're at a ? or #, check if you can fit groups[j] elements in a row
			run := 0

			for k := i; k < len(line); k++ {
				if run > groups[j] || line[k] == '.' || run == groups[j] && line[k] == '?' {
					break
				}
				run += 1
			}

			if run == groups[j] {
				sum += dp(min(len(line), i+run+1), j+1, line, groups, cache)
			}
		}
	}

	cache[i][j] = sum

	return sum
}
