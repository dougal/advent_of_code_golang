package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
  fmt.Println(summation())
}

var wordNumbers = map[string]int{
  "one":   1,
  "two":   2,
  "three": 3,
  "four":  4,
  "five":  5,
  "six":   6,
  "seven": 7,
  "eight": 8,
  "nine":  9,
  "1":     1,
  "2":     2,
  "3":     3,
  "4":     4,
  "5":     5,
  "6":     6,
  "7":     7,
  "8":     8,
  "9":     9,
}

// Process each line searching for all the number.
// Word-numbers may overlap, example: nineight is 9 and 8.
// Take the first and last digit.
// String concat these digits to form a two digit number.
// Sum the numbers from every line.
func summation() int {
  f, err := os.Open("input_a.txt")
  if err != nil {
    log.Fatal(err)
  }

  sum := 0
  scanner := bufio.NewScanner(f)
  scanner.Split(bufio.ScanLines)

  for scanner.Scan() {
    line := scanner.Text()
    sum += lineToCalibaration(line)
  }

  return sum
}

func lineToCalibaration(line string) int {
  digits := []int{}

  for i := 0; i < len(line); i++ {
    haystack := line[i:]
    for needle, n := range wordNumbers {
      if strings.HasPrefix(haystack, needle) {
	digits = append(digits, n)
	break
      }
    }
  }

  return digits[0] * 10 + digits[len(digits) - 1]
}
