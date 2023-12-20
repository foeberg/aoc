package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type coord struct {
	x int
	y int
}

type instruction struct {
	color     string
	direction string
	len       int
}

func (i instruction) dig(start coord) coord {
	switch i.direction {
	case "U":
		return coord{start.x, start.y - i.len}
	case "R":
		return coord{start.x + i.len, start.y}
	case "D":
		return coord{start.x, start.y + i.len}
	case "L":
		return coord{start.x - i.len, start.y}
	}

	return start
}

func main() {
	plan := parsePlan()

	vertices := []coord{{0, 0}}

	for i := 0; i < len(plan)-1; i++ {
		vertices = append(vertices, plan[i].dig(vertices[i]))
	}

	fmt.Println(calculateArea(vertices))
}

func calculateArea(vertices []coord) int {
	// Shoelace formula for interior area
	intArea := 0
	boundary := 0
	for i := range vertices {
		j := (i + 1) % len(vertices)
		xp := vertices[i].x*vertices[j].y - vertices[i].y*vertices[j].x
		intArea += xp

		boundary += diff(vertices[i], vertices[j])
	}
	intArea = intArea / 2

	// Picks theorem to find i+b
	return (intArea + 1 + boundary/2)
}

func diff(a, b coord) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func parsePlan() []instruction {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var instructions []instruction

	for {
		var c, d string
		var l int
		_, err := fmt.Fscanf(f, "%s %d %s\n", &d, &l, &c)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		c = strings.TrimPrefix(c, "(")
		c = strings.TrimSuffix(c, ")")

		instructions = append(instructions, instruction{c, d, l})
	}

	return instructions
}
