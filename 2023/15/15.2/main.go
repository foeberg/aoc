package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type lens struct {
	label  string
	length int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	boxes := make(map[int][]lens, 256)

	reader := bufio.NewReader(f)
	for {
		str, err := reader.ReadString(',')
		if err != nil && err != io.EOF {
			panic(err)
		}

		le := parseLens(str)
		if le.length == 0 {
			removeLens(le, boxes)
		} else {
			insertLens(le, boxes)
		}

		if err == io.EOF {
			break
		}
	}

	fmt.Println(calculateFocusingPower(boxes))
}

func calculateFocusingPower(boxes map[int][]lens) int {
	power := 0

	for k, v := range boxes {
		for i := range v {
			power += (1 + k) * (i + 1) * v[i].length
		}
	}

	return power
}

func parseLens(str string) lens {
	le := lens{}
	for i := range str {
		if str[i] == '=' {
			le.label = str[:i]
			len, _ := strconv.Atoi(str[i+1 : i+2])
			le.length = len
			break
		}

		if str[i] == '-' {
			le.label = str[:i]
			break
		}
	}

	return le
}

func insertLens(le lens, boxes map[int][]lens) {
	box := hash(le.label)
	lenses := boxes[box]

	for i := range lenses {
		if lenses[i].label == le.label {
			lenses[i] = le
			return
		}
	}

	boxes[box] = append(boxes[box], le)
}

func removeLens(le lens, boxes map[int][]lens) {
	box := hash(le.label)
	lenses := boxes[box]

	for i := range lenses {
		if lenses[i].label == le.label {
			boxes[box] = append(boxes[box][:i], boxes[box][i+1:]...)
			return
		}
	}
}

func hash(input string) int {
	currentValue := 0
	for _, c := range input {
		if c == '=' || c == '-' {
			break
		}
		currentValue += int(c)
		currentValue *= 17
		currentValue = currentValue % 256
	}

	return currentValue
}
