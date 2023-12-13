package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	patterns := parsePatterns()

	var sum int
	for _, pattern := range patterns {
		sum += summarizePattern(pattern)
	}

	fmt.Println(sum)
}

func summarizePattern(pattern []string) int {
	// Horizontal lines
	hDups := findHorizontalDuplicates(pattern)

	for _, dup := range hDups {
		isReflection := true

		for i := 1; i <= min(dup, len(pattern)-(dup+2)); i++ {
			if pattern[dup-i] != pattern[dup+1+i] {
				isReflection = false
				break
			}
		}

		if isReflection {
			// We found a line between rows dup & dup+1
			return (dup + 1) * 100
		}
	}

	// Vertical lines
	vDups := findVerticalDuplicates(pattern)

	for _, dup := range vDups {
		isReflection := true

		for i := 1; i <= min(dup, len(pattern[0])-(dup+2)); i++ {
			if getColumn(dup-i, pattern) != getColumn(dup+1+i, pattern) {
				isReflection = false
				break
			}
		}

		if isReflection {
			// We found a line between columns dup & dup+1
			return dup + 1
		}
	}

	return 0
}

func findHorizontalDuplicates(pattern []string) []int {
	var rows []int

	for i := 0; i < len(pattern)-1; i++ {
		if pattern[i] == pattern[i+1] {
			rows = append(rows, i)
		}
	}

	return rows
}

func findVerticalDuplicates(pattern []string) []int {
	var columns []int

	for i := 0; i < len(pattern[0])-1; i++ {
		if getColumn(i, pattern) == getColumn(i+1, pattern) {
			columns = append(columns, i)
		}
	}

	return columns
}

func getColumn(col int, pattern []string) string {
	var bs []byte
	for i := 0; i < len(pattern); i++ {
		bs = append(bs, byte(pattern[i][col]))
	}
	return string(bs)
}

func parsePatterns() [][]string {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)
	var patterns [][]string

	patterns = append(patterns, []string{})
	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if len(b) == 0 {
			patterns = append(patterns, []string{})
			continue
		}

		patterns[len(patterns)-1] = append(patterns[len(patterns)-1], string(b))
	}

	return patterns
}
