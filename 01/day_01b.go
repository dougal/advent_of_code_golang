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
	fmt.Println(most_calories_by_three_elves(os.Stdin)) // 197,400
}

func most_calories_by_three_elves(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	maxCalories0 := 0
	maxCalories1 := 0
	maxCalories2 := 0
	currentCalorieCount := 0

	for scanner.Scan() {
		c, err := strconv.Atoi(scanner.Text())
		currentCalorieCount += c

		if err != nil { // Empty line
			if currentCalorieCount > maxCalories0 {
				maxCalories0 = currentCalorieCount
			}
			// Low budget bubble sort
			if maxCalories0 > maxCalories1 {
				maxCalories0, maxCalories1 = maxCalories1, maxCalories0
			}
			// And again.
			if maxCalories1 > maxCalories2 {
				maxCalories1, maxCalories2 = maxCalories2, maxCalories1
			}
			currentCalorieCount = 0
		}
	}

	return maxCalories0 + maxCalories1 + maxCalories2
}
