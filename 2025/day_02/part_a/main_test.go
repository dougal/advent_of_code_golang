package main

import (
	"strings"
	"testing"
)

var cases = []struct {
	input    string
	expected int
}{
	{
		`11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`,
		1227775554},
}

func TestSomething(t *testing.T) {
	for _, c := range cases {
		actual := sumInvalidIds(strings.NewReader(c.input))
		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}

var validIdcases = []struct {
	input    int
	expected bool
}{
	{10, true},
	{11, false},
	{22, false},
	{123, true},
	{1212, false},
	{123123, false},
}

func TestValidId(t *testing.T) {
	for _, c := range validIdcases {
		actual := validId(c.input)
		if actual != c.expected {
			t.Errorf("Expected %t but got %t\n", c.expected, actual)
		}
	}
}

var rangeSumInvalidCases = []struct {
	input    string
	expected int
}{
	{"11-22", 33},
	{"95-115", 99},
}

func TestRangeCountInvalid(t *testing.T) {
	for _, c := range rangeSumInvalidCases {
		actual := rangeSumInvalid(c.input)
		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}
