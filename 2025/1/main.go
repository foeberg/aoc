package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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
	position := 50
	zeros := 0

	for _, instruction := range input {
		direction := instruction[0]
		num, _ := strconv.Atoi(string(instruction[1:]))

		switch direction {
		case 'L':
			position = (position - num) % 100
		default:
			position = (position + num) % 100
		}

		if position < 0 {
			position = 100 + position
		}

		if position == 0 {
			zeros++
		}
	}

	return zeros
}

func part2(input []string) int {
	position := 50
	zeros := 0

	for _, instruction := range input {
		direction := instruction[0]
		num, _ := strconv.Atoi(string(instruction[1:]))

		dir := 1
		if direction == 'L' {
			dir = -1
		}

		for i := num; i > 0; i-- {
			position += dir
			if position < 0 {
				position = 99
			}
			if position > 99 {
				position = 0
			}
			if position == 0 {
				zeros++
			}
		}
	}

	return zeros
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
