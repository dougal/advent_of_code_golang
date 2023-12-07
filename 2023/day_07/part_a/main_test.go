package main

import (
	"strings"
	"testing"
)

const testInput = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

// From https://www.reddit.com/r/adventofcode/comments/18cr4xr/2023_day_7_better_example_input_not_a_spoiler/
const moreTestInput = `2345A 1
Q2KJJ 13
Q2Q2Q 19
T3T3J 17
T3Q33 11
2345J 3
J345A 2
32T3K 5
T55J5 29
KK677 7
KTJJT 34
QQQJA 31
JJJJJ 37
JAAAA 43
AAAAJ 59
AAAAA 61
2AAAA 23
2JJJJ 53
JJJJ2 41`

var rankCases = []struct {
	cards    string
	expected string
}{
	{"AAAAA", "7mmmmm"},
	{"AAAAK", "6mmmml"},
	{"AAAKK", "5mmmll"},
	{"AAAQK", "4mmmkl"},
	{"AAKKQ", "3mmllk"},
	{"AAKQJ", "2mmlkj"},
	{"AKQJT", "1mlkji"},

	// Bonus tests based on failures
	{"T55J5", "4iddjd"},
	{"T3Q33", "4ibkbb"},
}

func TestHandsScore(t *testing.T) {
	exp := 6440
	actual := AllHandsScore(strings.NewReader(testInput))

	if exp != actual {
		t.Errorf("Expected %d but got %d", exp, actual)
	}
}

func TestMoreHandsScore(t *testing.T) {
	expMore := 6592
	actualMore := AllHandsScore(strings.NewReader(moreTestInput))

	if expMore != actualMore {
		t.Errorf("Expected %d but got %d", expMore, actualMore)
	}
}

func TestRankingScore(t *testing.T) {
	for _, c := range rankCases {
		h := NewHand(c.cards + " 123")
		act := h.RankingScore()

		if act != c.expected {
			t.Errorf("Expected %s but got %s\n", c.expected, act)
		}
	}
}
