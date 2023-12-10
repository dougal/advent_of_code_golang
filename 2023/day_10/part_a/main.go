package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(maxDistance(f))
}

func maxDistance(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var m Map

	for scanner.Scan() {
		m = append(m, parseMapLine(scanner.Text()))
	}

	return loopDistance(m) / 2
}

type Form rune

const startPipe Form = 'S'
const vPipe Form = '|'
const hPipe Form = '-'
const lPipe Form = 'L'
const jPipe Form = 'J'
const sevenPipe Form = '7'
const fPipe Form = 'F'
const noPipe Form = '.'

type Pipe struct {
	visited bool
	form    Form
}

type Map [][]Pipe

func parseMapLine(line string) []Pipe {
	var pipes []Pipe

	for _, r := range line {
		pipes = append(pipes, Pipe{form: Form(r)})
	}

	return pipes
}

func loopDistance(m Map) int {
	x, y := m.findStartingPosition()
	x, y = m.nextValidMove(x, y)
	d := 1

	for x != -1 && y != -1 {
		x, y = m.nextValidMove(x, y)
		d++
	}

	return d
}

func (m Map) findStartingPosition() (int, int) {
	for y, r := range m {
		for x, p := range r {
			if p.form == startPipe {
				return x, y
			}
		}
	}

	// Bad things
	return -1, -1
}

type Move struct {
	x int
	y int
}

var Up = Move{0, -1}
var Down = Move{0, 1}
var Left = Move{-1, 0}
var Right = Move{1, 0}

var validMoves = map[Form][]Move{
	startPipe: {Up, Down, Left, Right},
	vPipe:     {Up, Down},
	hPipe:     {Left, Right},
	lPipe:     {Up, Right},
	jPipe:     {Up, Left},
	sevenPipe: {Left, Down},
	fPipe:     {Right, Down},
	noPipe:    {},
}

var validForms = map[Move][]Form{
	Up:    {vPipe, lPipe, jPipe},
	Down:  {vPipe, sevenPipe, fPipe},
	Left:  {hPipe, jPipe, sevenPipe},
	Right: {hPipe, lPipe, fPipe},
}

func (m *Map) nextValidMove(x, y int) (int, int) {
	curPipe := (*m)[y][x]
	// ignore if already visited
	for _, move := range validMoves[curPipe.form] {
		nextX := x + move.x
		nextY := y + move.y

		// Outside of map bounds
		if nextX < 0 || nextY < 0 || nextX >= len((*m)[0]) || nextY >= len(*m) {
			continue
		}

		nextPipe := (*m)[y+move.y][x+move.x]

		if nextPipe.visited {
			fmt.Println("visited")
			continue
		}

		// Skip if not a valid form to move to.
		if !slices.Contains(validMoves[nextPipe.form], Move{move.x * -1, move.y * -1}) {
			continue
		}

		// fmt.Printf("%d, %d, %s, %t\n", x, y, string(nextPipe.form), nextPipe.visited)

		nextPipe.visited = true
		(*m)[y+move.y][x+move.x] = nextPipe

		return nextX, nextY
	}

	// Back at start
	return -1, -1
}
