package main

import (
	"strings"
	"testing"
)

var cases = []struct {
	input    string
	expected int
}{
	{`LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`, 6},
}

func TestCountSteps(t *testing.T) {
	for _, c := range cases {
		actual := countSteps(strings.NewReader(c.input))

		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}
