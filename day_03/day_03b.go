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
	group := []string{}

	for scanner.Scan() {
		group = append(group, scanner.Text())

		if len(group) == 3 {
			total += groupPriority(group)
			group = []string{}
		}
	}

	return total
}

func groupPriority(g []string) int {
	itemCounts := map[rune]int{}
	rucksack := map[rune]bool{}

	for _, r := range g {
		rucksack = map[rune]bool{} // reset rucksack
		for _, i := range r {
			if _, exists := rucksack[i]; exists {
				// Skip to next as this is duplicate within rucksack.
				continue
			}
			if _, exists := itemCounts[i]; !exists {
				itemCounts[i] = 0
			}

			itemCounts[i] += 1
			rucksack[i] = true
		}
	}

	var commonItem rune
	for item, count := range itemCounts {
		if count == 3 {
			commonItem = item
		}
	}

	return itemPriority(commonItem)
}

func itemPriority(item rune) int {
	i := int(item)
	if i >= 97 { // a..z
		return i - 96 // make 1..26
	}

	// A..Z
	return i - 64 + 26 // make 27..52
}
