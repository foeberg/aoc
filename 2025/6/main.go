package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input")

	var input []string
	for s := bufio.NewScanner(file); s.Scan(); {
		input = append(input, s.Text())
	}

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input []string) int {
	var worksheet [][]string
	for _, line := range input {
		worksheet = append(worksheet, strings.Fields(line))
	}

	var grandTotal int

	for col := range worksheet[0] {
		op := worksheet[len(worksheet)-1][col]

		var total int
		for row := 0; row < len(worksheet)-1; row++ {
			num, _ := strconv.Atoi(worksheet[row][col])

			switch op {
			case "*":
				total = max(1, total) * num
			case "+":
				total += num
			}
		}
		grandTotal += total
	}

	return grandTotal
}

func part2(input []string) int {
	var grandTotal int
	var col int

	for {
		op := input[len(input)-1][col]
		if op == ' ' {
			continue
		}
		var terms [][]byte

		for {
			terms = append(terms, []byte{})
			for row := 0; row < len(input)-1; row++ {
				if input[row][col] == ' ' {
					continue
				}

				terms[len(terms)-1] = append(terms[len(terms)-1], input[row][col])
			}

			latestTerm := terms[len(terms)-1]
			if len(latestTerm) == 0 {
				break
			}
			col++

			if col >= len(input[0]) {
				break
			}
		}

		var total int
		if op == '*' {
			total = 1
		}

		for _, term := range terms {
			if len(term) == 0 {
				continue
			}

			intTerm, _ := strconv.Atoi(string(term))
			if op == '*' {
				total *= intTerm
			}
			if op == '+' {
				total += intTerm
			}
		}

		grandTotal += total

		col++
		if col >= len(input[0]) {
			break
		}
	}

	return grandTotal
}
