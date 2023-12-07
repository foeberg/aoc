package main

import (
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

	var (
		sum    int
		first  byte
		last   byte
		substr []byte
	)

	for {
		buf := make([]byte, 1)
		_, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		r := buf[0]

		if r == 0xA {
			val, _ := strconv.Atoi(string(fmt.Sprintf("%c%c", first, last)))

			sum += val

			first = 0
			last = 0
		}

		if r < 0x40 && r > 0x2F {
			substr = []byte{}
			if first == 0 {
				first = r
			}
			last = r
			continue
		}

		substr = append(substr, r)

		if num := findNum(substr); num != 0 {
			substr = []byte{r}

			if first == 0 {
				first = num
			}
			last = num
		}
	}

	fmt.Println(sum)
}

func findNum(strb []byte) byte {
	nums := map[string]string{
		"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
	}

	str := string(strb)
	for strNum, num := range nums {
		if strings.Contains(str, strNum) {
			return num[0]
		}
	}

	return 0
}
