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
10 13 16 21 30 45`, 114},
}

var extrapolateCases = []struct {
	input    string
	expected int
}{
	{"0 3 6 9 12 15", 18},
	{"1 3 6 10 15 21", 28},
	{"10 13 16 21 30 45", 68},
}

func TestSumAll(t *testing.T) {
	for _, c := range sumCases {
		actual := sumAll(strings.NewReader(c.input))
		if actual != c.expected {
			fmt.Sprintf("expected %d but got %d", c.expected, actual)
		}
	}
}

func TestExtrapolate(t *testing.T) {
	for _, c := range extrapolateCases {
		actual := extrapolate(c.input)
		if actual != c.expected {
			fmt.Sprintf("expected %d but got %d", c.expected, actual)
		}
	}
}
