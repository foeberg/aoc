package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input")

	var input []string
	for s := bufio.NewScanner(file); s.Scan(); {
		input = append(input, s.Text())
	}

	i := slices.Index(input, "")
	db := input[:i]
	ingredients := input[i+1:]

	fmt.Printf("Part 1: %v\n", part1(db, ingredients))
	fmt.Printf("Part 2: %v\n", part2(db))
}

func part1(db, ingredients []string) int {
	ranges := parseRanges(db)

	var fresh int
	for _, ingredient := range ingredients {
		ing, _ := strconv.Atoi(ingredient)

		for _, rang := range ranges {
			if ing >= rang[0] && ing <= rang[1] {
				fresh++
				break
			}
		}
	}

	return fresh
}

func part2(db []string) int {
	ranges := parseRanges(db)
	ranges = reduceRanges(ranges)

	var count int

	for _, rang := range ranges {
		count += rang[1] - rang[0] + 1
	}

	return count
}

func reduceRanges(ranges [][]int) [][]int {
	slices.SortFunc(ranges, func(a, b []int) int {
		if a[0] < b[0] {
			return -1
		}
		return 0
	})

	var newRanges [][]int
	currStart := ranges[0][0]
	currEnd := ranges[0][1]
	i := 0
	for {
		i++
		if i >= len(ranges) {
			newRanges = append(newRanges, []int{currStart, currEnd})
			break
		}

		// Not overlapping
		if ranges[i][0] > currEnd {
			newRanges = append(newRanges, []int{currStart, currEnd})
			currStart = ranges[i][0]
			currEnd = ranges[i][1]
			continue
		}

		if ranges[i][1] > currEnd {
			currEnd = ranges[i][1]
		}
	}

	return newRanges
}

func parseRanges(db []string) [][]int {
	var ranges [][]int
	for _, row := range db {
		split := strings.Split(row, "-")
		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])

		ranges = append(ranges, []int{start, end})
	}

	return ranges
}
