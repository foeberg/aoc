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

	var equations []string

	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		equations = append(equations, string(b))
	}

	fmt.Printf("Part 1: %d\n", part1(equations))
	fmt.Printf("Part 2: %d\n", part2(equations))
}

func extractNums(equation string) (int, []int) {
	split := strings.Split(equation, ":")
	testVal, _ := strconv.Atoi(split[0])

	split = strings.Split(split[1][1:], " ")

	var nums []int

	for _, s := range split {
		n, _ := strconv.Atoi(s)
		nums = append(nums, n)
	}

	return testVal, nums
}

func isValidEquation(val, current int, numsLeft []int) bool {
	if current > val {
		return false
	}
	if current < val && len(numsLeft) == 0 {
		return false
	}
	if current == val && len(numsLeft) == 0 {
		return true
	}

	return isValidEquation(val, current*numsLeft[0], numsLeft[1:]) ||
		isValidEquation(val, current+numsLeft[0], numsLeft[1:])
}

func part1(equations []string) int {
	result := 0

	for _, eq := range equations {
		val, nums := extractNums(eq)

		if isValidEquation(val, 0, nums) {
			result += val
		}
	}

	return result
}

func concat(v1, v2 int) int {
	str := strconv.Itoa(v1) + strconv.Itoa(v2)
	con, _ := strconv.Atoi(str)
	return con
}

func isValidEquation2(val, current int, numsLeft []int) bool {
	if current > val {
		return false
	}
	if current < val && len(numsLeft) == 0 {
		return false
	}
	if current == val && len(numsLeft) == 0 {
		return true
	}

	return isValidEquation2(val, current*numsLeft[0], numsLeft[1:]) ||
		isValidEquation2(val, current+numsLeft[0], numsLeft[1:]) ||
		isValidEquation2(val, concat(current, numsLeft[0]), numsLeft[1:])
}

func part2(equations []string) int {
	result := 0

	for _, eq := range equations {
		val, nums := extractNums(eq)

		if isValidEquation2(val, 0, nums) {
			result += val
		}
	}

	return result
}
