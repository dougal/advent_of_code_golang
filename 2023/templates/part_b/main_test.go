package main

import (
	"strings"
	"testing"
)

var cases = []struct {
	input    string
	expected int
}{
	// {``, 999},
}

func TestSomething(t *testing.T) {
	for _, c := range cases {
		actual := doSomething(strings.NewReader(c.input))
		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}
