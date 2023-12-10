package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	north = iota
	east
	south
	west
)

func main() {
	x, y, ma := parseMap()

	direction := findFirstConnection(x, y, ma)

	steps := 0
	for {
		x, y = walk(x, y, direction)
		steps++

		if ma[y][x] == 'S' {
			break
		}

		direction = getNextDirection(direction, ma[y][x])
	}

	fmt.Println(steps / 2)
}

func walk(x, y, direction int) (int, int) {
	switch direction {
	case north:
		return x, y - 1
	case east:
		return x + 1, y
	case south:
		return x, y + 1
	case west:
		return x - 1, y
	}

	return x, y
}

func getNextDirection(lastDirection int, pipe byte) int {
	if lastDirection == north {
		switch pipe {
		case '|':
			return north
		case 'F':
			return east
		case '7':
			return west
		}
	}

	if lastDirection == east {
		switch pipe {
		case '-':
			return east
		case 'J':
			return north
		case '7':
			return south
		}
	}

	if lastDirection == south {
		switch pipe {
		case '|':
			return south
		case 'L':
			return east
		case 'J':
			return west
		}
	}

	if lastDirection == west {
		switch pipe {
		case '-':
			return west
		case 'F':
			return south
		case 'L':
			return north
		}
	}

	return -1
}

func parseMap() (int, int, [][]byte) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var ma [][]byte
	ma = append(ma, []byte{})

	var startX, startY, row int

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
			ma = append(ma, []byte{})
			continue
		}

		ma[row] = append(ma[row], b)

		if b == 'S' {
			startX = len(ma[row]) - 1
			startY = row
		}
	}

	return startX, startY, ma
}

func findFirstConnection(x, y int, ma [][]byte) int {
	// n
	if y > 0 {
		switch ma[y-1][x] {
		case '|', '7', 'F':
			return north
		}
	}

	// e
	if x < len(ma[y])-1 {
		switch ma[y][x+1] {
		case '-', 'J', '7':
			return east
		}
	}

	// s
	if y < len(ma)-1 {
		switch ma[y+1][x] {
		case '|', 'L', 'J':
			return south
		}
	}

	// w
	if x > 0 {
		switch ma[y][x-1] {
		case '-', 'L', 'F':
			return west
		}
	}

	return -1
}
