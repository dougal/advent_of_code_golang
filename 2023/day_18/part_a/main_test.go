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
  {`R 1 (#FF0000)
D 1 (#00FF00)
L 1 (#0000FF)
U 1 (#F0000F)`, 4},

  // 9x9 solid, starts in middle
	{`U 1 (#FF0000)
R 1 (#FF0000)
D 2 (#FF0000)
L 2 (#FF0000)
U 2 (#FF0000)`, 9},

  // 9x9 solid, starts going left
{`L 1 (#FF0000)
U 1 (#FF0000)
R 2 (#FF0000)
D 2 (#FF0000)
L 2 (#FF0000)`, 9},

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
U 2 (#7a21e3)`, 62},
}

func TestSomething(t *testing.T) {
	for _, c := range LavaVolumeCases {
		actual := LavaVolume(strings.NewReader(c.input))
		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}
