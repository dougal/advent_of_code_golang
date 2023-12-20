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

	// Start at "in", follow each branch until reach "A" or "R"
	// Sum combinations

	return workflows["in"].Follow(workflows, map[string]Range{})
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

func (w Workflow) Follow(workflows map[string]Workflow, rangeSet map[string]Range) int {
	fmt.Println(w.label)
	// Loop through each rule
	// If A, return the combinations based on the partition rules
	// If R, return 0
	// Merge the partition with the rule
	// Call Follow again with the new partition
	// Sum the returns and return
	var combos int
	for _, rule := range w.rules {
		if rule.dest == "A" {

			c := 1
			for _, r := range rangeSet {
				c *= (r.max - r.min) // TODO: needs a +1?
				fmt.Println(c)
			}
			combos += c
			break

		}

		if rule.dest == "R" {
			combos += 0
			break
		}

		newRangeset := rangeSet

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

		combos += workflows[rule.dest].Follow(workflows, rangeSet)
	}

	// Handle the defaultDest
	if w.defaultDest == "A" {

		c := 1
		for _, r := range rangeSet {
			c *= (r.max - r.min) // TODO: needs a +1?
				fmt.Println(c)
		}
		combos += c

	} else if w.defaultDest == "R" {
		combos += 0
	} else {
		combos += workflows[w.defaultDest].Follow(workflows, rangeSet)
	}

	return combos
}
