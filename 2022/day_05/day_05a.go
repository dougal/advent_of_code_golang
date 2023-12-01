package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println(topCrates(os.Stdin))
}

func topCrates(input io.Reader) string {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	stacks := parseStartState(scanner)

	// Ignore blank line
	scanner.Scan()

	// Parse integers out of instructions
	// Follow instructions
	re := regexp.MustCompile(`\d+`)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		nums := re.FindAllStringSubmatch(scanner.Text(), -1)
		fromStackIndex, _ := strconv.Atoi(nums[0][0])
		fromStackIndex--
		toStackIndex, _ := strconv.Atoi(nums[1][0])
		toStackIndex--
		crateCount, _ := strconv.Atoi(nums[2][0])

		// slicePoint := len(stacks[fromStackIndex]) - crateCount
		// stacks[toStackIndex] = append(stacks[toStackIndex], stacks[fromStackIndex][slicePoint:]...)
		// stacks[fromStackIndex] = stacks[fromStackIndex][:slicePoint]
		for i:=0; i<crateCount; i++ {
			fromStackLen := len(stacks[fromStackIndex])
			stacks[toStackIndex] = append(stacks[toStackIndex], stacks[fromStackIndex][fromStackLen - 1])
			stacks[fromStackIndex] = stacks[fromStackIndex][:fromStackLen-1]
		}
	}

	topmostCrates := []rune{}
	// Concatenate the final crates on each stack together
	for _, stack := range stacks {
		if len(stack) == 0 {
			continue
		}

		topmostCrates = append(topmostCrates, stack[len(stack) - 1])
	}
	return string(topmostCrates)
}

func parseStartState(scanner *bufio.Scanner) [][]rune {
	// Loop through stack rows.
	// Loop through crate identifier positions
	// Number prepend each crate to appropriate stack
	// When positions are numeric, have reached bottom.
	var (
		stackCount int
		stacks     [][]rune
	)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		// Stacks have trailing whitespace so line width is fixed for all lines of
		// stack initial positions.
		if stackCount == 0 {
			stackCount = (len(line) + 1) / 4
			stacks = make([][]rune, stackCount)
		}

		// Row like ' 1   2   3 '
		// No more rows.
		if line[1] == '1' {
			break
		}

		for stackIndex := 0; stackIndex < stackCount; stackIndex++ {
			char := rune(line[stackIndex*4+1])

			stacks[stackIndex] = append([]rune{char}, stacks[stackIndex]...)
		}
	}

	return stacks
}
