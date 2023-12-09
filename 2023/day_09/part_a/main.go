package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sumAll(f))
}

func sumAll(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	s := 0

	for scanner.Scan() {
		s += extrapolate(scanner.Text())
	}

	return s
}

func extrapolate(line string) int {
	var rows [][]int
	rows = append(rows, []int{})

	for _, s := range strings.Split(line, " ") {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		rows[0] = append(rows[0], n)
	}

	// differentiate until there is a row of all-zeroes
	for !isAllZeros(rows[len(rows)-1]) {
		rows = append(rows, differentiate(rows[len(rows)-1]))
	}

	slices.Reverse(rows)

	lastNum := 0
	for _, r := range rows[1:] {
		lastNum = r[len(r)-1] + lastNum
	}

	return lastNum
}

func differentiate(rowA []int) []int {
	var rowB []int

	for i, n := range rowA[1:] {
		prev := rowA[i]
		diff := n - prev
		rowB = append(rowB, diff)
	}

	return rowB
}

func isAllZeros(row []int) bool {
	for _, n := range row {
		if n != 0 {
			return false
		}
	}
	return true
}
