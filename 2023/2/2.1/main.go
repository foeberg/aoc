package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var (
		maxRed   = 12
		maxGreen = 13
		maxBlue  = 14

		sum = 0
	)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		n, r, g, b := parseGame(line)

		if r <= maxRed && g <= maxGreen && b <= maxBlue {
			sum += n
		}
	}

	fmt.Println(sum)
}

func parseGame(line []byte) (int, int, int, int) {
	var r, g, b int

	strLine := string(line)

	split := strings.Split(strLine, ":")

	n, err := strconv.Atoi(strings.Split(split[0], " ")[1])
	if err != nil {
		panic(err)
	}

	rounds := strings.Split(split[1], ";")

	for _, round := range rounds {
		cubes := strings.Split(round, ",")

		for _, cube := range cubes {
			cube := strings.TrimSpace(cube)

			numColor := strings.Split(cube, " ")

			amount, err := strconv.Atoi(numColor[0])
			if err != nil {
				panic(err)
			}

			if numColor[1] == "red" {
				r = max(r, amount)
				continue
			}

			if numColor[1] == "green" {
				g = max(g, amount)
				continue
			}

			if numColor[1] == "blue" {
				b = max(b, amount)
				continue
			}
		}
	}

	return n, r, g, b
}
