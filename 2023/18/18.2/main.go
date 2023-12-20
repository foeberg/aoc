package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int64
	y int64
}

type instruction struct {
	direction string
	len       int64
}

func (i instruction) dig(start coord) coord {
	switch i.direction {
	case "0":
		return coord{start.x + i.len, start.y}
	case "1":
		return coord{start.x, start.y + i.len}
	case "2":
		return coord{start.x - i.len, start.y}
	case "3":
		return coord{start.x, start.y - i.len}
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

func calculateArea(vertices []coord) int64 {
	// Shoelace formula for interior area
	var intArea, boundary int64

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

func diff(a, b coord) int64 {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(n int64) int64 {
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

	reader := bufio.NewReader(f)
	var instructions []instruction

	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		c := strings.Split(string(b), " ")[2]
		c = strings.TrimPrefix(c, "(#")
		c = strings.TrimSuffix(c, ")")

		l, _ := strconv.ParseInt(c[:len(c)-1], 16, 32)

		instructions = append(instructions, instruction{c[len(c)-1:], l})
	}

	return instructions
}
