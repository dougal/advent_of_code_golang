package main

import (
	"strings"
	"testing"
)

const example = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func TestPossibleLineId(t *testing.T) {
  impossible := "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"
  impossibleId := possibleLineId(impossible)
  if impossibleId != 0 {
    t.Fatalf("Expected 0 got %d", impossibleId)
  }

  possible := "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
  possibleId := possibleLineId(possible)
  if possibleId != 5 {
    t.Fatalf("Expected 5 got %d", possibleId)
  }
}

func TestSum(t *testing.T) {
  s := sum(strings.NewReader(example))
  if s != 8 {
    t.Fatalf("Expected 8 got %d", s)
  }
}
