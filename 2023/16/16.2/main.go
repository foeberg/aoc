package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type direction int

const (
	north direction = iota
	west
	south
	east
)

func main() {
	layout := parseLayout()

	maxTiles := 0

	// Leftmost & rightmost column
	for i := range layout {
		leftCache := make(map[key][]direction)
		rightCache := make(map[key][]direction)

		followRay(0, i, east, layout, leftCache)
		followRay(len(layout[i])-1, i, west, layout, rightCache)

		maxTiles = max(maxTiles, len(leftCache), len(rightCache))
	}

	// top & bottom row
	for i := range layout[0] {
		topCache := make(map[key][]direction)
		botCache := make(map[key][]direction)

		followRay(i, 0, south, layout, topCache)
		followRay(i, len(layout)-1, north, layout, botCache)

		maxTiles = max(maxTiles, len(topCache), len(botCache))
	}

	fmt.Println(maxTiles)
}

type key struct {
	x int
	y int
}

// Follows a ray of light starting at (x, y) heading in direction dir until a mirror is found.
func followRay(x, y int, dir direction, layout [][]byte, cache map[key][]direction) {
	for {
		// Ray has gone past the edge of the layout
		if y >= len(layout) || x >= len(layout[0]) || x < 0 || y < 0 {
			return
		}

		tile := layout[y][x]

		// If position has been visited before, with the same direction, return.
		if dirs, ok := cache[key{x, y}]; ok {
			for _, cdir := range dirs {
				if cdir == dir {
					return
				}
			}
		}

		cache[key{x, y}] = append(cache[key{x, y}], dir)

		if tile == '.' ||
			(tile == '-' && dir%2 != 0) ||
			(tile == '|' && dir%2 == 0) {

			x, y = nextTile(x, y, dir)
			continue
		}

		if tile == '\\' {
			if dir == north {
				nx, ny := nextTile(x, y, west)
				followRay(nx, ny, west, layout, cache)
				return
			}

			if dir == west {
				nx, ny := nextTile(x, y, north)
				followRay(nx, ny, north, layout, cache)
				return
			}

			if dir == south {
				nx, ny := nextTile(x, y, east)
				followRay(nx, ny, east, layout, cache)
				return
			}

			if dir == east {
				nx, ny := nextTile(x, y, south)
				followRay(nx, ny, south, layout, cache)
				return
			}
		}

		if tile == '/' {
			if dir == north {
				nx, ny := nextTile(x, y, east)
				followRay(nx, ny, east, layout, cache)
				return
			}

			if dir == west {
				nx, ny := nextTile(x, y, south)
				followRay(nx, ny, south, layout, cache)
				return
			}

			if dir == south {
				nx, ny := nextTile(x, y, west)
				followRay(nx, ny, west, layout, cache)
				return
			}

			if dir == east {
				nx, ny := nextTile(x, y, north)
				followRay(nx, ny, north, layout, cache)
				return
			}
		}

		if tile == '|' {
			followRay(x, y-1, north, layout, cache)
			followRay(x, y+1, south, layout, cache)
			return
		}

		if tile == '-' {
			followRay(x+1, y, east, layout, cache)
			followRay(x-1, y, west, layout, cache)
			return
		}
	}
}

func nextTile(x, y int, dir direction) (int, int) {
	switch dir {
	case north:
		return x, y - 1
	case west:
		return x - 1, y
	case south:
		return x, y + 1
	case east:
		return x + 1, y
	}

	return x, y
}

func parseLayout() [][]byte {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)
	var layout [][]byte

	for {
		b, err := reader.ReadBytes('\n')
		if err == io.EOF {
			layout = append(layout, b)
			break
		}
		if err != nil {
			panic(err)
		}

		layout = append(layout, b[:len(b)-1])
	}

	return layout
}
