package main

import (
	"fmt"
	"strings"
	"testing"
)

var sumCases = []struct {
	input    string
	expected int
}{
	{`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`, 2},
}

var extrapolateCases = []struct {
	input    string
	expected int
}{
	{"0 3 6 9 12 15", -3},
	{"1 3 6 10 15 21", 0},
	{"10 13 16 21 30 45", 5},
}

func TestSumAll(t *testing.T) {
	for _, c := range sumCases {
		actual := sumAll(strings.NewReader(c.input))
		if actual != c.expected {
			fmt.Errorf("expected %d but got %d", c.expected, actual)
		}
	}
}

func TestExtrapolate(t *testing.T) {
	for _, c := range extrapolateCases {
		actual := extrapolate(c.input)
		if actual != c.expected {
			fmt.Errorf("expected %d but got %d", c.expected, actual)
		}
	}
}
