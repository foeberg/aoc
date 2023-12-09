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
	seeds, mappings := populate()

	loc := -1
	for i := 1; i < len(seeds); i += 2 {
		seedStart, _ := strconv.Atoi(seeds[i-1])
		seedLen, _ := strconv.Atoi(seeds[i])

		for seed := seedStart; seed < seedStart+seedLen; seed++ {
			source := seed
			for j := range mappings {
				source = sourceToDest(source, mappings[j])
			}

			if loc == -1 {
				loc = source
				continue
			}

			loc = min(loc, source)
		}
	}

	fmt.Println(loc)
}

type mapping struct {
	destination int
	source      int
	len         int
}

func populate() ([]string, [][]mapping) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var seeds []string

	mapIndex := -1
	input := make([][]mapping, 7)

	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if len(b) == 0 {
			mapIndex++
			continue
		}

		if strings.HasPrefix(string(b), "seeds:") {
			seeds = strings.Split(string(b), " ")[1:]
			continue
		}

		if !isNum(b[0]) {
			continue
		}

		input[mapIndex] = append(input[mapIndex], createMapping(b))
	}

	return seeds, input
}

func createMapping(b []byte) mapping {
	values := strings.Split(string(b), " ")
	d, _ := strconv.Atoi(values[0])
	s, _ := strconv.Atoi(values[1])
	l, _ := strconv.Atoi(values[2])

	return mapping{destination: d, source: s, len: l}
}

func isNum(b byte) bool {
	if b < 59 && b > 47 {
		return true
	}

	return false
}

func sourceToDest(source int, mappings []mapping) int {
	for i := range mappings {
		mapping := mappings[i]

		if source >= mapping.source && source < mapping.source+mapping.len {
			diff := source - mapping.source
			return mapping.destination + diff
		}
	}
	return source
}
