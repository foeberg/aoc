package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
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

var mul = regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

func part1(input []string) int {
	res := 0

	var merged string
	for _, line := range input {
		merged += line
	}

	matches := mul.FindAllStringSubmatch(merged, -1)
	for _, match := range matches {
		res += strToInt(match[1]) * strToInt(match[2])
	}

	return res
}

var dont = regexp.MustCompile(`don't\(\).*$`)

func part2(input []string) int {
	res := 0

	var merged string
	for _, line := range input {
		merged += line
	}

	var processed string

	split := strings.Split(merged, "do()")

	for _, s := range split {
		s = dont.ReplaceAllString(s, "")
		processed += s
	}

	matches := mul.FindAllStringSubmatch(processed, -1)
	for _, match := range matches {
		res += strToInt(match[1]) * strToInt(match[2])
	}

	return res
}
