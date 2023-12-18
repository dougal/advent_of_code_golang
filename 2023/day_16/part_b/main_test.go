package main

import (
	"strings"
	"testing"
)

var MaxEnergizedTilesCases = []struct {
	input    string
	expected int
}{
  {`...`, 3},
  {`.-.`, 3},
  {`./.`, 2},
  {`.\.`, 2},
  {`.|.`, 2},

  {`..\
../`, 6},

  {`..\
|./
...`, 7},

  {`.\.
.-.`, 5},

	{`.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`, 51},
}

func TestSomething(t *testing.T) {
	for _, c := range MaxEnergizedTilesCases {
		actual := MaxEnergizedTiles(strings.NewReader(c.input))
		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}
