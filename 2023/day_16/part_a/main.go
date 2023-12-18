package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(CountEnergizedTiles(f))
}

func CountEnergizedTiles(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var matrix [][]rune

	for scanner.Scan() {
		row := []rune(scanner.Text())
		matrix = append(matrix, row)
	}

	visited := make([][]bool, len(matrix))
	for y := range visited {
		visited[y] = make([]bool, len(matrix[y]))
	}

	// x ->
	// y
	// |
	// v

	vecX, vecY := 1, 0
	x, y := -1, 0

	// Reset the cache
	MoveBeamHits = map[string]bool{}

	MoveBeam(matrix, &visited, x, y, vecX, vecY)

	var c int
	for _, row := range visited {
		for _, b := range row {
			if b {
				c++
			}
		}
	}

	return c
}

var MoveBeamHits = map[string]bool{}

func MoveBeam(matrix [][]rune, visited *[][]bool, x, y, vecX, vecY int) {
	key := fmt.Sprintf("%d,%d,%d,%d", x, y, vecX, vecY)

	if _, ok := MoveBeamHits[key]; ok {
		return
	} else {
		MoveBeamHits[key] = true
	}

	newX := x + vecX
	newY := y + vecY

	// Return if out of bounds
	if newX >= len(matrix[0]) ||
		newX < 0 ||
		newY >= len(matrix) ||
		newY < 0 {
		return
	}

	(*visited)[newY][newX] = true

	// TODO: Switch to having these return arg sets rather than calling MoveBeam
	// themselves.
	switch matrix[newY][newX] {
	case '.':
		MoveBeam(matrix, visited, newX, newY, vecX, vecY)
	case '-':
		dash(matrix, visited, newX, newY, vecX, vecY)
	case '|':
		pipe(matrix, visited, newX, newY, vecX, vecY)
	case '/':
		fSlash(matrix, visited, newX, newY, vecX, vecY)
	case '\\':
		bSlash(matrix, visited, newX, newY, vecX, vecY)
	}
}

func dash(matrix [][]rune, visited *[][]bool, x, y, vecX, vecY int) {
	// fmt.Println("dash", x, y, vecX, vecY)
	// Parallel, do nothing
	if vecY == 0 {
		MoveBeam(matrix, visited, x, y, vecX, vecY)
		return
	}

	// Perpendicular, split the beam
	// Left
	MoveBeam(matrix, visited, x, y, -1, 0)
	// Right
	MoveBeam(matrix, visited, x, y, 1, 0)
}

func pipe(matrix [][]rune, visited *[][]bool, x, y, vecX, vecY int) {
	// Parallel, do nothing
	if vecX == 0 {
		MoveBeam(matrix, visited, x, y, vecX, vecY)
		return
	}

	// Perpendicular, split the beam
	// Up
	MoveBeam(matrix, visited, x, y, 0, -1)
	// Down
	MoveBeam(matrix, visited, x, y, 0, 1)
}

func fSlash(matrix [][]rune, visited *[][]bool, x, y, vecX, vecY int) {
	// Right
	if vecX == 1 {
		// Up
		MoveBeam(matrix, visited, x, y, 0, -1)
		return
	}

	// Left
	if vecX == -1 {
		// Down
		MoveBeam(matrix, visited, x, y, 0, 1)
		return
	}

	// Down
	if vecY == 1 {
		// Left
		MoveBeam(matrix, visited, x, y, -1, 0)
		return
	}

	// Up
	// Right
	MoveBeam(matrix, visited, x, y, 1, 0)
}

func bSlash(matrix [][]rune, visited *[][]bool, x, y, vecX, vecY int) {
	// Right
	if vecX == 1 {
		// Down
		MoveBeam(matrix, visited, x, y, 0, 1)
		return
	}

	// Left
	if vecX == -1 {
		// Up
		MoveBeam(matrix, visited, x, y, 0, -1)
		return
	}

	// Down
	if vecY == 1 {
		// Right
		MoveBeam(matrix, visited, x, y, 1, 0)
		return
	}

	// Up
	// Left
	MoveBeam(matrix, visited, x, y, -1, 0)
}
