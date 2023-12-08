package main

import (
	"strings"
	"testing"
)

var cases = []struct{input string; expected int}{
	{`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`, 2},
  {`LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`, 6},
}

func TestCountSteps(t *testing.T) {
	for _, c := range cases {
		actual := countSteps(strings.NewReader(c.input))

		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}
