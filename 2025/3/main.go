package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var banks [][]int
	var row int

	banks = append(banks, []int{})
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if b == '\n' {
			row++
			banks = append(banks, []int{})
			continue
		}

		banks[row] = append(banks[row], int(b)-48)
	}

	fmt.Printf("Part 1: %v\n", part1(banks))
	fmt.Printf("Part 2: %v\n", part2(banks))
}

func part1(banks [][]int) int {
	var sum int

	for i := range banks {
		sum += calculateJoltage(banks[i])
	}

	return sum
}

func calculateJoltage(bank []int) int {
	var firstIndex int

	for i := 0; i < len(bank)-1; i++ {
		if bank[i] > bank[firstIndex] {
			firstIndex = i
		}
	}

	secondIndex := firstIndex + 1
	for i := secondIndex; i < len(bank); i++ {
		if bank[i] > bank[secondIndex] {
			secondIndex = i
		}
	}

	return bank[firstIndex]*10 + bank[secondIndex]
}

func part2(banks [][]int) int {
	var sum int

	for i := range banks {
		sum += calculateJoltage2(banks[i], 12)
	}

	return sum
}

func calculateJoltage2(bank []int, count int) int {
	var biggestIndex int
	for i := 0; i < len(bank)-count+1; i++ {
		if bank[i] > bank[biggestIndex] {
			biggestIndex = i
		}
	}

	count--
	if count == 0 {
		return bank[biggestIndex]
	}

	return pow(10, count)*bank[biggestIndex] + calculateJoltage2(bank[biggestIndex+1:], count)
}

func pow(base, exp int) int {
	if exp == 0 {
		return 0
	}

	res := base
	for i := 1; i < exp; i++ {
		res *= base
	}
	return res
}
