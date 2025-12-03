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

	var input string
	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		input = string(b)
	}

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input string) int {
	seqs := strings.Split(input, ",")

	var sum int
	for _, seq := range seqs {
		exp := expandSeq(seq)
		for _, num := range exp {
			sum += findDoubleSequence(num)
		}
	}

	return sum
}

func part2(input string) int {
	seqs := strings.Split(input, ",")

	var sum int
	for _, seq := range seqs {
		exp := expandSeq(seq)
		for _, num := range exp {
			sum += findSequences(num)
		}
	}

	return sum
}

func expandSeq(seq string) []string {
	lims := strings.Split(seq, "-")
	start, _ := strconv.Atoi(lims[0])
	end, _ := strconv.Atoi(lims[1])

	var res []string
	for i := start; i <= end; i++ {
		res = append(res, fmt.Sprint(i))
	}

	return res
}

func findDoubleSequence(num string) int {
	if len(num)%2 != 0 {
		return 0
	}

	if num[:len(num)/2] == num[len(num)/2:] {
		intNum, _ := strconv.Atoi(num)
		return intNum
	}

	return 0
}

func findSequences(num string) int {
	for i := 1; i < len(num); i++ {
		if num[i] == num[0] {
			if len(num[i:]) < i {
				return 0
			}

			if num[0:i] == num[i:i+len(num[0:i])] {
				// Found one repearing sequence
				// Check if rest of num is the same sequence
				seq := num[0:i]
				remainder := num[i+len(num[0:i]):]

				if len(remainder) == 0 {
					intNum, _ := strconv.Atoi(num)
					return intNum
				}

				occurrences := strings.Count(remainder, seq)

				if occurrences == 0 {
					return 0
				}

				if len(remainder)%occurrences == 0 && len(remainder)/occurrences == len(seq) {
					intNum, _ := strconv.Atoi(num)
					return intNum
				}
			}
		}
	}

	return 0
}
