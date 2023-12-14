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

const cycles int = 1_000_000_000

// const cycles int = 1000
const roundRock rune = 'O'
const cubeRock rune = '#'
const space rune = '.'

func totalLoad(input io.Reader) int {
	field := parseRocks(input)
	s := 0

	for i := 0; i < cycles; i++ {
		if i%1000_000 == 0 {
			fmt.Println(i, hits, misses)
		}
		field = tiltAndRotateField(field)
	}

	for _, line := range field {
		for i, c := range line {
			switch c {
			case roundRock:
				s += len(line) - i
			}
		}
	}

	return s
}

var tiltAndRotateFieldCache = map[string][][]rune{}

func tiltAndRotateField(field [][]rune) [][]rune {
	cacheKey := fieldCacheKey(field)
	if v, ok := tiltAndRotateFieldCache[cacheKey]; ok {
		hits++
		return v
	}
	misses++

	field = tiltField(field)
	field = rotateField(field)

	tiltAndRotateFieldCache[cacheKey] = field

	return field
}

var hits int
var misses int

func fieldCacheKey(field [][]rune) string {
	var k string

	for _, l := range field {
		k += string(l)
	}

	return k
}

func rotateField(field [][]rune) [][]rune {
	newField := make([][]rune, len(field[0]))

	for _, row := range field {
		for j, c := range row {
			newField[j] = append(newField[j], c)
		}
	}

	return newField
}

func tiltField(field [][]rune) [][]rune {
	for i, line := range field {
		field[i] = tiltLine(line)
	}

	return field
}

func tiltLine(line []rune) []rune {
	currentObstacle := -1

	for j, c := range line {
		switch c {
		case cubeRock:
			currentObstacle = j
		case roundRock:
			currentObstacle++
			if currentObstacle != j {
				line[currentObstacle] = roundRock
				line[j] = space
			}
		}
	}

	return line
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
