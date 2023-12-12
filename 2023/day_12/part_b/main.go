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

	fmt.Println(sumArrangements(f))
}

func sumArrangements(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	c := 0

	// var l int
	for scanner.Scan() {
		// fmt.Println(l)
		c += arrangements(scanner.Text())
		// l++
	}

	return c
}

func arrangements(s string) int {
	var c int
	line := NewLine(s)

	// permutations := int(math.Pow(float64(2), float64(line.unknownCount)))
	permutations := maxPermutation(line.unknownCount, line.brokenCount)
	fmt.Println(permutations)
Outer:
	for i := 0; i < permutations; i++ {
		arrangement := make([]string, line.unknownCount)

		var (
			n           = i
			j           int
			brokenAdded int
		)

		for n > 0 {
			if n%2 == 0 {
				arrangement[j] = "."
			} else {
				arrangement[j] = "#"
				brokenAdded++
				// fmt.Printf("%d > %d\n", brokenAdded, line.unknownBrokenCount)
				if brokenAdded > line.unknownBrokenCount {
					continue Outer
				}
			}

			n = n / 2
			j++
		}

		// fmt.Println(arrangement)
		if line.satisfiedBy(arrangement) {
			c++
		}
	}

	return c
}

type Line struct {
	cells              []string
	grouping           []int
	unknownCount       int
	brokenCount        int
	knownBrokenCount   int
	unknownBrokenCount int
}

func NewLine(s string) Line {
	var l Line
	sc, gs, _ := strings.Cut(s, " ")

	cellGroup := strings.Split(sc, "")
	for i := 0; i < 5; i++ {
		l.cells = append(l.cells, cellGroup...)
		l.cells = append(l.cells, "?")
	}

	var gg []int
	for _, g := range strings.Split(gs, ",") {
		gi, _ := strconv.Atoi(g)
		gg = append(gg, gi)
	}
	for i := 0; i < 5; i++ {
		l.grouping = append(l.grouping, gg...)
	}

	for _, c := range l.grouping {
		l.brokenCount += c
	}

	for _, c := range l.cells {
		if c == "?" {
			l.unknownCount++
		} else if c == "#" {
			l.knownBrokenCount++
		}
	}

	l.unknownBrokenCount = l.brokenCount - l.knownBrokenCount

	// fmt.Println(l.cells)
	// fmt.Println(l.grouping)
	// fmt.Println(l.brokenCount)
	// fmt.Println(l.knownBrokenCount)
	// fmt.Println(l.unknownBrokenCount)

	return l
}

// return true if the arrangement, when applied, satisfies the grouping
func (l Line) satisfiedBy(arrangement []string) bool {
	// Never satisfied if not enough broken wells
	var brokenCount int
	for _, c := range arrangement {
		if c == "#" {
			brokenCount++
		}
	}

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

	var gc int
	for _, g := range gs {
		if g == "" {
			continue
		}

		if len(g) != l.grouping[gc] {
			return false
		}
		gc++
	}

	return true
}

func maxPermutation(unknownCount, brokenCount int) int {
	// If 6 unknown, and 3 broken:
	// ...###

	var s int

	m := 1
	for i := 0; i < unknownCount; i++ {
		if i < unknownCount-brokenCount {
			continue
		}

		s += m
		m *= 2
	}

	return s
}
