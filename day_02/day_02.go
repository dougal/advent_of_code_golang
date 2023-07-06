package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

  score := 0
  lines := 0

	for scanner.Scan() {
    score += scoreForLine(scanner.Text())
    lines++
    // fmt.Println(scanner.Text())
    // fmt.Println(scoreForLine(scanner.Text()))
  }
	// fmt.Println(lines)
	fmt.Println(score)
}

var opponentScores = map[string]int{
  "A": 1,
  "B": 2,
  "C": 3,
}

var playerScores = map[string]int{
  "X": 1,
  "Y": 2,
  "Z": 3,
}

func scoreForLine(l string) int {
  // fmt.Println(l)
  x := strings.Split(l, " ")
  o, p := x[0], x[1]
  // fmt.Println([]string{o, p})
  oScore := opponentScores[o]
  pScore := playerScores[p]
  // fmt.Println([]int{oScore, pScore})

  if oScore > pScore {
    return pScore
  }

  if oScore == pScore {
    return 3 + pScore
  }

  return 6 + pScore
}
