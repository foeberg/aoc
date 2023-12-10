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

	res := 0
	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		history := getHistory(b)
		extr := extrapolate(history)

		res += extr[0]
	}

	fmt.Println(res)
}

func extrapolate(data []int) []int {
	var zeros int
	for _, n := range data {
		if n == 0 {
			zeros++
		}
	}
	if zeros == len(data) {
		return append(data, 0)
	}

	var diffs []int

	for i := len(data) - 1; i > 0; i-- {
		diffs = append([]int{data[i] - data[i-1]}, diffs...)
	}

	rec := extrapolate(diffs)
	return append([]int{data[0] - rec[0]}, data...)
}

func getHistory(b []byte) []int {
	var res []int
	split := strings.Split(string(b), " ")

	for _, s := range split {
		n, _ := strconv.Atoi(s)
		res = append(res, n)
	}

	return res
}
