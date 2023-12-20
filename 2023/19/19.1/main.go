package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	param    string
	operator string
	value    int
	result   string
}

func (r rule) pass(i item) bool {
	if r.param == "" {
		return true
	}

	if r.operator == "<" {
		return i[r.param] < r.value
	}

	return i[r.param] > r.value
}

type workflow struct {
	rules []rule
}

func (w workflow) handle(i item) string {
	for _, rule := range w.rules {
		if rule.pass(i) {
			return rule.result
		}
	}
	return ""
}

type item map[string]int

func main() {
	workflows, items := parseInput()

	var acceptedItems []item

	for _, item := range items {
		flowID := "in"
		for {
			res := workflows[flowID].handle(item)
			if res == "A" {
				acceptedItems = append(acceptedItems, item)
				break
			}
			if res == "R" {
				break
			}

			flowID = res
		}
	}

	var sum int
	for _, item := range acceptedItems {
		sum += item["x"] + item["m"] + item["a"] + item["s"]
	}

	fmt.Println(sum)
}

func parseInput() (map[string]workflow, []item) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	workflows := make(map[string]workflow)
	var items []item

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

	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		s := string(b)
		s = strings.TrimPrefix(s, "{")
		s = strings.TrimSuffix(s, "}")

		split := strings.Split(s, ",")

		item := make(item)

		for _, sp := range split {
			val, _ := strconv.Atoi(sp[2:])
			item[sp[0:1]] = val
		}

		items = append(items, item)
	}

	return workflows, items
}
