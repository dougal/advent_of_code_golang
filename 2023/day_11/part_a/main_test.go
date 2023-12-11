package main

import (
	"strings"
	"testing"
)

var cases = []struct {
	input    string
	expected int
}{
	{`#.
.#`, 2},
	{`#..
..#`, 4},
	{`#.
..
.#`, 4},
	{`...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`, 374},
}

func TestSumGalaxyDistances(t *testing.T) {
	for _, c := range cases {
		actual := sumGalaxyDistances(strings.NewReader(c.input))
		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}
