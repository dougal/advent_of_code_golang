package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(countOverlappingPairs(os.Stdin))
}

func countOverlappingPairs(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	overlappingAssignments := 0

	for scanner.Scan() {
		if isOverlappingPair(scanner.Text()) {
			overlappingAssignments++
		}
	}

	return overlappingAssignments
}

func isOverlappingPair(s string) bool {
	elvesSections := strings.Split(s, ",")
	elfALimits := strings.Split(elvesSections[0], "-")
	elfBLimits := strings.Split(elvesSections[1], "-")

	a, _ := strconv.Atoi(elfALimits[0])
	b, _ := strconv.Atoi(elfALimits[1])

	c, _ := strconv.Atoi(elfBLimits[0])
	d, _ := strconv.Atoi(elfBLimits[1])

	return c >= a && c <= b || // left after right
		d >= a && d <= b || // right after left
		a <= c && b >= d || // left surrounds right
		c <= a && d >= b // right surrounds left
}
