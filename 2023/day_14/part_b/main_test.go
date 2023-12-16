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

var tiltWestCases = []struct {
	in  string
	exp string
}{
	{`O`, `O`},

	{`O.`, `O.`},

	{`.O`, `O.`},

	{`..O`, `O..`},

	{`#.O`, `#O.`},

	{`#.O
.#..O#.#`,
		`#O.
.#O..#.#`},
}

func TestTiltWest(t *testing.T) {
	for caseNumber, c := range tiltWestCases {
		field := parseField(strings.NewReader(c.in))
		expected := parseField(strings.NewReader(c.exp))

		field.TiltWest()

		for i, row := range field {
			if !slices.Equal(row, expected[i]) {
				t.Errorf("Case #%d Expected:\n%s\nbut got:\n%s\n", caseNumber, expected.String(), field.String())
				break
			}
		}
	}
}

var tiltEastCases = []struct {
	in  string
	exp string
}{
	{`O`, `O`},

	{`O.`, `.O`},

	{`.O`, `.O`},

	{`.O.`, `..O`},

	{`O.#`, `.O#`},

	{`#O.
.#O..#.#`,
		`#.O
.#..O#.#`},
}

func TestTiltEast(t *testing.T) {
	for caseNumber, c := range tiltEastCases {
		field := parseField(strings.NewReader(c.in))
		expected := parseField(strings.NewReader(c.exp))

		field.TiltEast()

		for i, row := range field {
			if !slices.Equal(row, expected[i]) {
				t.Errorf("Case #%d Expected:\n%s\nbut got:\n%s\n", caseNumber, expected.String(), field.String())
				break
			}
		}
	}
}

var tiltNorthCases = []struct {
	in  string
	exp string
}{
	{`O`, `O`},

	{`.
O`, `O
.`},

	{`O
.`,
`O
.`},

	{`.
O
.`, `O
.
.`},

	{`#
.
O`, `#
O
.`},

	{`.#
#.
.O
..
O.
#.
..
#.`,
		`.#
#O
O.
..
..
#.
..
#.`},
}

func TestTiltNorth(t *testing.T) {
	for caseNumber, c := range tiltNorthCases {
		field := parseField(strings.NewReader(c.in))
		expected := parseField(strings.NewReader(c.exp))

		field.TiltNorth()

		for i, row := range field {
			if !slices.Equal(row, expected[i]) {
				t.Errorf("Case #%d Expected:\n%s\nbut got:\n%s\n", caseNumber, expected.String(), field.String())
				break
			}
		}
	}
}
