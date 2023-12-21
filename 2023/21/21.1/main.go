package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type tile struct {
	x         int
	y         int
	typ       byte
	neighbors []*tile
}

func (t *tile) Equals(b *tile) bool {
	if b == nil {
		return false
	}

	return t.x == b.x && t.y == b.y
}

type step struct {
	currTile  *tile
	prevTile  *tile
	stepsLeft int
}

type stepQueue []step

func (q *stepQueue) push(s step) {
	*q = append(*q, s)
}

func (q *stepQueue) pop() step {
	old := *q
	f := old[0]
	*q = old[1:]
	return f
}

func (q stepQueue) len() int {
	return len(q)
}

type coord struct {
	x int
	y int
}

func main() {
	start := buildGraph()

	stepQueue := make(stepQueue, 0)
	seen := make(map[step]struct{})

	stepQueue.push(step{
		currTile:  start,
		prevTile:  nil,
		stepsLeft: 6,
	})

	endTiles := make(map[coord]struct{})

	for stepQueue.len() > 0 {
		currStep := stepQueue.pop()
		currTile := currStep.currTile
		pos := coord{currTile.x, currTile.y}

		if _, ok := seen[currStep]; ok {
			continue
		}

		seen[currStep] = struct{}{}

		if currStep.stepsLeft%2 == 0 {
			endTiles[pos] = struct{}{}
		}

		if currStep.stepsLeft == 0 {
			continue
		}

		for _, neighbor := range currTile.neighbors {
			if neighbor.Equals(currStep.prevTile) {
				continue
			}

			stepQueue.push(step{
				currTile:  neighbor,
				prevTile:  currTile,
				stepsLeft: currStep.stepsLeft - 1,
			})
		}
	}

	fmt.Println(len(endTiles))
}

func buildGraph() *tile {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var tiles [][]*tile
	var starterTile *tile

	tiles = append(tiles, make([]*tile, 0))

	var x, y int
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if b == '\n' {
			tiles = append(tiles, make([]*tile, 0))
			y++
			x = 0
			continue
		}

		tile := &tile{x: x, y: y, typ: b, neighbors: make([]*tile, 0)}

		if b == '#' {
			x++
			tiles[y] = append(tiles[y], tile)
			continue
		}

		if b == 'S' {
			starterTile = tile
		}

		if x > 0 {
			leftNeighbor := tiles[y][x-1]
			if leftNeighbor.typ == '.' {
				leftNeighbor.neighbors = append(leftNeighbor.neighbors, tile)
				tile.neighbors = append(tile.neighbors, leftNeighbor)
			}
		}

		if y > 0 {
			topNeighbor := tiles[y-1][x]
			if topNeighbor.typ == '.' {
				topNeighbor.neighbors = append(topNeighbor.neighbors, tile)
				tile.neighbors = append(tile.neighbors, topNeighbor)
			}
		}

		tiles[y] = append(tiles[y], tile)

		x++
	}

	return starterTile
}
