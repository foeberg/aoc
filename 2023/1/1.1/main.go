package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var (
		sum   int
		first rune
		last  rune
	)

	for {
		buf := make([]byte, 1)
		_, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		r := rune(buf[0])

		if r == 0xA {
			val, _ := strconv.Atoi(string(fmt.Sprintf("%c%c", first, last)))
			sum += val

			first = 0
			last = 0
		}

		if r < 0x40 && r > 0x2F {
			if first == 0 {
				first = r
				last = r
				continue
			}
			last = r
		}
	}

	fmt.Println(sum)
}
