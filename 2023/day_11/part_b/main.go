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

	emptyRows := space.CalculateEmptyRows()
	emptyCols := space.CalculateEmptyCols()

	return space.SumDistances(emptyRows, emptyCols)
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

func (s *Space) CalculateEmptyRows() []int {
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
			emptyRows = append(emptyRows, x)
		}
	}

	return emptyRows
}

func (s *Space) CalculateEmptyCols() []int {
	var emptyCols []int
	for i := range (*s)[0] {
		empty := true
		for _, row := range *s {
			if row[i] == Galaxy {
				empty = false
			}
		}

		if empty {
			emptyCols = append(emptyCols, i)
		}
	}

	return emptyCols
}

func (s *Space) SumDistances(emptyRows, emptyCols []int) int {
	var coords [][]int

	for x, row := range *s {
		for y, g := range row {
			if g == Galaxy {
				var extraRows int
				for _, er := range emptyRows {
					if er < x {
						extraRows+=1000000-1
					}
				}

				var extraCols int
				for _, ec := range emptyCols {
					if ec < y {
						extraCols+=1000000-1
					}
				}

				coords = append(coords, []int{x+extraRows, y+extraCols})
			}
		}
	}

	var sum int

	for i, c := range coords {
		if i == len(coords)-1 {
			break
		}
		for _, c2 := range coords[i+1:] {
			sum += Abs(c[0]-c2[0]) + Abs(c[1]-c2[1])
		}
	}

	return sum
}

func Abs(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}
