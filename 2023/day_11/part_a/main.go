package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sumGalaxyDistances(f))
}

type Cell int

const Empty Cell = 0
const Galaxy Cell = 1

type Space [][]Cell

func sumGalaxyDistances(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var space Space
	for scanner.Scan() {
		space.AddRow(scanner.Text())
	}

	// fmt.Println(space.Print())

	space.ExpandEmptyRows()
	space.ExpandEmptyCols()

	// fmt.Println(space.Print())

	return space.SumDistances()
}

func (s Space) Print() string {
	var a string
	for _, row := range s {
		for _, c := range row {
			if c == Empty {
				a += " . "
			} else {
				a += " # "
			}
		}
		a += "\n"
	}

	return a
}

func (s *Space) AddRow(line string) {
	var o []Cell

	for _, r := range line {
		switch r {
		case '.':
			o = append(o, Empty)
		case '#':
			o = append(o, Galaxy)
		}
	}

	*s = append(*s, o)
}

func (s *Space) ExpandEmptyRows() {
	var emptyRows []int

	for x, row := range *s {
		empty := true
		for _, c := range row {
			if c == Galaxy {
				empty = false
				break
			}
		}
		if empty {
			// Prepend so as to not have to adjust all the indexes when inserting.
			emptyRows = append([]int{x}, emptyRows...)
		}
	}

	// Duplicate each empty row
	for _, i := range emptyRows {
		*s = slices.Insert(*s, i, (*s)[i])
	}
}

func (s *Space) ExpandEmptyCols() {
	var emptyCols []int
	for i := range (*s)[0] {
		empty := true
		for _, row := range *s {
			if row[i] == Galaxy {
				empty = false
			}
		}

		if empty {
			// Prepend so as to not have to adjust all the indexes when inserting.
			emptyCols = append([]int{i}, emptyCols...)
		}
	}

	// Duplicate each empty column
	for _, i := range emptyCols {
		for x, row := range *s {
			(*s)[x] = slices.Insert(row, i, Empty)
		}
	}
}

func (s *Space) SumDistances() int {
	var coords [][]int

	for x, row := range *s {
		for y, g := range row {
			if g == Galaxy {
				coords = append(coords, []int{x, y})
			}
		}
	}

	var sum float64

	for i, c := range coords {
		if i == len(coords)-1 {
			break
		}
		for _, c2 := range coords[i+1:] {
			sum += math.Abs(float64(c[0]-c2[0])) + math.Abs(float64(c[1]-c2[1]))
		}
	}

	return int(sum)
}
