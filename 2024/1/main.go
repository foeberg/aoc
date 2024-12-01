package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
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

func part1(input []string) int {
	var left, right []string

	for _, line := range input {
		parts := strings.Fields(line)
		left = append(left, parts[0])
		right = append(right, parts[1])
	}

	slices.Sort(left)
	slices.Sort(right)

	var sum int
	for i := 0; i < len(left); i++ {
		l, _ := strconv.Atoi(left[i])
		r, _ := strconv.Atoi(right[i])

		diff := l - r
		if diff < 0 {
			diff *= -1
		}

		sum += diff
	}

	return sum
}

func part2(input []string) int {
	var left, right []string

	for _, line := range input {
		parts := strings.Fields(line)
		left = append(left, parts[0])
		right = append(right, parts[1])
	}

	cache := make(map[string]int)

	for i := 0; i < len(right); i++ {
		cache[right[i]]++
	}

	var sum int
	for i := 0; i < len(left); i++ {
		l, _ := strconv.Atoi(left[i])
		sum += l * cache[left[i]]
	}

	return sum
}
