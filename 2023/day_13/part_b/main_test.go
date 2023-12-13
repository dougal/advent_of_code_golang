package main

import (
	"strings"
	"testing"
)

var sumReflectionCases = []struct {
	input    string
	expected int
}{
	// {`AB
	// AB`, 100},

	// {`AA
	// BB`, 1},

	// {`AB
	// AB
	// AA`, 100},

	// {`AA
	// AB
	// AB`, 200},

	// {`##.
	// ...`, 1},

	// Top Left
	{`#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.`, 300},

	// First symbol on second row.
	{`#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`, 100},

	{`#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`, 400},
}

func TestSumReflections(t *testing.T) {
	for _, c := range sumReflectionCases {
		actual := sumReflections(strings.NewReader(c.input))
		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}
