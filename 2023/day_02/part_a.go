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

const maxRed = 12
const maxGreen = 13
const maxBlue = 14

func main() {
  f, err := os.Open("input_a.txt")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(sum(f))
}

func sum(lines io.Reader) int {
  scanner := bufio.NewScanner(lines)
  scanner.Split(bufio.ScanLines)

  s := 0

  for scanner.Scan() {
    line := scanner.Text()
    s += possibleLineId(line)
  }

  return s
}

// Returns 0 if the game is not possible with the set maximums
func possibleLineId(line string) int {
  // Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

  mainParts := strings.Split(line, ": ")
  gameSuffix := mainParts[0]
  allCombinations := mainParts[1]

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
        if count > maxRed {
          return 0
        }
      case "green":
        if count > maxGreen {
          return 0
        }
      case "blue":
        if count > maxBlue {
          return 0
        }
      }
    }
  }

  gameParts := strings.Split(gameSuffix, " ")
  gameID, err := strconv.Atoi(gameParts[1])
  if err != nil {
    log.Fatal(err)
  }

  return gameID
}
