package main

import (
	"strings"
	"testing"
)

var cases = []struct {
	input    string
	expected int
}{
	{`L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`, 6},
}

func TestNumPassesZero(t *testing.T) {
	for _, c := range cases {
		actual := numPassesZero(strings.NewReader(c.input))
		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}

var divCases = []struct {
	s        int
	c        int
	expected int
}{
	{ 0, 0, 0},
	{ 0, -1, 0},
	{ 0, 1, 0},
	{ 0, 100, 1},
	{ 0, -100, 1},
	{ 0, 200, 2},
	{ 0, -200, 2},

	// From test example
	{ 50, -68, 1},
}

func TestIntDiv(t *testing.T) {
	for _, c := range divCases {
		actual := countPastZero(c.s, c.c)
		if actual != c.expected {
			t.Errorf("For %d to %d expected %d but got %d\n", c.s, c.c, c.expected, actual)
		}
	}
}
