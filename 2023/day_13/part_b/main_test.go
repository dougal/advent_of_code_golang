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
#.#.##.#.`, 100},

// First symbol on second row.
	{`#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`, 300},

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

var equalIfOneFlippedCases = []struct {
	r1    string
	r2    string
	expected bool
} {
		{"AAA", "AAA", true},
		{"BAA", "AAA", true},
		{"BBA", "AAA", false},
		{"BBB", "AAA", false},
}

func TestEqualIfOneFlipped(t *testing.T) {
	for _, c := range equalIfOneFlippedCases {
		actual := equalIfOneFlipped([]rune(c.r1), []rune(c.r2))
		if actual != c.expected {
			t.Errorf("Expected %t but got %t\n", c.expected, actual)
		}
	}
}
