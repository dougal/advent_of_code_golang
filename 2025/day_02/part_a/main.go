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
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sumInvalidIds(f))
}

func sumInvalidIds(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	s, err := io.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}

	ss := strings.TrimSuffix(string(s), "\n")

	var t int

	ranges := strings.Split(ss, ",")
	for _, r := range ranges {
		t += rangeSumInvalid(r)
	}

	return t
}

func rangeSumInvalid(s string) int {
	parts := strings.Split(s, "-")
	fromS, err := strconv.ParseInt(parts[0], 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	toS, err := strconv.ParseInt(parts[1], 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	var c int
	for i := int(fromS); i <= int(toS); i++ {
		if !validId(i) {
			c += i
		}
	}

	return c
}

// Could probably make this faster with maths.
func validId(n int) bool {
	s := strconv.Itoa(n)
	l := len(s)

	// Valid if length is odd
	if l%2 == 1 {
		return true
	}

	h := l / 2
	return s[0:h] != s[h:]
}
