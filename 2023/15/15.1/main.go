package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	hashSum := 0
	reader := bufio.NewReader(f)
	for {
		str, err := reader.ReadString(',')
		if err == io.EOF {
			hashSum += hash(str)
			break
		}
		if err != nil {
			panic(err)
		}

		hashSum += hash(str)
	}

	fmt.Println(hashSum)
}

func hash(input string) int {
	currentValue := 0
	for _, c := range input {
		if c == ',' {
			continue
		}
		currentValue += int(c)
		currentValue *= 17
		currentValue = currentValue % 256
	}

	return currentValue
}
