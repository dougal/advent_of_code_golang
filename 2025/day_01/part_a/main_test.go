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
L82`, 3},
}

func TestNumPositionZeros(t *testing.T) {
	for _, c := range cases {
		actual := numPositionZeros(strings.NewReader(c.input))
		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}
