package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(CountAcceptedCombinations(f))
}

func CountAcceptedCombinations(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	workflows := map[string]Workflow{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		workflow := NewWorkflow(line)
		workflows[workflow.label] = workflow
	}

	return RunWorkflow(workflows, map[string]Range{}, "in", []string{})
}

type Workflow struct {
	defaultDest string
	rules       []Rule
	label       string
}

type Rule struct {
	variable string
	operator string
	count    int
	dest     string
}

func NewWorkflow(line string) Workflow {
	var w Workflow

	label, rest, _ := strings.Cut(line, "{")

	rs := strings.Split(rest[0:len(rest)-1], ",")

	for _, r := range rs[0 : len(rs)-1] {

		variable := r[0:1]
		operator := r[1:2]

		parts := strings.Split(r[2:], ":")
		count, _ := strconv.Atoi(parts[0])
		dest := parts[1]

		w.rules = append(w.rules, Rule{variable, operator, count, dest})
	}

	w.defaultDest = rs[len(rs)-1]
	w.label = label

	return w
}

type Range struct {
	min int
	max int
}

func RunWorkflow(workflows map[string]Workflow, rangeSet map[string]Range, workflowLabel string, chain []string) int {
	chain = append(chain, workflowLabel)
	if workflowLabel == "A" {
		c := 1
		for _, r := range rangeSet {
			c *= (r.max - r.min) // TODO: needs a +1?
			// fmt.Println(c)
		}
		// fmt.Println(chain, rangeSet, c)
		return c

	}

	if workflowLabel == "R" {
		return 0
	}

	w := workflows[workflowLabel]
	var combos int
	for _, rule := range w.rules {
		newRangeset := map[string]Range{}

		for k, v := range rangeSet {
			newRangeset[k] = v
		}

		if _, ok := newRangeset[rule.variable]; !ok {
			newRangeset[rule.variable] = Range{0, 4000}
		}

		r := newRangeset[rule.variable]
		// Merge in the new rule
		switch rule.operator {
		case ">":
			if rule.count > r.min {
				newRangeset[rule.variable] = Range{rule.count, r.max}
			}
		case "<":
			if rule.count < r.max {
				newRangeset[rule.variable] = Range{r.min, rule.count}
			}
		}

		combos += RunWorkflow(workflows, newRangeset, rule.dest, chain)
	}

	combos += RunWorkflow(workflows, rangeSet, w.defaultDest, chain)

	return combos
}
