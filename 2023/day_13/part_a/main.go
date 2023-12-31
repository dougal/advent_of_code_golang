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
		s += checkReflectionHorizontal(p)
		s += checkReflectionVertical(p)
	}

	return s
}

func checkReflectionVertical(p [][]rune) int {
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

	return checkReflectionHorizontal(p2) / 100
}

func checkReflectionHorizontal(p [][]rune) int {
Outer:
	for i := 1; i < len(p); i++ {
		top := p[0:i]
		bottom := p[i:]

		var max int
		if len(top) > len(bottom) {
			max = len(bottom)
		} else {
			max = len(top)
		}

		for j := 0; j < max; j++ {
			if !slices.Equal(top[len(top)-1-j], bottom[j]) {
				continue Outer
			}
		}

		return i * 100

	}
	return 0
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
