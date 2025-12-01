package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(numPositionZeros(f))
}

func numPositionZeros(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	curPos := 50
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0]
		c, err := strconv.ParseInt(line[1:], 10, 0)
		if err != nil {
			log.Fatal(err)
		}

		var change int
		if dir == 'L' {
			change = (100 - (int(c) % 100))
		} else {
			change = int(c) % 100
		}

		curPos = (curPos + change) % 100

		if curPos == 0 {
			count += 1
		}
	}

	return count
}
