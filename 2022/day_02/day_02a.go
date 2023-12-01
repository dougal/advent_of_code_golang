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
	"B X": 1, // Paper, Rock, Lose
	"C Y": 2, // Scissors, Paper, Lose
	"A Z": 3, // Rock, Scissors, Lose
	"A X": 4, // Rock, Rock, Draw
	"B Y": 5, // Paper, Paper, Draw
	"C Z": 6, // Scissors, Scissors, Draw
	"C X": 7, // Scissors, Rock, Win
	"A Y": 8, // Rock, Paper, Win
	"B Z": 9, // Paper, Scissors, Win
}
