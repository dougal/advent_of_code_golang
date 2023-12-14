package main

import (
	"strings"
	"testing"
)

var cases = []struct {
	input    string
	expected int
}{
	{`O
O
O`, 6},
	{`O
.`, 2},
	{`.
O`, 2},
	{`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`, 136},
}

func TestTotalLoad(t *testing.T) {
	for _, c := range cases {
		actual := totalLoad(strings.NewReader(c.input))
		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}
