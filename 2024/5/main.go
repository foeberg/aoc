package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

type rule struct {
	a string
	b string
}

func (r rule) appliesTo(update []string) bool {
	return slices.Index(update, r.a) <= slices.Index(update, r.b)
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var rules []rule
	var updates [][]string

	parsingRules := true
	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if len(b) == 0 {
			parsingRules = false
			continue
		}

		if parsingRules {
			split := strings.Split(string(b), "|")
			rule := rule{split[0], split[1]}
			rules = append(rules, rule)

			continue
		}

		updates = append(updates, strings.Split(string(b), ","))
	}

	fmt.Printf("Part 1: %d\n", part1(rules, updates))
	fmt.Printf("Part 2: %d\n", part2(rules, updates))
}

func findApplicableRules(rules []rule, update []string) []rule {
	var applicableRules []rule

	for _, rule := range rules {
		if slices.Contains(update, rule.a) && slices.Contains(update, rule.b) {
			applicableRules = append(applicableRules, rule)
		}
	}

	return applicableRules
}

func part1(rules []rule, updates [][]string) int {
	var correctUpdates [][]string

	for _, update := range updates {
		applicableRules := findApplicableRules(rules, update)
		isCorrect := true

		for _, rule := range applicableRules {
			if !rule.appliesTo(update) {
				isCorrect = false
				break
			}
		}
		if isCorrect {
			correctUpdates = append(correctUpdates, update)
		}
	}

	sum := 0
	for _, update := range correctUpdates {
		mid := update[len(update)/2]
		midInt, _ := strconv.Atoi(mid)
		sum += midInt
	}
	return sum
}

func fixUpdate(rules []rule, update []string) {
	slices.SortStableFunc(update, func(a, b string) int {
		correct := true
		subSlice := []string{a, b}
		applicableRules := findApplicableRules(rules, subSlice)
		for _, rule := range applicableRules {
			if !rule.appliesTo(subSlice) {
				correct = false
				break
			}
		}

		if !correct {
			return 1
		}
		return -1
	})
}

func part2(rules []rule, updates [][]string) int {
	var incorrectUpdates [][]string

	for _, update := range updates {
		applicableRules := findApplicableRules(rules, update)
		isCorrect := true

		for _, rule := range applicableRules {
			if !rule.appliesTo(update) {
				isCorrect = false
				break
			}
		}
		if !isCorrect {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}

	for _, update := range incorrectUpdates {
		fixUpdate(rules, update)
	}

	sum := 0
	for _, update := range incorrectUpdates {
		mid := update[len(update)/2]
		midInt, _ := strconv.Atoi(mid)
		sum += midInt
	}

	return sum
}
