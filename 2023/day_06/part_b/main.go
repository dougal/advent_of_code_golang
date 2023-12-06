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

	fmt.Println(waysToBeat(f))
}

func waysToBeat(input io.Reader) int {
	is, err := io.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(is), "\n")

	_, timeDigits, _ := strings.Cut(lines[0], " ")
	ts := strings.ReplaceAll(timeDigits, " ", "")
	time, err := strconv.Atoi(ts)
	if err != nil {
		log.Fatal(err)
	}

	_, distDigits, _ := strings.Cut(lines[1], " ")
	ds := strings.ReplaceAll(distDigits, " ", "")
	record, err := strconv.Atoi(ds)
	if err != nil {
		log.Fatal(err)
	}

	waysBeaten := 0

	for buttonTime := 1; buttonTime < time; buttonTime++ {
		moveTime := time - buttonTime
		distMoved := buttonTime * moveTime
		// fmt.Printf("Moving: %d, Record: %d, this: %d\n", moveTime, record, distMoved)

		if distMoved > record {
			waysBeaten++
		}
	}

	return waysBeaten
}
