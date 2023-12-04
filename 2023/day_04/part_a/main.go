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

	fmt.Println(sumScores(f))
}

func sumScores(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	s := 0

	for scanner.Scan() {
		line := scanner.Text()

		s += cardScore(line)
	}

	return s
}

func cardScore(line string) int {
	s := 0
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
			if s == 0 {
				s = 1
			} else {
				s *= 2
			}
		}
	}

	return s
}
