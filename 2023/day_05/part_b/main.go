package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lowestLocation(f))
}

func lowestLocation(input io.Reader) int {
	seedLimits, mappings := parse(input)
	lowestLocation := math.MaxInt

	// For each pair of seed limits
	// Calculate mapped number through each mapping until get to the final one
	for i := 0; i < len(seedLimits); i += 2 {
		for seed := seedLimits[i]; seed < seedLimits[i]+seedLimits[i+1]; seed++ {
			number := seed

			for _, mapping := range mappings {
				number = mapping.mapNumber(number)
			}

			if number < lowestLocation {
				lowestLocation = number
			}
		}
	}

	return lowestLocation
}

type mapping struct {
	name   string
	ranges []rng
}

func (m mapping) mapNumber(n int) int {
	for _, r := range m.ranges {
		newNumber, err := r.mapNumber(n)
		if err == nil {
			return newNumber
		}
	}

	return n
}

type rng struct {
	destStart   int
	sourceStart int
	length      int
}

func newRng(destStart int, sourceStart int, length int) rng {
	return rng{
		destStart: destStart,
		sourceStart: sourceStart,
		length: length,
	}
}

func (r rng) mapNumber(n int) (int, error) {
	if n >= r.sourceStart && n < r.sourceStart+r.length {
		return n - r.sourceStart + r.destStart, nil
	}

	return -1, errors.New("Not within range")
}

func parse(input io.Reader) ([]int, []mapping) {
	sInput, err := io.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(string(sInput), "\n\n")
	sParts := strings.Split(parts[0], " ")
	seeds := []int{}

	for _, ss := range sParts[1:] {
		is, err := strconv.Atoi(ss)
		if err != nil {
			log.Fatal(err)
		}

		seeds = append(seeds, is)
	}

	mappings := []mapping{}
	for _, ms := range parts[1:] {
		lines := strings.Split(ms, "\n")
		m := mapping{}
		m.name = strings.Split(lines[0], " ")[0]

		for _, l := range lines[1:] {
			if l == "" {
				continue
			}
			vs := strings.Split(l, " ")
			ds, err := strconv.Atoi(vs[0])
			if err != nil {
				log.Fatal(err)
			}

			ss, err := strconv.Atoi(vs[1])
			if err != nil {
				log.Fatal(err)
			}

			lth, err := strconv.Atoi(vs[2])
			if err != nil {
				log.Fatal(err)
			}

			m.ranges = append(m.ranges, newRng(ds, ss, lth))
		}

		mappings = append(mappings, m)
	}

	return seeds, mappings
}
