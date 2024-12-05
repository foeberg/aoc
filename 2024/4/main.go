package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type coord struct {
	x int
	y int
}

func (c coord) move(d coord) coord {
	return coord{c.x + d.x, c.y + d.y}
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var input [][]byte
	input = append(input, []byte{})

	var col, row int
	var xCoords, aCoords []coord

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
			col = 0
			input = append(input, []byte{})
			continue
		}

		input[row] = append(input[row], b)

		if b == 'X' {
			xCoords = append(xCoords, coord{col, row})
		}
		if b == 'A' {
			aCoords = append(aCoords, coord{col, row})
		}

		col++
	}

	fmt.Printf("Part 1: %v\n", part1(input, xCoords))
	fmt.Printf("Part 2: %v\n", part2(input, aCoords))
}

var keyword = []byte{'M', 'A', 'S'}

var (
	N  = coord{0, -1}
	NE = coord{1, -1}
	E  = coord{1, 0}
	SE = coord{1, 1}
	S  = coord{0, 1}
	SW = coord{-1, 1}
	W  = coord{-1, 0}
	NW = coord{-1, -1}
)

func search(input [][]byte, count int, position, direction coord) int {
	newPos := position.move(direction)

	if newPos.y >= len(input) || newPos.y < 0 {
		return 0
	}
	if newPos.x >= len(input[newPos.y]) || newPos.x < 0 {
		return 0
	}
	if input[newPos.y][newPos.x] != keyword[count] {
		return 0
	}
	if keyword[count] == 'S' {
		return 1
	}

	count++
	return search(input, count, newPos, direction)
}

func part1(input [][]byte, xCoords []coord) int {
	wordCount := 0

	for _, coord := range xCoords {
		wordCount += search(input, 0, coord, N)
		wordCount += search(input, 0, coord, NE)
		wordCount += search(input, 0, coord, E)
		wordCount += search(input, 0, coord, SE)
		wordCount += search(input, 0, coord, S)
		wordCount += search(input, 0, coord, SW)
		wordCount += search(input, 0, coord, W)
		wordCount += search(input, 0, coord, NW)
	}

	return wordCount
}

func part2(input [][]byte, aCoords []coord) int {
	matches := 0

	for _, coord := range aCoords {
		var i int
		i += search(input, 0, coord.move(SW).move(SW), NE)
		i += search(input, 0, coord.move(NW).move(NW), SE)
		i += search(input, 0, coord.move(NE).move(NE), SW)
		i += search(input, 0, coord.move(SE).move(SE), NW)

		if i == 2 {
			matches++
		}
	}

	return matches
}
