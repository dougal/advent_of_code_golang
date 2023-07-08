package main

import (
	"bytes"
	"testing"
)

func TestCountOverlappingPairs(t *testing.T) {
	b := bytes.NewBufferString(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`)

	expected := 4

	result := countOverlappingPairs(b)

	if result != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, result)
	}
}

func TestIsOverlappingPair(t *testing.T) {
	examples := map[string]bool{
		"1-2,3-4":         false,
		"3-4,1-2":         false,
		"1-2,2-3":         true,
		"2-3,1-2":         true,
		"1-4,2-3":         true,
		"2-3,1-4":         true,
	}

	for s, expected := range examples {
		result := isOverlappingPair(s)

		if result != expected {
			t.Errorf("Expected %t, got %t instead.\n", expected, result)
		}
	}
}
