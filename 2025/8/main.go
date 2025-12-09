package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input")

	var input []string
	for s := bufio.NewScanner(file); s.Scan(); {
		input = append(input, s.Text())
	}

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

type box struct {
	ID          int
	CircuitID   int
	X, Y, Z     float64
	Connections []*box
}

func (b *box) connectedTo(a *box) bool {
	for _, connectedBox := range b.Connections {
		if connectedBox.ID == a.ID {
			return true
		}
	}
	return false
}

func part1(input []string) int {
	boxes := parseBoxes(input)

	for range 1000 {
		closest := closestBoxes(boxes)

		closest[0].Connections = append(closest[0].Connections, closest[1])
		closest[1].Connections = append(closest[1].Connections, closest[0])
	}

	sizes := calculateCircuitSizes(boxes)

	sum := 1
	for i := 0; i < 3; i++ {
		sum *= sizes[i]
	}

	return sum
}

func part2(input []string) int {
	boxes := parseBoxes(input)

	for {
		closest := closestBoxes(boxes)

		connectBoxes(closest[0], closest[1])

		var circuitIDs []int
		for _, box := range boxes {
			circuitIDs = append(circuitIDs, box.CircuitID)
		}

		slices.Sort(circuitIDs)
		circuitIDs = slices.Compact(circuitIDs)

		if len(circuitIDs) == 1 {
			return int(closest[0].X) * int(closest[1].X)
		}
	}
}

func connectBoxes(a, b *box) {
	a.Connections = append(a.Connections, b)
	b.Connections = append(b.Connections, a)

	if a.CircuitID == b.CircuitID {
		return
	}

	updateCircuitID(b, a.CircuitID)
}

func updateCircuitID(b *box, circuitID int) {
	if b.CircuitID == circuitID {
		return
	}

	b.CircuitID = circuitID
	for _, connectedBox := range b.Connections {
		updateCircuitID(connectedBox, circuitID)
	}
}

func closestBoxes(boxes []*box) []*box {
	closestDistance := math.Inf(1)
	var pair []*box
	for a := 0; a < len(boxes); a++ {
		for b := a + 1; b < len(boxes); b++ {
			if boxes[a].connectedTo(boxes[b]) {
				continue
			}

			if dist := distance(boxes[a], boxes[b]); dist <= closestDistance {
				closestDistance = dist
				pair = []*box{boxes[a], boxes[b]}
			}
		}
	}

	return pair
}

func distance(a, b *box) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2) + math.Pow(a.Z-b.Z, 2))
}

func parseBoxes(input []string) []*box {
	var boxes []*box
	for i := range input {
		split := strings.Split(input[i], ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])

		boxes = append(boxes, &box{
			ID:          len(boxes),
			CircuitID:   len(boxes),
			X:           float64(x),
			Y:           float64(y),
			Z:           float64(z),
			Connections: make([]*box, 0),
		})
	}

	return boxes
}

func calculateCircuitSizes(boxes []*box) []int {
	visitedBoxes := make(map[int]struct{})

	var sizes []int
	for _, box := range boxes {
		if _, ok := visitedBoxes[box.ID]; ok {
			continue
		}

		visited := make(map[int]struct{})

		sizes = append(sizes, traverse(box, visited))

		for vist := range visited {
			visitedBoxes[vist] = struct{}{}
		}
	}

	slices.SortFunc(sizes, func(a, b int) int {
		if a < b {
			return 1
		}
		return -1
	})

	return sizes
}

func traverse(currBox *box, visited map[int]struct{}) int {
	if _, ok := visited[currBox.ID]; ok {
		return 0
	}
	visited[currBox.ID] = struct{}{}

	count := 1
	for _, connectedBox := range currBox.Connections {
		count += traverse(connectedBox, visited)
	}

	return count
}
