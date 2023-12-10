package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	instructions, m := createNodes()

	currNode := "AAA"

	var i, count int
	for {
		if i == len(instructions) {
			i = 0
		}

		instruction := instructions[i]

		currNode = m[currNode][instruction]
		count++

		if currNode == "ZZZ" {
			break
		}
		i++
	}

	fmt.Println(count)
}

func createNodes() (string, map[string]map[byte]string) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	ib, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	instructions := string(ib)

	m := make(map[string]map[byte]string)

	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if len(b) == 0 {
			continue
		}

		id, l, r := extractNode(string(b))

		n := map[byte]string{
			'L': l,
			'R': r,
		}

		m[id] = n
	}

	return instructions, m
}

func extractNode(str string) (string, string, string) {
	split1 := strings.Split(str, " = ")
	split2 := strings.Split(split1[1], ", ")
	l := strings.TrimPrefix(split2[0], "(")
	r := strings.TrimSuffix(split2[1], ")")
	return split1[0], l, r
}
