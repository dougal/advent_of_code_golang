package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(totalPriority(os.Stdin))
}
	
func totalPriority(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	total := 0

	for scanner.Scan() {
		total += rucksackPriority(scanner.Text())
	}

	return total
}

func rucksackPriority(s string) int {
	leftSide := map[byte]byte{}
	sackSize := len(s)
	halfSackSize := sackSize / 2

	for i := 0; i < halfSackSize; i++ {
		leftSide[s[i]] = 0b0
	}

	var currentChar byte;
	for i := halfSackSize; i < sackSize; i++ {
		currentChar = s[i]
		_, exists := leftSide[currentChar]
		if exists {
			break
		}
	}
	
	return itemPriority(currentChar)
}

func itemPriority(item byte) int {
	i := int(item)
	if i >= 97 { // a..z
		return i - 96 // make 1..26
	}

	// A..Z
	return i - 64 + 26 // make 27..52
}
