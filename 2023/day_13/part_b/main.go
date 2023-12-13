package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sumReflections(f))
}

func sumReflections(input io.Reader) int {
	puzzles := parsePuzzles(input)
	s := 0

	for _, p := range puzzles {
		origH:= horizontalIndexOfReflection(p, false)
		origV:= verticalIndexOfReflection(p, false)

		newH := horizontalIndexOfReflection(p, true)
		newV := verticalIndexOfReflection(p, true)

		// Only add the horizontal value if it is changed by the smudge.
		if origH != newH {
			s += newH * 100
		}

		// Only add the vertical value if it is changed by the smudge.
		if origV != newV {
			s += newV
		}
	}

	return s
}

func verticalIndexOfReflection(p [][]rune, errorCorrect bool) int {
	// Rotate
	p2 := make([][]rune, len(p[0]))
	// for range p[0] {
	// p2 = append(p2, make([]rune))
	// }

	for _, row := range p {
		for j, c := range row {
			p2[j] = append(p2[j], c)
		}
	}

	return horizontalIndexOfReflection(p2, errorCorrect)
}

func horizontalIndexOfReflection(p [][]rune, errorCorrect bool) int {
	// Find difference of 1 between two lines.
	// Flip the different position
	// Test for reflection
Outer:
	for i := 1; i < len(p); i++ {
		var smudged bool
		top := p[0:i]
		bottom := p[i:]

		var max int
		if len(top) > len(bottom) {
			max = len(bottom)
		} else {
			max = len(top)
		}

		for j := 0; j < max; j++ {
			var s bool
			var equal bool
			if errorCorrect {
				equal, s = equalIfOneFlipped(top[len(top)-1-j], bottom[j])
			} else {
				equal = slices.Equal(top[len(top)-1-j], bottom[j])
			}
			if !equal {
				continue Outer
			}

			// A smudge was previously found
			if smudged && s {
				continue Outer
			}

			smudged = s
		}

		return i

	}
	return 0
}

func equalIfOneFlipped(r1, r2 []rune) (bool, bool) {
	var foundDiff bool
	// fmt.Println(r1, r2)

	for i, c := range r1 {
		if c != r2[i] {
			// More than one flip
			if foundDiff {
				// fmt.Println("More than one difference")
				return false, false
			} else {
				foundDiff = true
			}
		}
	}

	return true, foundDiff
}

func parsePuzzles(input io.Reader) [][][]rune {
	var puzzles [][][]rune

	all, err := io.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}

	for _, ps := range strings.Split(string(all), "\n\n") {
		puzzles = append(puzzles, parsePuzzle(ps))
	}

	return puzzles
}

func parsePuzzle(s string) [][]rune {
	var p [][]rune

	for _, ls := range strings.Split(s, "\n") {
		p = append(p, parseLine(ls))
	}

	return p
}

func parseLine(s string) []rune {
	var l []rune

	for _, r := range s {
		l = append(l, r)
	}

	return l
}
