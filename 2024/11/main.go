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
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var input []byte

	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		input = append(input, b)
	}

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func trimLeadingZeros(stone string) string {
	trimmed := strings.TrimLeft(stone, "0")
	if len(trimmed) == 0 {
		trimmed = "0"
	}
	return trimmed
}

func blink(stone string) []string {
	if stone == "0" {
		return []string{"1"}
	}

	if len(stone)%2 == 0 {
		return []string{stone[0 : len(stone)/2], trimLeadingZeros(stone[len(stone)/2:])}
	}

	i, _ := strconv.Atoi(stone)
	i *= 2024
	s := strconv.Itoa(i)
	return []string{s}
}

func part1(input []byte) int {
	stones := strings.Fields(string(input))

	for i := 0; i < 25; i++ {
		var newStones []string

		for _, stone := range stones {
			newStones = append(newStones, blink(stone)...)
		}
		stones = newStones
	}

	return len(stones)
}

type key struct {
	stone string
	steps int
}

var cache = make(map[key]int)

func countStones(stone string, blinksLeft int) int {
	if res, ok := cache[key{stone, blinksLeft}]; ok {
		return res
	}

	if blinksLeft == 0 {
		return 1
	}

	if stone == "0" {
		res := countStones("1", blinksLeft-1)
		cache[key{stone, blinksLeft}] = res
		return res
	}

	if len(stone)%2 == 0 {
		res := countStones(stone[:len(stone)/2], blinksLeft-1)
		res += countStones(trimLeadingZeros(stone[len(stone)/2:]), blinksLeft-1)
		cache[key{stone, blinksLeft}] = res
		return res
	}

	newStone, _ := strconv.Atoi(stone)
	newStone *= 2024
	res := countStones(strconv.Itoa(newStone), blinksLeft-1)

	cache[key{stone, blinksLeft}] = res

	return res
}

func part2(input []byte) int {
	stones := strings.Fields(string(input))

	count := 0
	for _, stone := range stones {
		count += countStones(stone, 75)
	}

	return count
}
