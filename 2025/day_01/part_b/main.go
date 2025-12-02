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

	fmt.Println(numPassesZero(f))
}

func numPassesZero(input io.Reader) int {
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

		change := int(c)
		if dir == 'L' {
			change *= -1
		}

		origPos := curPos
		curPos = (curPos + change) % 100
		if curPos < 0 {
			curPos = 100 + curPos
		}
		count += countPastZero(origPos, change)

		// fmt.Printf("%c, %v, %v, %v\n", dir, c, count, curPos)
	}

	return count
}

func countPastZero(s, c int) int {
	// Maths? No, don't be silly. Let's just do iteration.
	var n int

	f := s + c

	if c == 0 {
		return 0
	} else if c > 0 {
		for i := s+1; i <= f; i++ {
			if i % 100 == 0 {
				n++
			}
		}
	} else {
		for i := s-1; i >= f; i-- {
			if i % 100 == 0 {
				n++
			}
		}
	}

	return n
}
