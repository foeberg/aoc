package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var (
		lines []string
		sum   int
	)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		lines = append(lines, string(line))
	}

	for r, line := range lines {
		for c, b := range line {
			if isNum(byte(b)) || b == 46 {
				continue
			}

			sum += findAdjacentSum(lines, r, c)
		}
	}

	fmt.Println(sum)
}

func isNum(b byte) bool {
	if b < 59 && b > 47 {
		return true
	}

	return false
}

func findAdjacentSum(lines []string, row, column int) int {
	var sum int

	// check left
	n, _ := strconv.Atoi(string(findNum(lines, row, column-1, -1)))
	sum += n

	// check right
	n, _ = strconv.Atoi(string(findNum(lines, row, column+1, 1)))
	sum += n

	if row > 0 {
		// check top
		n, _ = strconv.Atoi(string(findNum(lines, row-1, column, 0)))
		sum += n

		n, _ = strconv.Atoi(string(findNum(lines, row-1, column-1, 0)))
		sum += n

		n, _ = strconv.Atoi(string(findNum(lines, row-1, column+1, 0)))
		sum += n
	}

	if row < len(lines)-1 {
		// check bottom
		n, _ = strconv.Atoi(string(findNum(lines, row+1, column, 0)))
		sum += n

		n, _ = strconv.Atoi(string(findNum(lines, row+1, column-1, 0)))
		sum += n

		n, _ = strconv.Atoi(string(findNum(lines, row+1, column+1, 0)))
		sum += n
	}

	return sum
}

func findNum(rows []string, row, column int, direction int) []byte {
	var b []byte

	line := rows[row]

	if column < 0 || column >= len(line) || !isNum(line[column]) {
		return b
	}

	b = append(b, line[column])
	rows[row] = replaceAt(line, 46, column)

	if direction >= 0 {
		b = append(b, findNum(rows, row, column+1, 1)...)
	}

	if direction <= 0 {
		b = append(findNum(rows, row, column-1, -1), b...)
	}

	return b
}

func replaceAt(str string, replace rune, i int) string {
	out := []rune(str)
	out[i] = replace
	return string(out)
}
