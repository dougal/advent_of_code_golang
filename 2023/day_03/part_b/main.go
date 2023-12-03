package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sumGearRatios(f))
}

func sumGearRatios(input io.Reader) int {
	// Each line
	// Look for gears
	// Look for adjacent numbers on the previous, current, and next lines
	// If there are two, multiply these together to get the ratio
	// Ignore more than two?
	// Sum every ratio
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	s := 0
	var (
		pLine Line
		cLine Line
		nLine Line
	)

	for scanner.Scan() {
		pLine = cLine
		cLine = nLine
		nLine = NewLine(scanner.Text())

		s += sumRatiosForLine(pLine, cLine, nLine)
	}

	// Handle the last line
	s += sumRatiosForLine(cLine, nLine, Line{})

	return s
}

func sumRatiosForLine(pLine Line, cLine Line, nLine Line) int {
	// For every gear in current line
	// Check to see if adjacent numbers on any lines
	s := 0

	for _, g := range cLine.gearIndexes {
		nums := []int{}

		vp := pLine.numAdjTo(g)
		nums = append(nums, vp...)

		vc := cLine.numAdjTo(g)
		nums = append(nums, vc...)

		vn := nLine.numAdjTo(g)
		nums = append(nums, vn...)

		if len(nums) == 2 {
			s += nums[0] * nums[1]
		}
	}

	return s
}

var numRe = regexp.MustCompile(`\d+`)
var gearRe = regexp.MustCompile(`\*`)

type Line struct {
	length      int
	numIndexes  [][]int
	nums        []int
	gearIndexes []int
}

func NewLine(s string) Line {
	l := Line{}
	l.length = len(s)
	l.numIndexes = numRe.FindAllStringIndex(s, -1)

	for _, i := range l.numIndexes {
		v, err := strconv.Atoi(s[i[0]:i[1]])
		if err != nil {
			log.Fatal(err)
		}
		l.nums = append(l.nums, v)
	}

	for _, i := range gearRe.FindAllStringIndex(s, -1) {
		l.gearIndexes = append(l.gearIndexes, i[0])
	}

	return l
}

func (l Line) numAdjTo(g int) []int {
	gStart := g - 1
	gEnd := g + 1
	nums := []int{}

	for i, n := range l.numIndexes {
		// Check if there is overlap +/- 1 of g with the num
		if (n[0] >= gStart && n[0] <= gEnd) ||
			((n[1]-1) >= gStart && (n[1]-1) <= gEnd) {
			nums = append(nums, l.nums[i])
		}
	}

	return nums
}
