package main

import (
	"bufio"
	"fmt"
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

	timeRow, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	times := toInt(strings.Fields(string(timeRow))[1:])

	distanceRow, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	distances := toInt(strings.Fields(string(distanceRow))[1:])

	res := 0

	for i := 0; i < len(times); i++ {
		wins := 0

		for speed := 1; speed < times[i]; speed++ {
			if speed*(times[i]-speed) > distances[i] {
				wins++
			}
		}
		if res == 0 {
			res = wins
			continue
		}
		res *= wins
	}

	fmt.Println(res)
}

func toInt(str []string) []int {
	var out []int
	for _, s := range str {
		v, _ := strconv.Atoi(s)
		out = append(out, v)
	}
	return out
}
