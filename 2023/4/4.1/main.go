package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	sum := 0

	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		line := string(b)
		points := 0

		card := strings.Split(strings.Split(line, ":")[1], "|")

		winning := strings.Split(strings.TrimSpace(card[0]), " ")
		got := strings.Split(strings.TrimSpace(card[1]), " ")

		for _, num := range got {
			if num == "" {
				continue
			}

			if slices.Contains(winning, num) {
				if points == 0 {
					points = 1
					continue
				}
				points = points << 1
			}
		}

		sum += points
	}

	fmt.Println(sum)
}
