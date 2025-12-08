package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input")

	var input []string
	for s := bufio.NewScanner(file); s.Scan(); {
		input = append(input, s.Text())
	}

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

var (
	N  = []int{0, -1}
	NE = []int{1, -1}
	E  = []int{1, 0}
	SE = []int{1, 1}
	S  = []int{0, 1}
	SW = []int{-1, 1}
	W  = []int{-1, 0}
	NW = []int{-1, -1}
)

var directions = [][]int{
	N, NE, E, SE, S, SW, W, NW,
}

func part1(input []string) int {
	accessible := 0
	for row := range input {
		for col := range input[row] {
			if input[row][col] == '.' {
				continue
			}

			if adjacentRolls(col, row, input) < 4 {
				accessible++
			}
		}
	}

	return accessible
}

func part2(input []string) int {
	var totalRemoved int

	for {
		var accessibleCoords [][]int
		var removed int

		for row := range input {
			for col := range input[row] {
				if input[row][col] == '.' {
					continue
				}

				if adjacentRolls(col, row, input) < 4 {
					accessibleCoords = append(accessibleCoords, []int{col, row})
				}
			}
		}

		for _, coord := range accessibleCoords {
			row := input[coord[1]]
			newRow := []byte(row)
			newRow[coord[0]] = '.'
			input[coord[1]] = string(newRow)
			removed++
		}

		if removed == 0 {
			break
		}

		totalRemoved += removed
	}

	return totalRemoved
}

func adjacentRolls(col, row int, input []string) int {
	rolls := 0

	for _, direction := range directions {
		newCol := col + direction[0]
		newRow := row + direction[1]
		if inRange(newCol, newRow, input) && input[newRow][newCol] == '@' {
			rolls++
		}
	}

	return rolls
}

func inRange(col, row int, input []string) bool {
	if col < 0 || col > len(input[0])-1 || row < 0 || row > len(input)-1 {
		return false
	}

	return true
}
