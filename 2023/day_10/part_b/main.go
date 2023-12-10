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

	fmt.Println(countEnclosedTiles(f))
}

func countEnclosedTiles(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var m Map

	for scanner.Scan() {
		m = append(m, parseMapLine(scanner.Text()))
	}

	m.markLoopTiles()
	m.markOutside()
	// m.markInside() // Makes debugging easier
	fmt.Println(m.Print())
	return m.countInsideTiles()
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
const inside Form = 'I'
const outside Form = 'O'

type Pipe struct {
	isStart   bool
	isLoop    bool
	// isOutside bool
	form      Form
}

type Map [][]Pipe

func parseMapLine(line string) []Pipe {
	var pipes []Pipe

	for _, r := range line {
		pipes = append(pipes, Pipe{form: Form(r)})
	}

	return pipes
}

func (m *Map) markLoopTiles() {
	startX, startY := m.findStartingPosition()
	// Mark the start
	(*m)[startY][startX].isStart = true

	// NOTE: Did this manually looking at input
	// (*m)[startY][startX].form = jPipe

	x, y := m.nextValidMove(startX, startY)

	for x != -1 && y != -1 {
		x, y = m.nextValidMove(x, y)
	}
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

func (m Map) width() int {
	return len(m[0])
}

func (m Map) height() int {
	return len(m)
}

type Move struct {
	x int
	y int
}

var UpLeft = Move{-1, -1}
var Up = Move{0, -1}
var UpRight = Move{1, -1}
var Left = Move{-1, 0}
var Right = Move{1, 0}
var DownLeft = Move{-1, 1}
var Down = Move{0, 1}
var DownRight = Move{1, 1}

var AllMoves = []Move{UpLeft, Up, UpRight, Left, Right, DownLeft, Down, DownRight}

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
		if nextX < 0 || nextY < 0 || nextX >= m.width() || nextY >= m.height() {
			continue
		}

		nextPipe := (*m)[y+move.y][x+move.x]

		if nextPipe.isLoop {
			// fmt.Println("visited")
			continue
		}

		// Skip if not a valid form to move to.
		if !slices.Contains(validMoves[nextPipe.form], Move{move.x * -1, move.y * -1}) {
			continue
		}

		// fmt.Printf("%d, %d, %s, %t\n", x, y, string(nextPipe.form), nextPipe.visited)

		nextPipe.isLoop = true
		(*m)[y+move.y][x+move.x] = nextPipe

		return nextX, nextY
	}

	// Back at start
	return -1, -1
}

func (m *Map) markOutside() {
	for y, r := range (*m) {
		tubeCount := 0
		var previousForm Form
		for x, tile := range r {
			// Only modify inside/outside if part of loop
			// Ignore horizontal pipes as they do not affect inside/outside
      if tile.isLoop && tile.form != hPipe {
				// TODO Replace the start letter with it's actual type
        tubeCount++
				// j+f and 7+l cancel each other out, so don't alter the tube count.
        if ((tile.form == jPipe && previousForm == fPipe) || (tile.form == sevenPipe && previousForm == lPipe)) {
          tubeCount--
        }

        previousForm = tile.form
      }

			if tile.isLoop {
				continue
			}

      if tubeCount % 2 == 1 {
				(*m)[y][x].form = inside
			} else {
				(*m)[y][x].form = outside
      }
    }
  }
}

// func (m *Map) markOutside() {
// 	// Top row
// 	for x := range (*m)[0] {
// 		m.markTileOutside(x, 0)
// 	}

// 	// Bottom row
// 	for x := range (*m)[m.height()-1] {
// 		m.markTileOutside(x, m.height()-1)
// 	}

// 	// Left column
// 	for y := range *m {
// 		m.markTileOutside(0, y)
// 	}

// 	// Right column
// 	for y := range *m {
// 		m.markTileOutside(m.width()-1, y)
// 	}

// 	m.markOutsideOfCorners()
// }

// var outsideCorners = map[Form][]Move{
// 	lPipe:     {Down, Left},
// 	jPipe:     {Down, Right},
// 	sevenPipe: {Up, Right},
// 	fPipe:     {Up, Left},
// }

// func (m *Map) markOutsideOfCorners() {
// 	// Loop through every tile
// 	for y, r := range *m {
// 		for x, t := range r {
// 			// If a corner tile
// 			if moves, ok := outsideCorners[t.form]; ok {
// 				// Loop through the tiles on the outside
// 				for _, move := range moves {
// 					nextX := x + move.x
// 					nextY := y + move.y

// 					// Outside of map bounds
// 					if nextX < 0 || nextY < 0 || nextX >= m.width() || nextY >= m.height() {
// 						continue
// 					}

// 					m.markTileOutside(nextX, nextY)
// 				}
// 			}
// 		}
// 	}
// }

// func (m *Map) markTileOutside(x, y int) {
// 	tile := (*m)[y][x]

// 	if tile.isLoop {
// 		return
// 	}

// 	if tile.isOutside {
// 		return
// 	}

// 	// Mark the tile as outside
// 	tile.isOutside = true
// 	tile.form = outside
// 	(*m)[y][x] = tile

// 	// Search all it's neighbours
// 	for _, move := range AllMoves {
// 		nextX := x + move.x
// 		nextY := y + move.y

// 		// Outside of map bounds
// 		if nextX < 0 || nextY < 0 || nextX >= m.width() || nextY >= m.height() {
// 			continue
// 		}

// 		m.markTileOutside(nextX, nextY)
// 	}
// }

func (m Map) countInsideTiles() int {
	var c int

	for _, r := range m {
		for _, t := range r {
			if t.form == inside {
				c++
			}
		}
	}

	return c
}

// func (m *Map) markInside() {
// 	for y, r := range *m {
// 		for x, t := range r {
// 			if t.form != inside && !t.isLoop {
// 				(*m)[y][x].form = inside
// 			}
// 		}
// 	}
// }

func (m Map) Print() string {
	var s string

	for _, r := range m {
		for _, t := range r {
			s += "" + string(t.form)
		}

		s += "\n"
	}

	return s
}
