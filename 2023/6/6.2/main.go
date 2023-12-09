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
	time := toInt(merge(strings.Fields(string(timeRow))[1:]))

	distanceRow, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	distance := toInt(merge(strings.Fields(string(distanceRow))[1:]))

	wins := 0

	for speed := 1; speed < time; speed++ {
		if speed*(time-speed) > distance {
			wins++
		}
	}

	fmt.Println(wins)
}

func merge(slice []string) string {
	var out []byte

	for _, str := range slice {
		out = append(out, []byte(str)...)
	}

	return string(out)
}

func toInt(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}
