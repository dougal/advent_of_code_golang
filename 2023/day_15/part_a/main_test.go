package main

import (
	"strings"
	"testing"
)

var SumHashescases = []struct {
	input    string
	expected int
}{
  {"HASH", 52},
	{"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7", 1320},
}

func TestSumHashes(t *testing.T) {
	for _, c := range SumHashescases {
		actual := SumHashes(strings.NewReader(c.input))
		if actual != c.expected {
			t.Errorf("Expected %d but got %d\n", c.expected, actual)
		}
	}
}
