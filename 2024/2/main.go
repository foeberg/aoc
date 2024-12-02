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
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var input []string
	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		input = append(input, string(b))
	}

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func isSafe(levels []string) (int, bool) {
	last := 0

	for i := 0; i < len(levels)-1; i++ {
		diff := strToInt(levels[i]) - strToInt(levels[i+1])
		if abs(diff) > 3 || diff == 0 {
			return i, false
		}

		if last == 0 {
			last = diff
			continue
		}

		if (last < 0 && diff > 0) || (last > 0 && diff < 0) {
			return i, false
		}
	}

	return -1, true
}

func part1(input []string) int {
	count := 0

	for _, line := range input {
		levels := strings.Fields(line)

		if _, s := isSafe(levels); s {
			count++
		}
	}

	return count
}

func part2(input []string) int {
	count := 0

	for _, line := range input {
		levels := strings.Fields(line)

		i, s := isSafe(levels)
		if s {
			count++
			continue
		}

		var s1, s2, s3 bool
		var l1, l2, l3 []string

		l1 = append(l1, levels[0:i]...)
		l1 = append(l1, levels[i+1:]...)
		_, s1 = isSafe(l1)

		l2 = append(l2, levels[0:i+1]...)
		l2 = append(l2, levels[i+2:]...)
		_, s2 = isSafe(l2)

		if i > 0 {
			l3 = append(l3, levels[0:i-1]...)
			l3 = append(l3, levels[i:]...)
			_, s3 = isSafe(l3)
		}

		if s1 || s2 || s3 {
			count++
		}

	}
	return count
}
