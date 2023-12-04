package main

import (
	"bufio"
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

	fmt.Println(countAllCards(f))
}

func countAllCards(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	cardScores := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		cardScores = append(cardScores, cardMatches(line))
	}

	cc := 0

	for i := range cardScores {
		cc += countCards(cardScores, i)
	}

	return cc
}

func cardMatches(line string) int {
	m := 0
	winners := []int{}

	_, after, _ := strings.Cut(line, ": ")
	sWinners, sCandidates, _ := strings.Cut(after, " | ")

	for _, ws := range strings.Split(sWinners, " ") {
		if ws == "" {
			continue
		}

		w, err := strconv.Atoi(ws)
		if err != nil {
			log.Fatal(err)
		}

		winners = append(winners, w)
	}

	for _, cs := range strings.Split(sCandidates, " ") {
		if cs == "" {
			continue
		}

		c, err := strconv.Atoi(cs)
		if err != nil {
			log.Fatal(err)
		}

		if slices.Contains(winners, c) {
			m++
		}
	}

	return m
}

func countCards(cardScores []int, i int) int {
	cc := 1

	nCards := cardScores[i]
	sIndex := i + 1
	eIndex := i + 1 + nCards

	for ii := range cardScores[sIndex:eIndex] {
		cc += countCards(cardScores, ii+sIndex)
	}

	return cc
}
