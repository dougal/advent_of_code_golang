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

	for scanner.Scan() {
		l := NewLine(scanner.Text())
		c += arrangements(l)
	}

	return c
}

func arrangements(l Line) int {
	// Finished, no more to check.
	// Return 1 if all groups accounted for.
	if l.cells == "" {
		if len(l.grouping) == 0 {
			return 1
		} else {
			return 0
		}
	}

	// There are no more groups remaining
	// Return 1 if all wells accounted for.
	if len(l.grouping) == 0 {
		if strings.Index(l.cells, "#") == -1 {
			return 1
		} else {
			return 0
		}
	}

	c := 0

	if l.cells[0] == '.' || l.cells[0] == '?' {
		lNext := l
		lNext.cells = lNext.cells[1:]
		c += arrangements(lNext)
	}

	if l.cells[0] == '#' || l.cells[0] == '?' {
		// Valid if:
		// - Enough wells left to complete
		// - First chars equaling first group must be wells
		// - Next spring after must be operational, or be out of bounds.
		if l.grouping[0] <= len(l.cells) &&
			strings.Index(l.cells[:l.grouping[0]], ".") == -1 &&
			(l.grouping[0] == len(l.cells) || l.cells[l.grouping[0]] != '#') {
			lNext := l
			// Pass empty string if first group is same length as remaining cells
			if l.grouping[0] == len(l.cells) {
				lNext.cells = ""
			} else {
				lNext.cells = lNext.cells[l.grouping[0] + 1:]
			}
			lNext.grouping = lNext.grouping[1:]
			c += arrangements(lNext)
		}
	}

	return c
}

type Line struct {
	cells    string
	grouping []int
}

func NewLine(s string) Line {
	var l Line
	sc, gs, _ := strings.Cut(s, " ")

	l.cells = sc

	for _, g := range strings.Split(gs, ",") {
		gi, _ := strconv.Atoi(g)
		l.grouping = append(l.grouping, gi)
	}

	return l
}
