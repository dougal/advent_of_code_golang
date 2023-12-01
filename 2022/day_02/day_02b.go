package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	score := 0

	for scanner.Scan() {
		score += scores[scanner.Text()]
	}

	fmt.Println(score)
}

var scores = map[string]int{
	"B X": 1, // Paper,    Lose, Rock,     0 + 1
	"C X": 2, // Scissors, Lose, Paper,    0 + 2
	"A X": 3, // Rock,     Lose, Scissors, 0 + 3
	"A Y": 4, // Rock,     Draw, Rock,     3 + 1
	"B Y": 5, // Paper,    Draw, Paper,    3 + 2
	"C Y": 6, // Scissors, Draw, Scissors, 3 + 3
	"C Z": 7, // Scissors, Win,  Rock,     6 + 1
	"A Z": 8, // Rock,     Win,  Paper,    6 + 2
	"B Z": 9, // Paper,    Win,  Scissors, 6 + 3
}
