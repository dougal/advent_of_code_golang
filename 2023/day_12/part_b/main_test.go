package main

import (
	"strings"
	"testing"
)

var cases = []struct {
	input    string
	expected int
}{
	{`???.### 1,1,3`, 1},
	{`.??..??...?##. 1,1,3`, 16384},
	{`?#?#?#?#?#?#?#? 1,3,1,6`, 1},
	{`????.#...#... 4,1,1`, 16},
	{`????.######..#####. 1,6,5`, 2500},
	{`?###???????? 3,2,1`, 506250},
	{`???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`, 525152},
}

func TestSumArrangements(t *testing.T) {
	for _, c := range cases {
		actual := sumArrangements(strings.NewReader(c.input))
		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}

// var validPrefixCases = []struct{
// 	input string
// 	groupings []int
// 	expected bool
// }{
// 	{"?", []int{ 1 }, true},
// 	{"#", []int{ 1 }, true},
// 	{".", []int{ 1 }, false},
// 	{"?.", []int{ 1 }, true},
// 	{"#.", []int{ 1 }, true},
// 	{"..", []int{ 1 }, false},

// 	{"???.###", []int{1,1,3}, true},

// 	{"#??.###", []int{1,1,3}, true},
// 	{"#.?.###", []int{1,1,3}, true},
// 	{"#.#.###", []int{1,1,3}, true},
// 	{"#...###", []int{1,1,3}, false},
// 	{"##?.###", []int{1,1,3}, false},

// 	{".??.###", []int{1,1,3}, true},
// 	{".#?.###", []int{1,1,3}, true},
// 	{".##.###", []int{1,1,3}, false},
// 	{".#..###", []int{1,1,3}, false},
// 	{"..?.###", []int{1,1,3}, true},
// 	{"..#.###", []int{1,1,3}, false},
// 	{"....###", []int{1,1,3}, false},
// }

// func TestPrefixSatisifies(t *testing.T) {
// 	for i, c := range validPrefixCases {
// 		l := Line{c.input, c.groupings}
// 		actual := l.validPrefix()
// 		if actual != c.expected {
// 			t.Errorf("Expected case %d to be %t but got %t\n", i, c.expected, actual)
// 		}
// 	}
// }


func BenchmarkSumArrangements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			sumArrangements(strings.NewReader(c.input))
		}
	}
}
