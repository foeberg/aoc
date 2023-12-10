package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

const (
	north = iota
	east
	south
	west
)

type vx struct {
	x int
	y int
}

var edges []byte = []byte{'7', 'F', 'J', 'L'}

func main() {
	x, y, ma := parseMap()

	var vertices []vx
	var xMin, xMax, yMin, yMax int
	vertices = append(vertices, vx{x, y})

	direction := findFirstConnection(x, y, ma)

	for {
		x, y = walk(x, y, direction)

		if ma[y][x] == 'S' {
			ma[y][x] = '#'
			break
		}

		if isEdge(ma[y][x]) {
			xMin = min(xMin, x)
			xMax = max(xMax, x)
			yMin = min(yMin, y)
			yMax = max(yMax, y)

			vertices = append(vertices, vx{x, y})
		}

		direction = getNextDirection(direction, ma[y][x])

		ma[y][x] = '#'
	}

	var lines [][]vx

	for i := 1; i < len(vertices); i++ {
		lines = append(lines, []vx{vertices[i-1], vertices[i]})
	}
	lines = append(lines, []vx{vertices[len(vertices)-1], vertices[0]})

	enclosed := 0

	// We only have horizontal and vertical lines. We count line intersections to
	// each point (x,y) from (0,y). This means we can ignore horizontal lines.
	// Even number of intersections means the point is outside.
	for y := range ma {
		for x := range ma[y] {
			if ma[y][x] == '#' {
				continue
			}

			if x >= xMax || x <= xMin || y >= yMax || y <= yMin {
				continue
			}

			intersections := 0
			for _, line := range lines {
				if line[0].x == line[1].x {
					// vertical lines
					if x > line[0].x {
						if y > min(line[0].y, line[1].y) && y <= max(line[0].y, line[1].y) {
							intersections++
						}
					}
					continue
				}
			}

			if intersections%2 == 0 {
				continue
			}
			enclosed++
		}
	}

	fmt.Println(enclosed)
}

func isEdge(b byte) bool {
	return slices.Contains(edges, b)
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
