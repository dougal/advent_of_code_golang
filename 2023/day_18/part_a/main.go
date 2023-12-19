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

	fmt.Println(LavaVolume(f))
}

func LavaVolume(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var steps []Step

	for scanner.Scan() {
		steps = append(steps, NewStep(scanner.Text()))
	}

	var x, y, maxUp, maxDown, maxLeft, maxRight int

	for _, s := range steps {
		switch s.direction {
		case "U":
			x -= s.distance
			if x < maxUp {
				maxUp = x
			}
		case "D":
			x += s.distance
			if x > maxDown {
				maxDown = x
			}
		case "L":
			y -= s.distance
			if y < maxLeft {
				maxLeft = y
			}
		case "R":
			y += s.distance
			if y > maxRight {
				maxRight = y
			}
		}
	}

	// Work out the matrix size, and the start position
	width := maxRight - maxLeft + 1
	height := maxDown - maxUp + 1
	y = maxUp * -1
	x = maxLeft * -1

	matrix := make([][]rune, height)
	for i := range matrix {
		matrix[i] = make([]rune, width)
		for j := range matrix[i] {
			matrix[i][j] = '.'
		}
	}

	// Dig the initial hole
	matrix[y][x] = '#'

	for _, s := range steps {
		switch s.direction {
		case "U":
			for i := 0; i < s.distance; i++ {
				y--
				matrix[y][x] = '#'
			}
		case "D":
			for i := 0; i < s.distance; i++ {
				y++
				matrix[y][x] = '#'
			}
		case "L":
			for i := 0; i < s.distance; i++ {
				x--
				matrix[y][x] = '#'
			}
		case "R":
			for i := 0; i < s.distance; i++ {
				x++
				matrix[y][x] = '#'
			}
		}
	}
	PrintMatrix(matrix)

	// Mark outside
	// Left side
	for i := 0; i < len(matrix); i++ {
		matrix = MarkOutside(matrix, 0, i)
	}

	// Right side
	for i := 0; i < len(matrix); i++ {
		matrix = MarkOutside(matrix, len(matrix[0]) - 1, i)
	}

	// Top side
	for i := 0; i < len(matrix[0]); i++ {
		matrix = MarkOutside(matrix, i, 0)
	}

	// Bottom side
	for i := 0; i < len(matrix[0]); i++ {
		matrix = MarkOutside(matrix, i, len(matrix)-1)
	}

	PrintMatrix(matrix)

	// Count total - outside
	var s int
	for _, row := range matrix {
		for _, c := range row {
			if c != 'O' {
				s++
			}
		}
	}

	return s
}

func PrintMatrix(matrix [][]rune) {
	for _, row := range matrix {
		fmt.Println(string(row))
	}
}

func MarkOutside(matrix [][]rune, x, y int) [][]rune {
	// Out of bounds
	if y >= len(matrix) || y < 0 || x >= len(matrix[0]) || x < 0 {
		return matrix
	}

	c := matrix[y][x]

	if c == '#' || c == 'O' {
		return matrix
	}

	matrix[y][x] = 'O'

	// Above
	matrix = MarkOutside(matrix, x, y-1)
	// Below
	matrix = MarkOutside(matrix, x, y+1)
	// Left
	matrix = MarkOutside(matrix, x-1, y)
	// Right
	matrix = MarkOutside(matrix, x+1, y)

	return matrix
}

type Step struct {
	direction string
	distance  int
	// colour    string
}

func NewStep(in string) Step {
	s := Step{}
	parts := strings.Split(in, " ")

	s.direction = parts[0]

	disti, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}

	s.distance = disti

	// TODO parse colour

	return s
}
