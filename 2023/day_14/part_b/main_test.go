package main

import (
	"slices"
	"strings"
	"testing"
)

var totalLoadCases = []struct {
	input    string
	expected int
}{
	{`O`, 1},
	// {`O
	// O
	// O`, 6},
	// {`O
	// .`, 2},
	// {`.
	// O`, 2},
	{`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`, 64},
}

func TestTotalLoad(t *testing.T) {
	for _, c := range totalLoadCases {
		actual := totalLoad(strings.NewReader(c.input))
		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}

var rotateClockwiseCases = []struct {
	in  string
	exp string
}{
	{`O`, `O`},

	{`O.`, `O
.`},

	{`O
.`, `.O`},

	{`AB
CD`, `CA
DB`},

	{`ABC`, `A
B
C`},

	{`A
B
C`, `CBA`},
}

func TestRotateClockwise(t *testing.T) {
	for ci, c := range rotateClockwiseCases {
		input := parseField(strings.NewReader(c.in))
		expected := parseField(strings.NewReader(c.exp))

		actual := rotateClockwise(input)

		for i, row := range actual {
			if !slices.Equal(row, expected[i]) {
				t.Errorf("Case #%d Expected:\n%s\nbut got:\n%s\n", ci, FieldToString(expected), FieldToString(actual))
				break
			}
		}
	}
}
