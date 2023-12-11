package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type coord struct {
	x int
	y int
}

type pair struct {
	a coord
	b coord
}

func main() {
	image := parseImage()

	expandedImage := expandImage(image)

	galaxies := findGalaxies(expandedImage)

	pairs := createPairs(galaxies)

	tot := 0
	for _, pair := range pairs {
		tot += distance(pair, expandedImage)
	}

	fmt.Println(tot)
}

func distance(p pair, image []string) int {
	xDiff := max(p.a.x, p.b.x) - min(p.a.x, p.b.x)
	yDiff := max(p.a.y, p.b.y) - min(p.a.y, p.b.y)

	dist := xDiff + yDiff

	for x := min(p.a.x, p.b.x); x < max(p.a.x, p.b.x); x++ {
		if image[0][x] == 'X' {
			dist += 999998
		}
	}

	for y := min(p.a.y, p.b.y); y < max(p.a.y, p.b.y); y++ {
		if image[y][0] == 'X' {
			dist += 999998
		}
	}

	return dist
}

func createPairs(items []coord) []pair {
	var pairs []pair

	for i := 0; i < len(items); i++ {
		for j := i + 1; j < len(items); j++ {
			pairs = append(pairs, pair{items[i], items[j]})
		}
	}

	return pairs
}

func findGalaxies(image []string) []coord {
	var galaxies []coord
	for y, row := range image {
		for x, c := range row {
			if c == '#' {
				galaxies = append(galaxies, coord{x, y})
			}
		}
	}

	return galaxies
}

func expandImage(image []string) []string {
	var expandedImage []string

	for _, row := range image {
		expandedImage = append(expandedImage, row)

		var found bool
		for _, c := range row {
			if c == '#' {
				found = true
				break
			}
		}

		if !found {
			expandedImage = append(expandedImage, strings.Repeat("X", len(row)))
		}
	}

	var expandIndices []int

	for col := 0; col < len(expandedImage[0]); col++ {
		var found bool
		for row := 0; row < len(expandedImage); row++ {
			if expandedImage[row][col] == '#' {
				found = true
				break
			}
		}

		if !found {
			expandIndices = append(expandIndices, col)
		}
	}

	inserts := 0
	for _, i := range expandIndices {
		for j := range expandedImage {
			bs := []byte(expandedImage[j])
			expandedImage[j] = string(slices.Insert(bs, i+inserts, 'X'))
		}
		inserts++
	}

	return expandedImage
}

func parseImage() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var lines []string

	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		lines = append(lines, string(b))
	}

	return lines
}
