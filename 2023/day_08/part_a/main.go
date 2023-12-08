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

	node := "AAA"
	stepCount := 0
	currentStep := 0
	for node != "ZZZ" {
		node = nodes[node][steps[currentStep]]

		currentStep++
		stepCount++
		if currentStep == len(steps) {
			currentStep = 0
		}
	}

	return stepCount
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
