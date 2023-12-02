package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
  f, err := os.Open("../part_a/input_a.txt")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(sumPowers(f))
}

func sumPowers(lines io.Reader) int {
  scanner := bufio.NewScanner(lines)
  scanner.Split(bufio.ScanLines)

  s := 0

  for scanner.Scan() {
    line := scanner.Text()
    s += gamePower(line)
  }

  return s
}

// Returns 0 if the game is not possible with the set maximums
func gamePower(line string) int {
  // Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
  largestRed := 0
  largestGreen := 0
  largestBlue := 0

  allCombinations := strings.Split(line, ": ")[1]

  combinations := strings.Split(allCombinations, "; ")

  for _, c := range combinations {
    colourPairs := strings.Split(c, ", ")
    for _, p := range colourPairs {
      pairParts := strings.Split(p, " ")
      count, err := strconv.Atoi(pairParts[0])
      if err != nil {
        log.Fatal(err)
      }
      color := pairParts[1]

      switch color {
      case "red":
        if count > largestRed {
          largestRed = count
        }
      case "green":
        if count > largestGreen {
          largestGreen = count
        }
      case "blue":
        if count > largestBlue {
          largestBlue = count
        }
      }
    }
  }

  return largestRed * largestGreen * largestBlue
}
