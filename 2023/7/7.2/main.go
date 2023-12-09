package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"slices"
)

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

var cardValues = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

type hand struct {
	hand string
	bid  int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var hands []hand

	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		handBids := strings.Split(string(b), " ")

		hands = append(hands, hand{hand: handBids[0], bid: toInt(handBids[1])})
	}

	slices.SortFunc(hands, func(a hand, b hand) int {
		aType := getType(a.hand)
		bType := getType(b.hand)

		if aType < bType {
			return -1
		}

		if aType > bType {
			return 1
		}

		return cmpEqualHands(a.hand, b.hand)
	})

	sum := 0
	for i, h := range hands {
		sum += (i + 1) * h.bid
	}

	fmt.Println(sum)
}

func getType(hand string) int {
	var pairs, triples, quads, fulls int
	m := createCountMap()

	for _, card := range hand {
		m[card]++
	}

	for _, counts := range m {
		switch counts {
		case 2:
			pairs++
		case 3:
			triples++
		case 4:
			quads++
		case 5:
			fulls++
		}
	}

	jokers := m['J']

	if fulls == 1 {
		return FiveOfKind
	}

	if quads == 1 {
		if jokers == 1 || jokers == 4 {
			return FiveOfKind
		}
		return FourOfKind
	}

	if triples == 1 && pairs == 1 {
		if jokers >= 2 {
			return FiveOfKind
		}
		return FullHouse
	}

	if triples == 1 {
		if jokers == 1 || jokers == 3 {
			return FourOfKind
		}
		return ThreeOfKind
	}

	if pairs == 2 {
		if jokers == 2 {
			return FourOfKind
		}
		if jokers == 1 {
			return FullHouse
		}
		return TwoPair
	}

	if pairs == 1 {
		if jokers >= 1 {
			return ThreeOfKind
		}
		return OnePair
	}

	if jokers == 1 {
		return OnePair
	}

	return HighCard
}

func cmpEqualHands(a, b string) int {
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}

		if cardValues[rune(a[i])] > cardValues[rune(b[i])] {
			return 1
		}
		if cardValues[rune(a[i])] < cardValues[rune(b[i])] {
			return -1
		}
	}
	return 0
}

func toInt(b string) int {
	v, _ := strconv.Atoi(b)
	return v
}

func createCountMap() map[rune]int {
	return map[rune]int{
		'A': 0,
		'K': 0,
		'Q': 0,
		'T': 0,
		'9': 0,
		'8': 0,
		'7': 0,
		'6': 0,
		'5': 0,
		'4': 0,
		'3': 0,
		'2': 0,
		'J': 0,
	}
}
