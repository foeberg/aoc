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

	var input [][]byte

	var col, row int

	var currRow []byte
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			if len(currRow) > 0 {
				input = append(input, currRow)
			}
			break
		}
		if err != nil {
			panic(err)
		}
		if b == '\n' {
			row++
			col = 0

			if len(currRow) > 0 {
				input = append(input, currRow)
				currRow = []byte{}
			}
			continue
		}

		currRow = append(currRow, b)
		col++
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input [][]byte) int {
	return 0
}
