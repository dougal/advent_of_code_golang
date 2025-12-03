package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sumJoltages(f))
}

func sumJoltages(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	t := 0

	for scanner.Scan() {
		line := scanner.Text()

		t += highestPair(line)
	}

	return t
}

func highestPair(l string) int {
	var curMax int

	for _, j := range l {
		one := curMax % 10
		ten := curMax - one
		newDigit := int(j) - 48

		if one > ten/10 {
			curMax = one*10 + newDigit
		} else if newDigit > one {
			curMax = ten + newDigit
		}
	}

	return curMax
}
