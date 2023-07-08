package main

import (
	"bytes"
	"testing"
)

func TestTotalPriority( t *testing.T) {
  b := bytes.NewBufferString(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`)

  expected := 157

	result := totalPriority(b)

	if result != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, result)
	}
}

func TestRucksackPriority( t *testing.T) {
  examples := map[string]int{
		"vJrwpWtwJgWrhcsFMMfFFhFp": 16,
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL": 38,
		"PmmdzqPrVvPwwTWBwg": 42,
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn": 22,
		"ttgJtRGJQctTZtZT": 20,
		"CrZsJsPPZsGzwwsLwLmpwMDw": 19,
	}

	for s, expected := range examples {
		result := rucksackPriority(s)

		if result != expected {
			t.Errorf("Expected %d, got %d instead.\n", expected, result)
		}
	}
}

func TestItemPriority( t *testing.T) {
  examples := map[rune]int{
		'a': 1,
		'b': 2,
		'c': 3,
		'x': 24,
		'y': 25,
		'z': 26,
		'A': 27,
		'B': 28,
		'C': 29,
		'X': 50,
		'Y': 51,
		'Z': 52,
	}

	for s, expected := range examples {
		result := itemPriority(byte(s))

		if result != expected {
			t.Errorf("For '%c', expected %d, got %d instead.\n", s, expected, result)
		}
	}
}
