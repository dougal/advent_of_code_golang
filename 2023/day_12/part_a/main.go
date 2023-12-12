package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sumArrangements(f))
}

func sumArrangements(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	c := 0

	for scanner.Scan() {
		c += arrangements(scanner.Text())
	}

	return c
}

func arrangements(s string) int {
	var c int
	line := NewLine(s)

	permutations := int(math.Pow(float64(2), float64(line.unknownCount)))
	for i := 0; i < permutations; i++ {
		var arrangement []string
		for k := 0; k < line.unknownCount; k++ {
			arrangement = append(arrangement, ".")
		}

		var (
			n = i
			j int
		)
		
		for n > 0 {
			if n%2 == 0 {
				arrangement[j] = "."
			} else {
				arrangement[j] = "#"
			}

			n = n / 2
			j++
		}

		if line.satisfiedBy(arrangement) {
			c++
		}
	}

	return c
}

type Line struct {
	cells []string
	grouping []int
	unknownCount int
}

func NewLine(s string) Line {
	var l Line
	sc, gs, _ := strings.Cut(s, " ")

	l.cells = strings.Split(sc, "")

	for _, g := range strings.Split(gs, ",") {
		gi, _ := strconv.Atoi(g)
		l.grouping = append(l.grouping, gi)
	}

	for _, c := range l.cells {
		if c == "?" {
			l.unknownCount++
		}
	}

	return l
}

// return true if the arrangement, when applied, satisfies the grouping
func (l Line) satisfiedBy(arrangement []string) bool {
	var missingIndex int
	var cells []string

	for _, c := range l.cells {
		if c != "?" {
			cells = append(cells, c)
			continue
		}

		cells = append(cells, arrangement[missingIndex])
		missingIndex++
	}

	s := strings.Join(cells, "")

	gs := strings.Split(s, ".")

	var grouping []int
	for _, g := range gs {
		if g == "" {
			continue
		}

		grouping = append(grouping, len(g))
	}

	return slices.Equal(grouping, l.grouping)
}
