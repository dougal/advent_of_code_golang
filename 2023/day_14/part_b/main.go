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

type Field [][]rune

func totalLoad(input io.Reader) int {
	field := parseField(input)

	var checkPoint Field
	for i := 0; i < cycles; i++ {
		checkPoint = field

		field.TiltNorth()
		field.TiltWest()
		field.TiltSouth()
		field.TiltEast()

		// Stable loop
		if field.EqualTo(checkPoint) {
			break
		}
	}

	return field.LoadOnNorth()
}

func (f Field) EqualTo(other Field) bool {
	return false
}

func (f Field) LoadOnNorth() int {
	// TODO: Re-adjust for north.
	s := 0
	// for _, line := range field {
	// 	for i, c := range line {
	// 		switch c {
	// 		case roundRock:
	// 			s += len(line) - i
	// 		}
	// 	}
	// }

	return s
}

func (f *Field) TiltWest() {
	for i, line := range *f {
		currentObstacle := -1

		for j := 0; j < len(line); j++ {
			c := line[j]
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

		(*f)[i] = line
	}
}

func (f *Field) TiltEast() {
	for i, line := range *f {
		currentObstacle := len(line)

		for j := len(line) - 1; j >= 0; j-- {
			c := line[j]
			switch c {
			case cubeRock:
				currentObstacle = j
			case roundRock:
				currentObstacle--
				if currentObstacle != j {
					line[currentObstacle] = roundRock
					line[j] = space
				}
			}
		}

		(*f)[i] = line
	}
}

func (f *Field) TiltNorth() {
	for x := 0; x < len((*f)[0]); x++ {
		currentObstacle := -1

		for y := 0; y < len((*f)); y++ {
			c := (*f)[y][x]
			switch c {
			case cubeRock:
				currentObstacle = y
			case roundRock:
				currentObstacle++
				if currentObstacle != y {
					(*f)[currentObstacle][x] = roundRock
					(*f)[y][x] = space
				}
			}
		}
	}
}

func (f *Field) TiltSouth() {
	for x := 0; x < len((*f)[0]); x++ {
		currentObstacle := len((*f))

		for y := len((*f)) - 1; y >= 0; y-- {
			c := (*f)[y][x]
			switch c {
			case cubeRock:
				currentObstacle = y
			case roundRock:
				currentObstacle--
				if currentObstacle != y {
					(*f)[currentObstacle][x] = roundRock
					(*f)[y][x] = space
				}
			}
		}
	}
}

func parseField(input io.Reader) Field {
	s, err := io.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(s), "\n")

	var f Field

	for _, line := range lines {
		var row []rune
		for _, c := range string(line) {
			row = append(row, c)
		}
		f = append(f, row)
	}

	return f
}

func (f Field) String() string {
	var s string

	for _, row := range f {
		s += string(row) + "\n"
	}

	return s
}
