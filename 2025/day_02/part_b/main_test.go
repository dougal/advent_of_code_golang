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
		4174379265},
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
	{121212, false},
	{111111, false},
	{111, false},
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
	{"95-115", 210},
}

func TestRangeCountInvalid(t *testing.T) {
	for _, c := range rangeSumInvalidCases {
		actual := rangeSumInvalid(c.input)
		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}

var isRepeatCases = []struct {
	input    string
	splits   int
	expected bool
}{
	{"1", 1, true},
	{"1", 2, false},
	{"1212", 2, true},
	{"12121", 2, false},
	{"121212", 2, false},
	{"123123", 2, true},
	{"121212", 3, true},
}

func TestIsRepeatCases(t *testing.T) {
	for _, c := range isRepeatCases {
		actual := isRepeat(c.input, c.splits)
		if actual != c.expected {
			t.Errorf("For %s, %d expected %t but got %t\n", c.input, c.splits, c.expected, actual)
		}
	}
}
