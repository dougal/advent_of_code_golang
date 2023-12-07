package main

import (
	"bufio"
	"cmp"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(AllHandsScore(f))
}

var cardRank = map[string]string{
	"A": "m",
	"K": "l",
	"Q": "k",
	"T": "j",
	"9": "i",
	"8": "h",
	"7": "g",
	"6": "f",
	"5": "e",
	"4": "d",
	"3": "c",
	"2": "b",
	"J": "a",
}

type Hand struct {
	cards []string
	bid   int
}

func (h Hand) Print() string {
	return fmt.Sprintf("{%v % 3d}", h.cards, h.bid)
}

func NewHand(line string) Hand {
	cards, bidS, _ := strings.Cut(line, " ")

	h := Hand{}

	for _, r := range strings.Split(cards, "") {
		h.cards = append(h.cards, r)
	}

	b, err := strconv.Atoi(bidS)
	if err != nil {
		log.Fatal(err)
	}

	h.bid = b

	return h
}

func (h Hand) RankingScore() string {
	perms := map[string]int{}
	rankHand := []string{}

	for _, c := range h.cards {
		rankHand = append(rankHand, cardRank[c])
		if _, ok := perms[c]; !ok {
			perms[c] = 0
		}

		perms[c]++
	}

	jCount, ok := perms["J"]
	if !ok {
		jCount = 0
	}

	var counts []int
	for k, v := range perms {
    if k == "J" {
			continue
		}
		counts = append(counts, v)
	}

	slices.Sort(counts)
	slices.Reverse(counts)

	// Cover case where only Jokers
	if len(counts) == 0 {
		counts = append(counts, 5)
	} else {
		// Add the count of Jokers to the highest ranking group.
		counts[0] += jCount
	}

	// Default is High Card
	typeScore := 1
	switch {
	// five of a kind
	case counts[0] == 5:
		typeScore = 7
	// four of a kind
	case counts[0] == 4:
		typeScore = 6
	// full house
	case counts[0] == 3 && counts[1] == 2:
		typeScore = 5
	// three of a kind
	case counts[0] == 3:
		typeScore = 4
	// two pair
	case counts[0] == 2 && counts[1] == 2:
		typeScore = 3
	// one pair
	case counts[0] == 2:
		typeScore = 2
	}

	rankHand = append([]string{strconv.Itoa(typeScore)}, rankHand...)

	return strings.Join(rankHand, "")
}

func AllHandsScore(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var (
		s     = 0
		hands []Hand
	)

	for scanner.Scan() {
		hands = append(hands, NewHand(scanner.Text()))
	}

	slices.SortStableFunc(hands, func(a, b Hand) int {
		return cmp.Compare(a.RankingScore(), b.RankingScore())
	})

	for i, h := range hands {
		// fmt.Printf("%s - %s - %d\n", h.Print(), h.RankingScore(), i + 1)
		s += h.bid * (i + 1)
	}

	return s
}
