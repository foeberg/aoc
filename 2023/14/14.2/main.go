package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	platform := parsePlatform()

	platform = cycle(platform, 1000000000)

	sum := 0
	for r := range platform {
		for c := range platform[r] {
			if platform[r][c] == 'O' {
				sum += len(platform) - r
			}
		}
	}

	fmt.Println(sum)
}

func cycle(platform [][]byte, n int) [][]byte {
	cache := make(map[string]int)
	cache[fmt.Sprint(platform)] = 0

	rem := 0

	for i := 1; i <= n; i++ {
		tiltNorth(platform)
		tiltWest(platform)
		tiltSouth(platform)
		tiltEast(platform)

		sp := fmt.Sprint(platform)
		if start, ok := cache[sp]; ok {
			rem = (n - i) % (start - i)
			break
		}

		cache[sp] = i
	}

	for i := 0; i < rem; i++ {
		tiltNorth(platform)
		tiltWest(platform)
		tiltSouth(platform)
		tiltEast(platform)
	}

	return platform
}

func tiltNorth(platform [][]byte) {
	for column := range platform[0] {
		nextPos := 0
		for row := range platform {
			if platform[row][column] == 'O' {
				if nextPos == row {
					nextPos++
					continue
				}

				platform[nextPos][column] = 'O'
				platform[row][column] = '.'

				nextPos++
				continue
			}

			if platform[row][column] == '#' {
				nextPos = row + 1
				continue
			}
		}
	}
}

func tiltWest(platform [][]byte) {
	for row := range platform {
		nextPos := 0
		for column := range platform[row] {
			if platform[row][column] == 'O' {
				if nextPos == column {
					nextPos++
					continue
				}

				platform[row][nextPos] = 'O'
				platform[row][column] = '.'

				nextPos++
				continue
			}

			if platform[row][column] == '#' {
				nextPos = column + 1
				continue
			}
		}
	}
}

func tiltSouth(platform [][]byte) {
	for column := range platform[0] {
		nextPos := len(platform) - 1
		for row := len(platform) - 1; row >= 0; row-- {
			if platform[row][column] == 'O' {
				if nextPos == row {
					nextPos--
					continue
				}

				platform[nextPos][column] = 'O'
				platform[row][column] = '.'

				nextPos--
				continue
			}

			if platform[row][column] == '#' {
				nextPos = row - 1
				continue
			}
		}
	}
}

func tiltEast(platform [][]byte) {
	for row := range platform {
		nextPos := len(platform[row]) - 1
		for column := len(platform[row]) - 1; column >= 0; column-- {
			if platform[row][column] == 'O' {
				if nextPos == column {
					nextPos--
					continue
				}

				platform[row][nextPos] = 'O'
				platform[row][column] = '.'

				nextPos--
				continue
			}

			if platform[row][column] == '#' {
				nextPos = column - 1
				continue
			}
		}
	}
}

func parsePlatform() [][]byte {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)
	var platform [][]byte

	for {
		b, err := reader.ReadBytes('\n')
		if err == io.EOF {
			platform = append(platform, b)
			break
		}
		if err != nil {
			panic(err)
		}

		platform = append(platform, b[:len(b)-1])
	}

	return platform
}
