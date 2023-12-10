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

		res += extr[len(extr)-1]
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

	for i := 1; i < len(data); i++ {
		diffs = append(diffs, data[i]-data[i-1])
	}

	rec := extrapolate(diffs)
	return append(data, data[len(data)-1]+rec[len(rec)-1])
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
