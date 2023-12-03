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

  fmt.Println(partNumbersSum(f))
}

func partNumbersSum(input io.Reader) int {
  scanner := bufio.NewScanner(input)
  scanner.Split(bufio.ScanLines)

  s := 0
  pLine := ""
  cLine := ""
  nLine := ""

  for scanner.Scan() {
    pLine = cLine
    cLine = nLine
    nLine = scanner.Text()

    s += parsePartNumberSum(pLine, cLine, nLine)
  }

  // Handle the last line
  s += parsePartNumberSum(cLine, nLine, "")

  return s
}

var numRe = regexp.MustCompile(`\d+`)

func parsePartNumberSum(pLine string, cLine string, nLine string) int {
  // Find indexes of numbers in current line
  // Look for adjacent symbols in the previous, current, and next lines
  sum := 0
  partIndex := numRe.FindAllStringIndex(cLine, -1)

  for _, i := range partIndex {
    startRange := i[0]
    if startRange > 0 {
      startRange -= 1
    }

    endRange := i[1]
    if endRange < (len(cLine) - 1) {
      endRange += 1
    }

    if (len(pLine) > 0 && containsSymbol(pLine[startRange:endRange])) ||
         (len(cLine) > 0 && containsSymbol(cLine[startRange:endRange])) ||
         (len(nLine) > 0 && containsSymbol(nLine[startRange:endRange])) {
           d, err := strconv.Atoi(cLine[i[0]:i[1]])
           if err != nil {
             log.Fatal(err)
           }
           sum += d
         }
  }

  return sum
}

func containsSymbol(s string) bool {
  for _, r := range s {
    if r >= '0' && r <= '9' {
      continue
    } else if r == '.' {
      continue
    } else {
      return true
    }
  }

  return false
}
