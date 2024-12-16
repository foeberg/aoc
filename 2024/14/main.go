package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type robot struct {
	x  int
	y  int
	vx int
	vy int
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var robots []robot

	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		var x, y, vx, vy int
		_, err = fmt.Sscanf(string(b), "p=%d,%d v=%d,%d", &x, &y, &vx, &vy)
		if err != nil {
			panic(err)
		}

		robots = append(robots, robot{x, y, vx, vy})
	}

	fmt.Printf("Part 1: %d\n", part1(robots))
	part2(robots)
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func findCycle(sizeX, sizeY int, robot robot) int {
	type coord struct {
		x int
		y int
	}

	visitedPositions := make(map[coord]struct{})
	x, y := robot.x, robot.y
	for i := 0; ; i++ {
		if _, ok := visitedPositions[coord{x, y}]; ok {
			return i
		}
		visitedPositions[coord{x, y}] = struct{}{}

		x = mod((x + robot.vx), sizeX)
		y = mod((y + robot.vy), sizeY)
	}
}

func simulate(sizeX, sizeY, seconds int, robot robot) (int, int) {
	x, y := robot.x, robot.y

	for i := 0; i < seconds; i++ {
		x = mod((x + robot.vx), sizeX)
		y = mod((y + robot.vy), sizeY)
	}

	return x, y
}

func part1(robots []robot) int {
	var q1, q2, q3, q4 int
	for _, robot := range robots {
		cycle := findCycle(101, 103, robot)

		seconds := 100 % cycle
		x, y := simulate(101, 103, seconds, robot)

		if x < 101/2 && y < 103/2 {
			q1++
			continue
		}
		if x > 101/2 && y < 103/2 {
			q2++
			continue
		}
		if x > 101/2 && y > 103/2 {
			q3++
			continue
		}
		if x < 101/2 && y > 103/2 {
			q4++
			continue
		}
	}

	return q1 * q2 * q3 * q4
}

func draw(sizeX, sizeY int, robots []robot) {
	floor := make([][]int, sizeY)
	for i := range floor {
		floor[i] = make([]int, sizeX)
	}

	for _, robot := range robots {
		floor[robot.y][robot.x] = 1
	}

	for r := range floor {
		for c := range floor[r] {
			if floor[r][c] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Print("\n")
	}
}

// Visualized each step, found two interesting patterns on step x & y
//
// With some help from reddit hints and wolframalpha I then calculated
// the n that is congruent to x mod 103 && y mod 101
func part2(robots []robot) {
	for i := 0; ; i++ {
		fmt.Println()
		fmt.Println(i)
		draw(101, 103, robots)
		for i := range robots {
			x, y := simulate(101, 103, 1, robots[i])
			robots[i].x = x
			robots[i].y = y
		}
		time.Sleep(1 * time.Second)
	}
}
