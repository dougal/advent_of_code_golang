package main

import (
	"strings"
	"testing"
)

var LavaVolumeCases = []struct {
	input    string
	expected int
}{
	// Solid square
	{`R 1 (#000010)
D 1 (#000011)
L 1 (#000012)
U 1 (#000013)`, 4},

	// Provided example
	{`R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`, 952408144115},
}

func TestSomething(t *testing.T) {
	for _, c := range LavaVolumeCases {
		actual := LavaVolume(strings.NewReader(c.input))
		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}
