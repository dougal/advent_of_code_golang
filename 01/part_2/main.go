package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Usage:
func main() {
	fmt.Println(most_calories_by_three_elves(os.Stdin))
}

func most_calories_by_three_elves(r io.Reader) int64 {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var topThreeCalorieCounts = make([]int64, 3)
	var currentCalorieCount int64 = 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			for i, v := range topThreeCalorieCounts {
				if currentCalorieCount > v {
					topThreeCalorieCounts[i] = currentCalorieCount
					break
				}
			}
			fmt.Println(topThreeCalorieCounts)
			currentCalorieCount = 0
			continue
		}

		if c, err := strconv.ParseInt(line, 10, 64); err == nil {
			currentCalorieCount += c
		}
	}

	for i, v := range topThreeCalorieCounts {
		if currentCalorieCount > v {
			topThreeCalorieCounts[i] = currentCalorieCount
			break
		}
	}

	return topThreeCalorieCounts[0] + topThreeCalorieCounts[1] + topThreeCalorieCounts[2]
}
