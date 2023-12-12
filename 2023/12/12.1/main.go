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
		lines = append(lines, split[0])

		var g []int
		for _, s := range strings.Split(split[1], ",") {
			n, _ := strconv.Atoi(s)
			g = append(g, n)
		}

		groups = append(groups, g)
	}

	sum := 0
	for i := range lines {
		sum += dp(0, 0, lines[i], groups[i])
	}

	fmt.Println(sum)
}

func dp(i, j int, line string, groups []int) int {
	if i >= len(line) {
		if j < len(groups) {
			return 0
		}
		return 1
	}

	var sum int

	if line[i] == '.' {
		return dp(i+1, j, line, groups)
	}

	if line[i] == '?' {
		sum += dp(i+1, j, line, groups) // Case where we don't put spring on ?
	}

	if j < len(groups) {
		// We're at a ? or #, check if you can fit groups[j] elements in a row
		run := 0

		for k := i; k < len(line); k++ {
			if run > groups[j] {
				return sum
			}

			if line[k] == '.' || run == groups[j] && line[k] == '?' {
				break
			}
			run += 1
		}

		if run == groups[j] {
			sum += dp(min(len(line), i+run+1), j+1, line, groups)
		}
	}

	return sum
}
