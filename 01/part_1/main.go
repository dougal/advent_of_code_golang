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
	fmt.Println(most_calories_by_one_elf(os.Stdin))
}

func most_calories_by_one_elf(r io.Reader) int64 {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var highestCalorieCount int64 = 0
	var currentCalorieCount int64 = 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if currentCalorieCount > highestCalorieCount {
				highestCalorieCount = currentCalorieCount
			}
			currentCalorieCount = 0
			continue // Is this right?
		}

		if c, err := strconv.ParseInt(line, 10, 64); err == nil {
			currentCalorieCount += c
		}
	}

	if currentCalorieCount > highestCalorieCount {
		highestCalorieCount = currentCalorieCount
	}

	return highestCalorieCount
}
