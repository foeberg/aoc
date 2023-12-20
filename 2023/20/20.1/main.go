package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type power int

const (
	low power = iota
	high
)

type module interface {
	handle(p pulse)
}

type broadcaster struct {
	name         string // debug purposes
	destinations []module
	queue        *pulseQueue
}

func (b *broadcaster) handle(p pulse) {
	for _, dest := range b.destinations {
		b.queue.push(pulse{
			power:       low,
			source:      b,
			destination: dest,
		})
	}
}

type flipflop struct {
	name         string
	on           bool
	destinations []module
	queue        *pulseQueue
}

func (f *flipflop) handle(p pulse) {
	if p.power == high {
		return
	}

	f.on = !f.on

	power := low
	if f.on {
		power = high
	}

	for _, dest := range f.destinations {
		f.queue.push(pulse{
			power:       power,
			source:      f,
			destination: dest,
		})
	}
}

type conjunction struct {
	name         string
	lastInputs   map[module]power
	sources      []module
	destinations []module
	queue        *pulseQueue
}

func (c *conjunction) handle(p pulse) {
	c.lastInputs[p.source] = p.power
	power := high

	if len(c.lastInputs) == len(c.sources) {
		power = low
		for _, pow := range c.lastInputs {
			if pow == low {
				power = high
				break
			}
		}
	}

	for _, dest := range c.destinations {
		c.queue.push(pulse{
			power:       power,
			source:      c,
			destination: dest,
		})
	}
}

type noop struct{}

func (*noop) handle(pulse) {}

type button struct {
	destination module
	queue       *pulseQueue
}

func (b *button) push() {
	b.queue.push(pulse{
		power:       low,
		destination: b.destination,
	})
}

type pulse struct {
	power       power
	source      module
	destination module
}

type pulseQueue []pulse

func (q *pulseQueue) push(p pulse) {
	*q = append(*q, p)
}

func (q *pulseQueue) pop() pulse {
	old := *q
	f := old[0]
	*q = old[1:]
	return f
}

func (q pulseQueue) len() int {
	return len(q)
}

func main() {
	button, pulseQueue := buildModules()
	lows := 0
	highs := 0

	for i := 0; i < 1000; i++ {
		button.push()

		for pulseQueue.len() > 0 {
			pulse := pulseQueue.pop()

			if pulse.power == low {
				lows++
			} else {
				highs++
			}

			pulse.destination.handle(pulse)
		}
	}

	fmt.Println("highs", highs)
	fmt.Println("lows", lows)
	fmt.Println(highs * lows)
}

func buildModules() (*button, *pulseQueue) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	modules := make(map[string]module)
	destinations := make(map[string][]string)
	pulseQueue := make(pulseQueue, 0)

	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		split := strings.Split(string(b), " -> ")

		name := split[0][1:]
		typ := split[0][0:1]

		if typ == "%" {
			modules[name] = &flipflop{
				name:  name,
				queue: &pulseQueue,
			}
		}

		if typ == "&" {
			modules[name] = &conjunction{
				name:       name,
				lastInputs: make(map[module]power),
				queue:      &pulseQueue,
			}
		}

		if typ == "b" {
			name = "broadcaster"
			modules[name] = &broadcaster{
				name:  name,
				queue: &pulseQueue,
			}
		}

		dests := strings.Split(split[1], ", ")
		destinations[name] = append(destinations[name], dests...)
	}

	for name, dests := range destinations {
		module := modules[name]
		for _, dest := range dests {
			destModule := modules[dest]
			if destModule == nil {
				destModule = &noop{}
			}
			addDestination(module, destModule)

			if c, ok := destModule.(*conjunction); ok {
				c.sources = append(c.sources, module)
			}
		}
	}

	button := &button{
		queue:       &pulseQueue,
		destination: modules["broadcaster"],
	}

	return button, &pulseQueue
}

func addDestination(src, dest module) {
	switch t := src.(type) {
	case *broadcaster:
		t.destinations = append(t.destinations, dest)
	case *flipflop:
		t.destinations = append(t.destinations, dest)
	case *conjunction:
		t.destinations = append(t.destinations, dest)
	}
}
