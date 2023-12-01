package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
  fmt.Println(summation())
}

// Process each line for the first and last digit.
// String concat these digits to form a two digit number.
// Sum the numbers from every line
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
    digits := []rune{}

    for _, r := range line {
      if r >= '0' && r <= '9' {
	digits = append(digits, r)
      }
    }

    val, err := strconv.Atoi(string([]rune{digits[0], digits[len(digits) - 1]}))
    if err != nil {
      log.Fatal(err)
    }

    sum += val
  }

  return sum
}
