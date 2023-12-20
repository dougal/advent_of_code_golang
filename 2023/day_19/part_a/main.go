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

	fmt.Println(SumAccepted(f))
}

func SumAccepted(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	workflows := map[string]Workflow{}
	var inParts bool
	var s int

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			inParts = true
			continue
		}

		if inParts {
			s += AcceptedPartRating(workflows, line)
		} else {
		  label, workflow := NewWorkflow(line)
			workflows[label] = workflow
		}
	}

	return s
}

type Workflow struct {
  defaultDest string
  rules []Rule
}

type Rule struct {
  variable string
  operator string
	count int
	dest string
}

func NewWorkflow(line string) (string, Workflow) {
	var w Workflow

  label, rest, _ := strings.Cut(line, "{")

	rs := strings.Split(rest[0:len(rest) - 1], ",")

	for _, r := range rs[0:len(rs)-1] {

    variable := r[0:1]
    operator := r[1:2]

		parts := strings.Split(r[2:], ":")
		count, _ := strconv.Atoi(parts[0])
		dest := parts[1]

		w.rules = append(w.rules, Rule{variable, operator, count, dest})
	}

	w.defaultDest = rs[len(rs) - 1]

  
	return label, w
}

func (w Workflow) CheckPart(p map[string]int) string {
  // Loop through each rule, return dest if matches
	for _, rule :=  range w.rules {
	  // Check if part has label
		if _, ok := p[rule.variable]; !ok {
		  continue
		}

    switch rule.operator {
		case ">":
		  if p[rule.variable] > rule.count {
			  return rule.dest
			}
		case "<":
		  if p[rule.variable] < rule.count {
			  return rule.dest
			}
		}
	}

	// Else return the default
	return w.defaultDest
}

func AcceptedPartRating(workflows map[string]Workflow, line string) int {
	part := map[string]int{}
	var totalRating int
	pairs := strings.Split(line[1:len(line) - 1], ",")

	for _, pair := range pairs {
	  label, countS, _ := strings.Cut(pair, "=")

		count, _ := strconv.Atoi(countS)
		totalRating += count

		part[label] = count
	}

	workflowLabel := "in"
	for workflowLabel != "A" && workflowLabel != "R" {
	  workflowLabel = workflows[workflowLabel].CheckPart(part)
	}

  if workflowLabel == "A" {
    return totalRating
	}

	return 0
}
