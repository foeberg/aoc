package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	instructions, starterNodes, m := createNodes()

	var loops []int

	for _, sn := range starterNodes {
		var i, count int
		for {
			if i == len(instructions) {
				i = 0
			}

			sn = m[sn][instructions[i]]
			count++

			if strings.HasSuffix(sn, "Z") {
				loops = append(loops, count)
				break
			}

			i++
		}
	}

	fmt.Println(lcm(loops...))
}

// Table method
func lcm(nums ...int) int {
	res := 1

	divisor := 2
	for {
		var changed bool
		var done int

		for i := range nums {
			if nums[i] == 1 {
				done++
				continue
			}

			if nums[i]%divisor == 0 {
				nums[i] = nums[i] / divisor
				changed = true
			}
		}

		if done == len(nums) {
			break
		}

		if changed {
			res = res * divisor
			continue
		}

		if divisor == 2 {
			divisor++
			continue
		}

		divisor += 2
	}

	return res
}

func createNodes() (string, []string, map[string]map[byte]string) {
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
	var starterNodes []string

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

		if strings.HasSuffix(id, "A") {
			starterNodes = append(starterNodes, id)
		}
	}

	return instructions, starterNodes, m
}

func extractNode(str string) (string, string, string) {
	split1 := strings.Split(str, " = ")
	split2 := strings.Split(split1[1], ", ")
	l := strings.TrimPrefix(split2[0], "(")
	r := strings.TrimSuffix(split2[1], ")")
	return split1[0], l, r
}
