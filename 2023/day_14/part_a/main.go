package main

import (
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

	fmt.Println(totalLoad(f))
}

const roundRock rune = 'O'
const cubeRock rune = '#'
const space rune = '.'

func totalLoad(input io.Reader) int {
	field := parseRocks(input)
	s := 0

	// For each line
	for _, line := range field {
		currentObstacle := -1
		for i, c := range line {
			switch c {
			case cubeRock:
				currentObstacle = i
			case roundRock:
				currentObstacle++
				s += len(line) - currentObstacle
			}
		}
	}
	return s
}

func parseRocks(input io.Reader) [][]rune {
	s, err := io.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(s), "\n")

	field := make([][]rune, len(lines[0]))

	// Put rows into columns, so move west/left instead of to top
	for _, line := range lines {
		for i, c := range string(line) {
			field[i] = append(field[i], c)
		}
	}

	return field
}
