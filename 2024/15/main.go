package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

type coord struct {
	x, y int
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var input [][]byte
	var start coord
	var directions []byte

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
			if len(currRow) > 0 {
				input = append(input, currRow)
				currRow = []byte{}
			}
			continue
		}

		if b == '@' {
			start = coord{len(currRow), len(input)}
		}

		if slices.Contains([]byte{'<', '^', '>', 'v'}, b) {
			directions = append(directions, b)
			continue
		}

		currRow = append(currRow, b)
	}

	fmt.Printf("Part 1: %d\n", part1(copyInput(input), start, directions))
	fmt.Printf("Part 2: %d\n", part2(copyInput(input), directions))
}

func copyInput(input [][]byte) [][]byte {
	var new [][]byte

	for y := range input {
		var row []byte
		row = append(row, input[y]...)
		new = append(new, row)
	}

	return new
}

func draw(input [][]byte) {
	for r := range input {
		for c := range input[r] {
			fmt.Printf("%c", input[r][c])
		}
		fmt.Print("\n")
	}
}

func part1(input [][]byte, start coord, directions []byte) int {
	robotPos := start
	for _, dir := range directions {
		robotPos = move(input, robotPos, dir)
	}

	draw(input)

	res := 0
	for r := range input {
		for c := range input[r] {
			if input[r][c] == 'O' {
				res += r*100 + c
			}
		}
	}

	return res
}

func expandInput(input [][]byte) ([][]byte, coord) {
	var newMap [][]byte
	var start coord
	for r := range input {
		newMap = append(newMap, []byte{})
		for c := range input[r] {
			tile := input[r][c]

			if tile == '#' {
				newMap[r] = append(newMap[r], '#', '#')
			}

			if tile == 'O' {
				newMap[r] = append(newMap[r], '[', ']')
			}

			if tile == '.' {
				newMap[r] = append(newMap[r], '.', '.')
			}

			if tile == '@' {
				newMap[r] = append(newMap[r], '@', '.')
				start = coord{len(newMap[r]) - 2, r}
			}
		}
	}
	return newMap, start
}

func canPush(input [][]byte, position coord, direction byte) bool {
	tile := input[position.y][position.x]
	if tile == '#' {
		return false
	}
	if tile == '.' {
		return true
	}

	var nextPos coord
	switch direction {
	case '<':
		nextPos = coord{position.x - 1, position.y}
	case '^':
		nextPos = coord{position.x, position.y - 1}
	case '>':
		nextPos = coord{position.x + 1, position.y}
	case 'v':
		nextPos = coord{position.x, position.y + 1}
	}

	if direction == '<' || direction == '>' {
		return canPush(input, nextPos, direction)
	}

	if tile == '[' {
		return canPush(input, nextPos, direction) &&
			canPush(input, coord{nextPos.x + 1, nextPos.y}, direction)
	}

	if tile == ']' {
		return canPush(input, nextPos, direction) &&
			canPush(input, coord{nextPos.x - 1, nextPos.y}, direction)
	}

	return canPush(input, nextPos, direction)
}

func push(input [][]byte, position coord, direction byte) bool {
	tile := input[position.y][position.x]
	if tile == '#' {
		return false
	}
	if tile == '.' {
		return true
	}

	var nextPos coord
	switch direction {
	case '<':
		nextPos = coord{position.x - 1, position.y}
	case '^':
		nextPos = coord{position.x, position.y - 1}
	case '>':
		nextPos = coord{position.x + 1, position.y}
	case 'v':
		nextPos = coord{position.x, position.y + 1}
	}

	if direction == '<' || direction == '>' {
		success := push(input, nextPos, direction)
		if success {
			input[nextPos.y][nextPos.x] = input[position.y][position.x]
			input[position.y][position.x] = '.'
		}

		return success
	}

	if tile == '[' {
		success := push(input, nextPos, direction) &&
			push(input, coord{position.x + 1, position.y}, direction)
		if success {
			input[nextPos.y][nextPos.x] = input[position.y][position.x]
			input[position.y][position.x] = '.'

			input[nextPos.y][nextPos.x+1] = input[position.y][position.x+1]
			input[position.y][position.x+1] = '.'
		}

		return success
	}

	if tile == ']' {
		success := push(input, nextPos, direction) &&
			push(input, coord{nextPos.x - 1, nextPos.y}, direction)
		if success {
			input[nextPos.y][nextPos.x] = input[position.y][position.x]
			input[position.y][position.x] = '.'

			input[nextPos.y][nextPos.x-1] = input[position.y][position.x-1]
			input[position.y][position.x-1] = '.'
		}

		return success
	}

	success := push(input, nextPos, direction)
	if success {
		input[nextPos.y][nextPos.x] = input[position.y][position.x]
		input[position.y][position.x] = '.'
	}
	return success
}

func move(input [][]byte, position coord, direction byte) coord {
	var nextPos coord
	switch direction {
	case '<':
		nextPos = coord{position.x - 1, position.y}
	case '^':
		nextPos = coord{position.x, position.y - 1}
	case '>':
		nextPos = coord{position.x + 1, position.y}
	case 'v':
		nextPos = coord{position.x, position.y + 1}
	}

	tile := input[nextPos.y][nextPos.x]
	if tile == '#' {
		return position
	}
	if tile == '.' {
		input[position.y][position.x] = '.'
		input[nextPos.y][nextPos.x] = '@'

		return nextPos
	}

	if tile == '[' {
		if canPush(input, nextPos, direction) &&
			canPush(input, coord{nextPos.x + 1, nextPos.y}, direction) {
			if success := push(input, nextPos, direction); success {
				input[position.y][position.x] = '.'
				input[nextPos.y][nextPos.x] = '@'

				return nextPos
			}
		}
	}

	if tile == ']' {
		if canPush(input, nextPos, direction) &&
			canPush(input, coord{nextPos.x - 1, nextPos.y}, direction) {
			if success := push(input, nextPos, direction); success {
				input[position.y][position.x] = '.'
				input[nextPos.y][nextPos.x] = '@'

				return nextPos
			}
		}
	}

	return position
}

func part2(input [][]byte, directions []byte) int {
	input, start := expandInput(input)

	robotPos := start
	// for _, dir := range directions {
	// 	robotPos = move(input, robotPos, dir)
	// 	fmt.Printf("Direction: %c, Position: (%d, %d)\n", dir, robotPos.x, robotPos.y)
	// 	draw(input)
	// 	time.Sleep(500 * time.Millisecond)
	// }
	reader := bufio.NewReader(os.Stdin)
	for {
		char, _, _ := reader.ReadRune()

		var direction byte
		switch char {
		case 'w':
			direction = '^'
		case 'a':
			direction = '<'
		case 's':
			direction = 'v'
		case 'd':
			direction = '>'
		}
		robotPos = move(input, robotPos, direction)
		draw(input)
	}

	// draw(input)

	res := 0
	for r := range input {
		for c := range input[r] {
			if input[r][c] == '[' {
				res += r*100 + c
			}
		}
	}

	return res
}
