package main

import (
	"strings"
	"testing"
)

func TestSumGearRatios(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	expected := 467835
	actual := sumGearRatios(strings.NewReader(input))

	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}

func TestFirstThreeLines(t *testing.T) {
	// Third line has a * replaced with a !
	input := `....................................18..........889.270.....748.280...997.................617..............622........763............476....
...529......434.....191..489...717...@.....................&....................939*7.....*....................606............760....*......
....*...473....!221................$........182......812........493.84....793..........794.......589..407..41...*.....................68....`

	// 939 × 7 + 617 × 794 + 476 × 68
	expected := 528839
	actual := sumGearRatios(strings.NewReader(input))

	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}

type Case struct {
	input    string
	expected int
}

var cases = []Case{
	{input: "...\n...\n...", expected: 0},
	{input: "1..\n1..\n*..", expected: 0},
	{input: "1..\n.*.\n..2", expected: 2},
}

func TestSumGearRatioCases(t *testing.T) {
	for i, c := range cases {
		actual := sumGearRatios(strings.NewReader(c.input))
		if c.expected != actual {
			t.Fatalf("Expected %d but got %d for case #%d\n", c.expected, actual, i)
		}
	}
}
