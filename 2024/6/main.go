package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

type coord struct {
	x int
	y int
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
	var start coord

	var directions = []byte{'<', '^', '>', 'v'}

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

		if slices.Contains(directions, b) {
			start = coord{col, row}
		}

		col++
	}

	fmt.Printf("Part 1: %d\n", part1(copyMap(input), start))
	fmt.Printf("Part 2: %d\n", part2(copyMap(input), start))
}

func copyMap(input [][]byte) [][]byte {
	var new [][]byte

	for y := range input {
		var row []byte
		row = append(row, input[y]...)
		new = append(new, row)
	}

	return new
}

func changeDirection(b byte) byte {
	switch b {
	case '<':
		return '^'
	case '^':
		return '>'
	case '>':
		return 'v'
	case 'v':
		return '<'
	default:
		panic(fmt.Sprintf("changeDirection: unexpected direction %c", b))
	}
}

func outOfBounds(input [][]byte, pos coord) bool {
	if pos.y < 0 || pos.x < 0 || pos.y >= len(input) || pos.x >= len(input[pos.y]) {
		return true
	}
	return false
}

func obstructed(input [][]byte, pos coord, direction byte) bool {
	x, y := pos.x, pos.y

	switch direction {
	case '<':
		if !outOfBounds(input, coord{x - 1, y}) && input[y][x-1] == '#' {
			return true
		}
	case '^':
		if !outOfBounds(input, coord{x, y - 1}) && input[y-1][x] == '#' {
			return true
		}
	case '>':
		if !outOfBounds(input, coord{x + 1, y}) && input[y][x+1] == '#' {
			return true
		}
	case 'v':
		if !outOfBounds(input, coord{x, y + 1}) && input[y+1][x] == '#' {
			return true
		}
	default:
		panic(fmt.Sprintf("obstructed: unexpected direction %c", direction))
	}

	return false
}

func calculateNewPos(pos coord, direction byte) coord {
	x, y := pos.x, pos.y

	switch direction {
	case '<':
		return coord{x - 1, y}
	case '^':
		return coord{x, y - 1}
	case '>':
		return coord{x + 1, y}
	case 'v':
		return coord{x, y + 1}
	default:
		panic(fmt.Sprintf("calculateNewPos: unexpected direction %c", direction))
	}
}

func part1(input [][]byte, start coord) int {
	currPos := start

	visited := make(map[coord]struct{})

	direction := input[currPos.y][currPos.x]
	for {
		if obstructed(input, currPos, direction) {
			direction = changeDirection(direction)
			continue
		}

		visited[currPos] = struct{}{}

		newPos := calculateNewPos(currPos, direction)
		if outOfBounds(input, newPos) {
			break
		}

		input[currPos.y][currPos.x] = 'X'
		currPos = newPos
	}

	return len(visited)
}

type coordAndDirection struct {
	coord     coord
	direction byte
}

func loopsIfObstructed(input [][]byte, pos coord, direction byte) bool {
	newPos := calculateNewPos(pos, direction)
	if outOfBounds(input, newPos) {
		return false
	}
	if input[newPos.y][newPos.x] == 'X' {
		return false
	}

	input[newPos.y][newPos.x] = '#'
	defer func() {
		input[newPos.y][newPos.x] = '.'
	}()

	visited := make(map[coordAndDirection]struct{})
	cad := coordAndDirection{pos, direction}
	visited[cad] = struct{}{}

	direction = changeDirection(direction)
	for {
		cad := coordAndDirection{pos, direction}
		if _, ok := visited[cad]; ok {
			return true
		}

		visited[cad] = struct{}{}

		if obstructed(input, pos, direction) {
			direction = changeDirection(direction)
			continue
		}

		pos = calculateNewPos(pos, direction)
		if outOfBounds(input, pos) {
			return false
		}
	}
}

func part2(input [][]byte, start coord) int {
	currPos := start

	obs := 0

	direction := input[currPos.y][currPos.x]
	for {

		if obstructed(input, currPos, direction) {
			direction = changeDirection(direction)
			continue
		}

		if loopsIfObstructed(input, currPos, direction) {
			obs++
		}

		newPos := calculateNewPos(currPos, direction)
		if outOfBounds(input, newPos) {
			break
		}

		input[currPos.y][currPos.x] = 'X'
		currPos = newPos
	}

	return obs
}
