package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(SumHashes(f))
}

func SumHashes(input io.Reader) int {
	s := 0

	all, _ := io.ReadAll(input)
	words := strings.Split(string(all), ",")

	for _, w := range words {
		s += hash(w)
	}

	return s
}

func hash(in string) int {
	var s int

	for _, c := range in {
		if string(c) == "\n" {
			continue
		}

		s += int(c)
		s *= 17
		s = s % 256
	}

	return s
}
