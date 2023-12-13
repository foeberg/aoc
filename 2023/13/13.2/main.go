package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Confusing name
type duplicate struct {
	index      int
	withSmudge bool
}

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
		diffs := 0

		for i := 1; i <= min(dup.index, len(pattern)-(dup.index+2)); i++ {
			if dup.withSmudge {
				diffs += getDiffs(pattern[dup.index-i], pattern[dup.index+1+i])
				if diffs > 0 {
					break
				}
			} else {
				diffs += getDiffs(pattern[dup.index-i], pattern[dup.index+1+i])
				if diffs > 1 {
					break
				}
			}
		}

		if dup.withSmudge && diffs == 0 || !dup.withSmudge && diffs == 1 {
			// We found a line between rows dup & dup+1
			return (dup.index + 1) * 100
		}
	}

	// Vertical lines
	vDups := findVerticalDuplicates(pattern)

	for _, dup := range vDups {
		diffs := 0

		for i := 1; i <= min(dup.index, len(pattern[0])-(dup.index+2)); i++ {
			if dup.withSmudge {
				diffs += getDiffs(getColumn(dup.index-i, pattern), getColumn(dup.index+1+i, pattern))
				if diffs > 0 {
					break
				}
			} else {
				diffs += getDiffs(getColumn(dup.index-i, pattern), getColumn(dup.index+1+i, pattern))
				if diffs > 1 {
					break
				}
			}
		}

		if dup.withSmudge && diffs == 0 || !dup.withSmudge && diffs == 1 {
			// We found a line between columns dup & dup+1
			return dup.index + 1
		}
	}

	return 0
}

func findHorizontalDuplicates(pattern []string) []duplicate {
	var rows []duplicate

	for i := 0; i < len(pattern)-1; i++ {
		diffs := getDiffs(pattern[i], pattern[i+1])

		if diffs == 1 {
			rows = append(rows, duplicate{i, true})
		}
		if diffs == 0 {
			rows = append(rows, duplicate{i, false})
		}
	}

	return rows
}

func findVerticalDuplicates(pattern []string) []duplicate {
	var columns []duplicate

	for i := 0; i < len(pattern[0])-1; i++ {
		diffs := getDiffs(getColumn(i, pattern), getColumn(i+1, pattern))

		if diffs == 1 {
			columns = append(columns, duplicate{i, true})
		}
		if diffs == 0 {
			columns = append(columns, duplicate{i, false})
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

func getDiffs(a, b string) int {
	diffs := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diffs++

			if diffs > 1 {
				break
			}
		}
	}

	return diffs
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
