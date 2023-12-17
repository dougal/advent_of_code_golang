package main

import (
	"strings"
	"testing"
)

var CountEnergizedTilesCases = []struct {
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
..//.|....`, 46},
}

func TestSomething(t *testing.T) {
	for _, c := range CountEnergizedTilesCases {
		actual := CountEnergizedTiles(strings.NewReader(c.input))
		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}
