package main

import (
	"strings"
	"testing"
)

var testInput = `Time:      7  15   30
Distance:  9  40  200`

func TestCalculateProduct(t *testing.T) {
	exp := 71503
	act := waysToBeat(strings.NewReader(testInput))

	if exp != act {
		t.Fatalf("Expected %d but got %d\n", exp, act)
	}
}
