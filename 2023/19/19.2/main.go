package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	param    string
	operator string
	value    int
	result   string
	not      bool
}

type workflow struct {
	rules []rule
}

func main() {
	workflows := parseInput()

	fmt.Println(follow("in", make([]rule, 0), workflows))
}

func calculateCombinations(rules []rule) int {
	min := make(map[string]int)
	max := make(map[string]int)
	for _, rule := range rules {
		if rule.param == "" {
			continue
		}

		if rule.operator == ">" {
			if rule.not {
				max[rule.param] = rule.value
				continue
			}
			min[rule.param] = rule.value + 1
			continue
		}

		if rule.not {
			min[rule.param] = rule.value
			continue
		}
		max[rule.param] = rule.value - 1
	}

	combs := 1
	params := []string{"x", "m", "a", "s"}
	for _, param := range params {
		max := max[param]
		min := min[param]
		if max == 0 {
			max = 4000
		}
		if min == 0 {
			min = 1
		}

		combs = combs * ((max - min) + 1)
	}

	return combs
}

func follow(flowID string, rules []rule, workflows map[string]workflow) int {
	if flowID == "R" {
		return 0
	}
	if flowID == "A" {
		return calculateCombinations(rules)
	}

	workflow := workflows[flowID]
	combs := 0
	for i, rule := range workflow.rules {
		prevRules := rules
		if i > 0 {
			for j := 0; j < i; j++ {
				r := workflow.rules[j]
				r.not = true
				prevRules = append(prevRules, r)
			}
		}
		prevRules = append(prevRules, rule)
		combs += follow(rule.result, prevRules, workflows)
	}

	return combs
}

func parseInput() map[string]workflow {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	workflows := make(map[string]workflow)

	for {
		b, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		if len(b) == 0 {
			break
		}

		line := string(b)
		i := strings.IndexRune(line, '{')

		name := line[:i]

		rulesStrs := strings.Split(line[i+1:len(line)-1], ",")

		var rules []rule
		for i := 0; i < len(rulesStrs); i++ {
			if i == len(rulesStrs)-1 {
				rules = append(rules, rule{result: rulesStrs[i]})
				break
			}
			str := rulesStrs[i]

			colonIndex := strings.Index(rulesStrs[i], ":")
			value, _ := strconv.Atoi(str[2:colonIndex])

			rules = append(rules, rule{
				param:    str[0:1],
				operator: str[1:2],
				value:    value,
				result:   str[colonIndex+1:],
			})
		}

		workflows[name] = workflow{rules}
	}

	return workflows
}
