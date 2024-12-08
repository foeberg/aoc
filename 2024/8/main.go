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

	var input [][]byte
	input = append(input, []byte{})

	var col, row int
	antennas := make(map[byte][]coord)

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

		if b != '.' {
			antennas[b] = append(antennas[b], coord{col, row})
		}

		input[row] = append(input[row], b)

		col++
	}

	fmt.Printf("Part 1: %d\n", part1(copyMap(input), antennas))
	fmt.Printf("Part 2: %d\n", part2(copyMap(input), antennas))
}

func outOfBounds(input [][]byte, pos coord) bool {
	if pos.y < 0 || pos.x < 0 || pos.y >= len(input) || pos.x >= len(input[pos.y]) {
		return true
	}
	return false
}

func calculateDistance(a, b coord) coord {
	return coord{a.x - b.x, a.y - b.y}
}

func findAntinodes(antennaCoords []coord) []coord {
	var antinodes []coord
	for i := 0; i < len(antennaCoords)-1; i++ {
		for j := i + 1; j < len(antennaCoords); j++ {
			dist := calculateDistance(antennaCoords[i], antennaCoords[j])

			n1 := coord{
				x: antennaCoords[i].x + dist.x,
				y: antennaCoords[i].y + dist.y,
			}
			n2 := coord{
				x: antennaCoords[j].x - dist.x,
				y: antennaCoords[j].y - dist.y,
			}

			antinodes = append(antinodes, n1, n2)
		}
	}

	return antinodes
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

func part1(input [][]byte, antennas map[byte][]coord) int {
	nodeCount := 0
	for _, antennaCoords := range antennas {
		antinodes := findAntinodes(antennaCoords)
		for _, node := range antinodes {
			if outOfBounds(input, node) {
				continue
			}
			if input[node.y][node.x] == '#' {
				continue
			}

			input[node.y][node.x] = '#'
			nodeCount++
		}
	}

	return nodeCount
}

func findAntinodes2(input [][]byte, antennaCoords []coord) []coord {
	var antinodes []coord
	for i := 0; i < len(antennaCoords)-1; i++ {
		for j := i + 1; j < len(antennaCoords); j++ {
			dist := calculateDistance(antennaCoords[i], antennaCoords[j])

			for mult := 0; ; mult++ {
				n1 := coord{
					x: antennaCoords[i].x + mult*dist.x,
					y: antennaCoords[i].y + mult*dist.y,
				}
				n2 := coord{
					x: antennaCoords[j].x - mult*dist.x,
					y: antennaCoords[j].y - mult*dist.y,
				}

				newNode := false
				if !outOfBounds(input, n1) {
					newNode = true
					antinodes = append(antinodes, n1)
				}

				if !outOfBounds(input, n2) {
					newNode = true
					antinodes = append(antinodes, n2)
				}

				if !newNode {
					break
				}
			}
		}
	}

	return antinodes
}

func part2(input [][]byte, antennas map[byte][]coord) int {
	nodeCount := 0

	for _, antennaCoords := range antennas {
		antinodes := findAntinodes2(input, antennaCoords)
		for _, node := range antinodes {
			if input[node.y][node.x] == '#' {
				continue
			}

			input[node.y][node.x] = '#'
			nodeCount++
		}
	}

	return nodeCount
}
