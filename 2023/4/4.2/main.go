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

	var (
		sum   int
		lines []string
	)

	m := make(map[int]int)
	i := 0

	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		lines = append(lines, string(b))
		m[i] = 1
		i++
	}

	for i, line := range lines {
		points := 0

		card := strings.Split(strings.Split(line, ":")[1], "|")

		winning := strings.Split(strings.TrimSpace(card[0]), " ")
		got := strings.Split(strings.TrimSpace(card[1]), " ")

		for _, num := range got {
			if num == "" {
				continue
			}

			if slices.Contains(winning, num) {
				points++
			}
		}

		for j := 1; j <= points; j++ {
			_, ok := m[i+j]
			if !ok {
				continue
			}

			m[i+j] += m[i]
		}
	}

	for _, n := range m {
		sum += n
	}

	fmt.Println(sum)
}
