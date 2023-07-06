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

type TopThreeElves []int64

// Take a calorie count, add to array if larger than any of the existing value.
func (t TopThreeElves) CompareAndAddElf(c int64) {
	for i, v := range t {
		if c > v {
			t[i] = c
			break
		}
	}
}

func InitTopThreeElves() TopThreeElves {
	return make([]int64, 3)
}

func (t TopThreeElves) Sum() int64 {
	return t[0] + t[1] + t[2]
}

func most_calories_by_three_elves(r io.Reader) int64 {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var topThreeCalorieCounts = InitTopThreeElves()
	var currentCalorieCount int64 = 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			topThreeCalorieCounts.CompareAndAddElf(currentCalorieCount)
			fmt.Println(topThreeCalorieCounts)
			currentCalorieCount = 0
			continue
		}

		if c, err := strconv.ParseInt(line, 10, 64); err == nil {
			currentCalorieCount += c
		}
	}

	topThreeCalorieCounts.CompareAndAddElf(currentCalorieCount)

	return topThreeCalorieCounts.Sum()
}
