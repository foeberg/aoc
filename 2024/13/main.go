package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)
	state := 1
	var as, bs, rs []float64

	for {
		bytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if len(bytes) == 0 {
			state = 1
			continue
		}

		if state == 1 {
			var a, b float64
			_, err := fmt.Sscanf(string(bytes), "Button A: X+%f, Y+%f", &a, &b)
			if err != nil {
				panic(err)
			}
			as = append(as, a)
			bs = append(bs, b)

			state = 2
			continue
		}

		if state == 2 {
			var a, b float64
			fmt.Sscanf(string(bytes), "Button B: X+%f, Y+%f", &a, &b)
			as = append(as, a)
			bs = append(bs, b)

			state = 3
			continue
		}

		if state == 3 {
			var r1, r2 float64
			fmt.Sscanf(string(bytes), "Prize: X=%f, Y=%f", &r1, &r2)
			rs = append(rs, r1, r2)
		}
	}

	fmt.Printf("Part 1: %d\n", part1(as, bs, rs))
	fmt.Printf("Part 2: %d\n", part2(as, bs, rs))
}

func part1(as, bs, rs []float64) int {
	tokens := 0
	for i := 0; i < len(as)-1; i += 2 {
		a := as[i]
		b := as[i+1]
		c := bs[i]
		d := bs[i+1]

		r1 := rs[i]
		r2 := rs[i+1]

		// Solution to the linear equation system
		A := (b*r2 - d*r1) / (b*c - a*d)
		B := (r1 - a*A) / b

		if A > 100 || B > 100 {
			continue
		}

		if math.Trunc(A) != A || math.Trunc(B) != B {
			continue
		}

		tokens += int(A * 3)
		tokens += int(B)
	}

	return tokens
}

func part2(as, bs, rs []float64) int {
	tokens := 0
	for i := 0; i < len(as)-1; i += 2 {
		a := as[i]
		b := as[i+1]
		c := bs[i]
		d := bs[i+1]

		r1 := rs[i] + 10000000000000
		r2 := rs[i+1] + 10000000000000

		// Solution to the linear equation system
		A := (b*r2 - d*r1) / (b*c - a*d)
		B := (r1 - a*A) / b

		if math.Trunc(A) != A || math.Trunc(B) != B {
			continue
		}

		tokens += int(A * 3)
		tokens += int(B)
	}

	return tokens
}
