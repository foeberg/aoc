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

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var input [][]int

	var col, row int
	var zeros, nines []coord

	var currRow []int
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
				currRow = []int{}
			}
			continue
		}

		currRow = append(currRow, int(b-0x30))

		if b == '0' {
			zeros = append(zeros, coord{col, row})
		}
		if b == '9' {
			nines = append(nines, coord{col, row})
		}

		col++
	}

	fmt.Printf("Part 1: %d\n", part1(input, zeros))
	fmt.Printf("Part 2: %d\n", part2(input, nines))
}

func part1(input [][]int, zeros []coord) int {
	sum := 0
	for _, zero := range zeros {
		sum += ascendTrail(input, zero)
	}
	return sum
}

func ascendTrail(input [][]int, start coord) int {
	nextPositions := []coord{start}
	nines := map[coord]struct{}{}

	for {
		if len(nextPositions) == 0 {
			return len(nines)
		}

		currPos := nextPositions[0]
		currNum := input[currPos.y][currPos.x]
		nextPositions = nextPositions[1:]

		x, y := currPos.x, currPos.y

		if currNum == 9 {
			nines[currPos] = struct{}{}
			continue
		}

		// N
		if y > 0 {
			if input[y-1][x] == currNum+1 {
				nextPositions = append(nextPositions, coord{x, y - 1})
			}
		}
		// E
		if x < len(input[y])-1 {
			if input[y][x+1] == currNum+1 {
				nextPositions = append(nextPositions, coord{x + 1, y})
			}
		}
		// S
		if y < len(input)-1 {
			if input[y+1][x] == currNum+1 {
				nextPositions = append(nextPositions, coord{x, y + 1})
			}
		}
		// W
		if x > 0 {
			if input[y][x-1] == currNum+1 {
				nextPositions = append(nextPositions, coord{x - 1, y})
			}
		}
	}
}

func part2(input [][]int, nines []coord) int {
	sum := 0
	for _, nine := range nines {
		sum += descendTrail(input, nine)
	}
	return sum
}

func descendTrail(input [][]int, start coord) int {
	nextPositions := []coord{start}
	score := 0

	for {
		if len(nextPositions) == 0 {
			return score
		}

		currPos := nextPositions[0]
		currNum := input[currPos.y][currPos.x]
		nextPositions = nextPositions[1:]

		x, y := currPos.x, currPos.y

		if currNum == 0 {
			score++
			continue
		}

		// N
		if y > 0 {
			if input[y-1][x] == currNum-1 {
				nextPositions = append(nextPositions, coord{x, y - 1})
			}
		}
		// E
		if x < len(input[y])-1 {
			if input[y][x+1] == currNum-1 {
				nextPositions = append(nextPositions, coord{x + 1, y})
			}
		}
		// S
		if y < len(input)-1 {
			if input[y+1][x] == currNum-1 {
				nextPositions = append(nextPositions, coord{x, y + 1})
			}
		}
		// W
		if x > 0 {
			if input[y][x-1] == currNum-1 {
				nextPositions = append(nextPositions, coord{x - 1, y})
			}
		}
	}
}
