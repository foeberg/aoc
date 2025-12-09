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

type coord struct {
	X int
	Y int
}

func part1(input []string) int {
	start := findStart(input)

	visited := make(map[coord]struct{})

	followBeam(start, input, visited)

	return len(visited)
}

func part2(input []string) int {
	start := findStart(input)

	cache := make(map[coord]int)

	return followBeam2(start, input, cache)
}

func findStart(input []string) coord {
	for row := range input {
		for col := range input[row] {
			if input[row][col] == 'S' {
				return coord{col, row}
			}
		}
	}
	return coord{}
}

func followBeam(start coord, input []string, visited map[coord]struct{}) {
	currCoord := start
	for {
		if currCoord.Y >= len(input) || currCoord.X >= len(input[0]) || currCoord.Y < 0 || currCoord.X < 0 {
			return
		}
		if input[currCoord.Y][currCoord.X] == '^' {
			if _, ok := visited[currCoord]; ok {
				return
			}

			visited[currCoord] = struct{}{}

			followBeam(coord{currCoord.X - 1, currCoord.Y}, input, visited)
			followBeam(coord{currCoord.X + 1, currCoord.Y}, input, visited)
			return
		}
		currCoord.Y++
	}
}

func followBeam2(start coord, input []string, cache map[coord]int) int {
	currCoord := start
	for {
		if currCoord.Y >= len(input) || currCoord.X >= len(input[0]) || currCoord.Y < 0 || currCoord.X < 0 {
			return 1
		}
		if input[currCoord.Y][currCoord.X] == '^' {
			if downstreamTimelines, ok := cache[currCoord]; ok {
				return downstreamTimelines
			}

			downstreamTimelines := followBeam2(coord{currCoord.X - 1, currCoord.Y}, input, cache) + followBeam2(coord{currCoord.X + 1, currCoord.Y}, input, cache)
			cache[currCoord] = downstreamTimelines
			return downstreamTimelines
		}
		currCoord.Y++
	}
}
