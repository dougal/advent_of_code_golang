package main

import (
	"strings"
	"testing"
)

var cases = []struct {
	input    string
	expected int
}{
	{`.....
.S-7.
.|.|.
.L-J.
.....`, 4},
	{`..F7.
.FJ|.
SJ.L7
|F--J
LJ...`, 8},
}

func TestMaxDistance(t *testing.T) {
	for _, c := range cases {
		actual := maxDistance(strings.NewReader(c.input))
		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}
