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

	for i:=2; i<=l; i++ {
		r := isRepeat(s, i)
		if r {
			return false
		}
	}

	return true
}

func isRepeat(s string, n int) bool {
	l := len(s)

	if l % n != 0 {
		return false
	}

	p := l / n

	f := s[0:p]

	for i:=1; i<n; i++ {
		g := s[p*i:p+p*i]

		if g != f {
			return false
		}
	}

	return true
}
