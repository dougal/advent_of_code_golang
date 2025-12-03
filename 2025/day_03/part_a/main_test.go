package main

import (
	"strings"
	"testing"
)

var cases = []struct {
	input    string
	expected int
}{
	{`987654321111111
811111111111119
234234234234278
818181911112111`, 357},
}

func TestSumJoltages(t *testing.T) {
	for _, c := range cases {
		actual := sumJoltages(strings.NewReader(c.input))
		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}

var highestPairCases = []struct {
	input    string
	expected int
}{
	{"12", 12},
	{"123", 23},
	{"987654321111111", 98},
	{"811111111111119", 89},
	{"234234234234278", 78},
	{"818181911112111", 92},
}

func TestHighestPair(t *testing.T) {
	for _, c := range highestPairCases {
		actual := highestPair(c.input)
		if actual != c.expected {
			t.Errorf("For %s expected %d but got %d\n", c.input, c.expected, actual)
		}
	}
}
