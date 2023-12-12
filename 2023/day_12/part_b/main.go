package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
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
	l := 0

	for scanner.Scan() {
		c += arrangements(scanner.Text())
		l++
		fmt.Println(l)
	}

	return c
}

func arrangements(s string) int {
	var c int
	line := NewLine(s)

	c += replaceNextUnknown(line, "#")
	c += replaceNextUnknown(line, ".")

	return c
}

func replaceNextUnknown(l Line, r string) int {
	before, after, _ := strings.Cut(l.cells, "?")

	l.cells = before + r + after

	if strings.Index(l.cells, "?") == -1 {
		if l.validPrefix() {
			return 1
		} else {
			return 0
		}
	}

	var c int
	if l.validPrefix() {
		c += replaceNextUnknown(l, "#")
		c += replaceNextUnknown(l, ".")
	}

	return c
}

type Line struct {
	cells    string
	grouping []int
	brokenCount int
	workingCount int
}

func NewLine(s string) Line {
	var l Line
	sc, gs, _ := strings.Cut(s, " ")

	for i := 0; i < 5; i++ {
		if i > 0 {
			l.cells += "?"
		}
		l.cells += sc
	}

	for i := 0; i < 5; i++ {
		// Split goes brrrrrrr 5 times
		for _, g := range strings.Split(gs, ",") {
			gi, _ := strconv.Atoi(g)
			l.grouping = append(l.grouping, gi)
			l.brokenCount+= gi
		}
	}

	l.workingCount = len(l.cells) - l.brokenCount

	return l
}

// Returns false if no extend prefix can satisfy the groupings.
// Returns true if:
// - There are no groupings in the prefix
// - Any groupings in the prefix partially satisfy the groupings
// TODO: Forward-checking there are enough ? left to insert enough # or . to satisfy.
func (l Line) validPrefix() bool {
	prefix, _, _ := strings.Cut(l.cells, "?")
	var groups []int
	var curGroup int
	for _, c := range prefix {
		if c == '.' {
			if curGroup > 0 {
				groups = append(groups, curGroup)
				curGroup = 0
			}
		} else if c == '#' {
			curGroup++
		}
	}

	if curGroup > 0 {
		groups = append(groups, curGroup)
		curGroup = 0
	}

	// Too many groups
	if len(groups) > len(l.grouping) {
		return false
	}

	for i, g := range groups {
		isLast := i == len(groups)-1

		if isLast {
			if g > l.grouping[i] {
				return false
			}
		} else {
			if g != l.grouping[i] {
				return false
			}
		}
	}

	if strings.Index(l.cells, "?") == -1 {
		return slices.Equal(groups, l.grouping)
	}

	return true
}
