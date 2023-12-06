package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(calculateProduct(f))
}

func calculateProduct(input io.Reader) int {
	is, err := io.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(is), "\n")

	var times []int
	for _, t := range strings.Split(lines[0], " ")[1:] {
		if t == "" {
			continue
		}

		ti, err := strconv.Atoi(t)
		if err != nil {
			log.Fatal(err)
		}

		times = append(times, ti)
	}

	var distances []int
	for _, t := range strings.Split(lines[1], " ")[1:] {
		if t == "" {
			continue
		}

		ti, err := strconv.Atoi(t)
		if err != nil {
			log.Fatal(err)
		}

		distances = append(distances, ti)
	}

	product := 1
	for i, t := range times {
		record := distances[i]
		waysBeaten := 0

		for buttonTime := 1; buttonTime < t; buttonTime++ {
			moveTime := t - buttonTime
			distMoved := buttonTime * moveTime
			// fmt.Printf("Moving: %d, Record: %d, this: %d\n", moveTime, record, distMoved)

			if distMoved > record {
				waysBeaten++
			}
		}
		product *= waysBeaten
	}

	return product
}
