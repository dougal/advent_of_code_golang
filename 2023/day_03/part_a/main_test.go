package main

import (
	"strings"
	"testing"
)

func Test(t *testing.T) {
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

  expected := 4361

  actual := partNumbersSum(strings.NewReader(input))

  if expected != actual {
    t.Fatalf("Expected %d but got %d", expected, actual)
  }
}

func TestParsePartNumberSum(t *testing.T) {
  pLine := `..592.....`
  cLine := `......755.`
  nLine := `...$.*....`

  expected := 755

  actual := parsePartNumberSum(pLine, cLine, nLine)

  if expected != actual {
    t.Fatalf("Expected %d but got %d", expected, actual)
  }
}

// func TestContainsymbol(t *testing.T) {
//   if containsSymbol("
// }
