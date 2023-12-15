package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	rows := parseRows()

	sum := 0
	for i := range rows[0] {
		s := countColumn(getColumn(i, rows))
		sum += s
	}

	fmt.Println(sum)
}

func countColumn(column string) int {
	weight := 0
	multiplier := len(column)

	for i := range column {
		if column[i] == 'O' {
			weight += multiplier
			multiplier--
			continue
		}

		if column[i] == '#' {
			multiplier = len(column) - i - 1
			continue
		}
	}

	return weight
}

func parseRows() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)
	var rows []string

	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		rows = append(rows, string(b))
	}
	return rows
}

func getColumn(col int, rows []string) string {
	var bs []byte
	for i := 0; i < len(rows); i++ {
		bs = append(bs, byte(rows[i][col]))
	}
	return string(bs)
}
