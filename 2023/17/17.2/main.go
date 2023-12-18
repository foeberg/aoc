package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
)

type direction int

const (
	north direction = iota
	west
	south
	east
)

type state struct {
	index int

	x           int
	y           int
	dist        int
	consecutive int
	direction   direction
}

func (s state) key() key {
	return key{s.x, s.y, s.consecutive, s.direction}
}

type key struct {
	x           int
	y           int
	consecutive int
	direction   direction
}

type PriorityQueue []*state

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].dist == pq[j].dist {
		return pq[i].consecutive < pq[j].consecutive
	}
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	state := x.(*state)
	state.index = n
	*pq = append(*pq, state)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	state := old[n-1]
	old[n-1] = nil   // avoid memory leak
	state.index = -1 // for safety
	*pq = old[0 : n-1]
	return state
}

func (pq *PriorityQueue) update(state *state, dist int) {
	state.dist = dist
	heap.Fix(pq, state.index)
}

func main() {
	nodes := parseNodes()

	fmt.Println(djikstra(nodes))
}

func djikstra(graph [][]int) int {
	seen := make(map[key]struct{})

	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &state{x: 0, y: 0, dist: 0, consecutive: 0, direction: east})
	heap.Push(&pq, &state{x: 0, y: 0, dist: 0, consecutive: 0, direction: south})

	heap.Init(&pq)

	for pq.Len() > 0 {
		u := heap.Pop(&pq)
		currNode := u.(*state)

		if _, ok := seen[currNode.key()]; ok {
			continue
		}

		seen[currNode.key()] = struct{}{}

		if currNode.y == len(graph)-1 && currNode.x == len(graph[0])-1 && currNode.consecutive >= 4 {
			return currNode.dist
		}

		neighbours := findNeighbors(currNode, graph)
		for _, neigbour := range neighbours {
			if neigbour.direction == currNode.direction {
				if currNode.consecutive < 10 {
					heap.Push(&pq, &state{
						x:           neigbour.x,
						y:           neigbour.y,
						dist:        currNode.dist + graph[neigbour.y][neigbour.x],
						consecutive: currNode.consecutive + 1,
						direction:   neigbour.direction,
					})
				}
				continue
			}

			if currNode.consecutive >= 4 {
				heap.Push(&pq, &state{
					x:           neigbour.x,
					y:           neigbour.y,
					dist:        currNode.dist + graph[neigbour.y][neigbour.x],
					consecutive: 1,
					direction:   neigbour.direction,
				})

			}
		}
	}

	return 0
}

type neighbour struct {
	x         int
	y         int
	direction direction
}

func findNeighbors(s *state, graph [][]int) []neighbour {
	var neighbors []neighbour
	x, y := s.x, s.y

	if x > 0 && s.direction != east {
		neighbors = append(neighbors, neighbour{x - 1, y, west})
	}

	if x < len(graph[y])-1 && s.direction != west {
		neighbors = append(neighbors, neighbour{x + 1, y, east})
	}

	if y > 0 && s.direction != south {
		neighbors = append(neighbors, neighbour{x, y - 1, north})
	}

	if y < len(graph)-1 && s.direction != north {
		neighbors = append(neighbors, neighbour{x, y + 1, south})
	}

	return neighbors
}

func parseNodes() [][]int {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var nodes [][]int
	nodes = append(nodes, make([]int, 0))

	x := 0
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if b == '\n' {
			nodes = append(nodes, make([]int, 0))
			x = 0
			continue
		}

		nodes[len(nodes)-1] = append(nodes[len(nodes)-1], int(b)-48)

		x++
	}

	return nodes
}
