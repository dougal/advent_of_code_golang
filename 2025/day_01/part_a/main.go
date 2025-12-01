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

		change := int(c)
		if dir == 'L' {
			change *= -1
		} 

		curPos = (curPos + change) % 100

		if curPos == 0 {
			count += 1
		}
	}

	return count
}
