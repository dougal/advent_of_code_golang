package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(countSteps(f))
}

func countSteps(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	steps := parseSteps(scanner.Text())

	// Blank line
	scanner.Scan()

	nodes := map[string][]string{}
	for scanner.Scan() {
		label, dirs := parseNode(scanner.Text())
		nodes[label] = dirs
	}

	var startingNodes []string
	for label := range nodes {
		if label[2] == 'A' {
			startingNodes = append(startingNodes, label)
		}
	}

	var pathSteps []int

	for _, startingNode := range startingNodes {
		fmt.Println(startingNode)
		node := startingNode
		stepCount := 0
		currentStep := 0
		for node[2] != 'Z' {
			fmt.Println(node)
			node = nodes[node][steps[currentStep]]

			currentStep++
			stepCount++
			if currentStep == len(steps) {
				currentStep = 0
			}
		}

		pathSteps = append(pathSteps, stepCount)
	}

	lcm := pathSteps[0]

	for _, n := range pathSteps[1:] {
		lcm = LCM(lcm, n)
	}

	return lcm
}

func parseNode(line string) (string, []string) {
	label, rest, _ := strings.Cut(line, " = ")

	dirs := []string{rest[1:4], rest[6:9]}

	return label, dirs
}

func parseSteps(line string) []int {
	var steps []int

	for _, r := range line {
		if r == 'L' {
			steps = append(steps, 0)
		} else if r == 'R' {
			steps = append(steps, 1)
		} else {
			log.Fatalln("Whoops, no such direction.")
		}
	}

	return steps
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
