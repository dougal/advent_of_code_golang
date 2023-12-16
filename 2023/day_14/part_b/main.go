package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
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
	north := parseField(input)
	s := 0

	// Spin from North to West, as Westward is how the tilt calculation works.
	east := rotateClockwise(north)
	south := rotateClockwise(east)
	west := rotateClockwise(south)
	lastWest := west

	for i := 0; i < cycles; i++ {
		north = tiltAndRotateField(west)
		east = tiltAndRotateField(north)
		south = tiltAndRotateField(east)
		west = tiltAndRotateField(east)

		// Stable loop
		for i, line := range west {
      if !slices.Equal(line, lastWest[i]) {
				continue
			}
		}
		break
	}

	for _, line := range west {
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

	field = rotateClockwise(field)
	field = tiltField(field)

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

func rotateClockwise(field [][]rune) [][]rune {
	newField := make([][]rune, len(field[0]))

	for _, row := range field {
		for j, c := range row {
			newField[j] = append([]rune{c}, newField[j]...)
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

func parseField(input io.Reader) [][]rune {
	s, err := io.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(s), "\n")

	var field [][]rune

	for _, line := range lines {
		var row []rune
		for _, c := range string(line) {
			row = append(row, c)
		}
		field = append(field, row)
	}

	return field
}

func FieldToString(field [][]rune) string {
	var s string

	for _, row := range field {
		s += string(row) + "\n"
	}

	return s
}
