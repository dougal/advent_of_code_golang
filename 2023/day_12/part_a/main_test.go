package main

import (
	"strings"
	"testing"
)

var cases = []struct {
	input    string
	expected int
}{
	{"? 1", 1},
	{"?? 1", 2},
	{`???.### 1,1,3`, 1},
	{`.??..??...?##. 1,1,3`, 4},
	{`?#?#?#?#?#?#?#? 1,3,1,6`, 1},
	{`????.#...#... 4,1,1`, 1},
	{`????.######..#####. 1,6,5`, 4},
	{`?###???????? 3,2,1`, 10},
	{`???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`, 21},
}

func TestSumArrangements(t *testing.T) {
	for _, c := range cases {
		actual := sumArrangements(strings.NewReader(c.input))
		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}

func BenchmarkSumArrangements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			sumArrangements(strings.NewReader(c.input))
		}
	}
}
