package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var input []byte

	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		input = append(input, b)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func calculateSize(input []byte) int {
	size := 0
	for _, b := range input {
		size += int(b - 0x30)
	}
	return size
}

func expandMap(input []byte, output []int) {
	currID := 0
	currOutputIndex := 0
	for i, n := range input {
		if i%2 == 0 {
			for j := 0; j < int(n-0x30); j++ {
				output[currOutputIndex+j] = currID
			}
			currID++
		} else {
			for j := 0; j < int(n-0x30); j++ {
				output[currOutputIndex+j] = -1
			}
		}
		currOutputIndex += int(n - 0x30)
	}
}

func compact(input []int) {
	for i := range input {
		if input[i] == -1 {
			for j := len(input) - 1; j >= 0; j-- {
				if j <= i {
					return
				}

				if input[j] != -1 {
					input[i] = input[j]
					input[j] = -1
					break
				}
			}
		}
	}
}

func part1(input []byte) int {
	blocks := make([]int, calculateSize(input))

	expandMap(input, blocks)
	compact(blocks)

	checksum := 0
	for i, id := range blocks {
		if id == -1 {
			break
		}

		checksum += i * id
	}

	return checksum
}

func findGap(input []int, size int) (int, bool) {
	gapSize := 0
	for i, n := range input {
		if gapSize == size {
			return i - (gapSize), true
		}
		if n != -1 {
			gapSize = 0
			continue
		}
		gapSize++
	}

	return 0, false
}

// Returns the start index of a file, and it size
func findLastFile(input []int) (int, int) {
	id := 0
	fileFound := false
	size := 0
	for i := len(input) - 1; i >= 0; i-- {
		if !fileFound {
			if input[i] != -1 {
				fileFound = true
				size++
				id = input[i]
			}

			continue
		}

		if input[i] != id {
			return i + 1, size
		}
		size++
	}

	return 0, size
}

func compactFiles(input []int) {
	remainder := input
	for {
		fileIndex, size := findLastFile(remainder)
		if size == 0 {
			return
		}

		gapIndex, ok := findGap(remainder, size)
		if !ok {
			remainder = remainder[:fileIndex]
			continue
		}

		fileID := remainder[fileIndex]

		slice := remainder[gapIndex : gapIndex+size]
		for i := range slice {
			slice[i] = fileID
			remainder[fileIndex+i] = -1
		}

		remainder = remainder[:fileIndex]
	}
}

func part2(input []byte) int {
	blocks := make([]int, calculateSize(input))

	expandMap(input, blocks)
	compactFiles(blocks)

	checksum := 0
	for i, id := range blocks {
		if id == -1 {
			continue
		}

		checksum += i * id
	}

	return checksum
}
