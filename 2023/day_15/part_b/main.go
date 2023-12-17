package main

import (
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

	fmt.Println(SumFocussingPower(f))
}

func SumFocussingPower(input io.Reader) int {

	all, _ := io.ReadAll(input)
	steps := strings.Split(string(all), ",")
	boxes := make([]Box, 256)

	for _, s := range steps {
		label, op, focalLength := parseStep(s)
		i := hash(label)

		switch op {
		case "-":
			boxes[i] = boxes[i].RemoveLens(label)
		case "=":
			boxes[i] = boxes[i].AddLens(label, focalLength)
		}
	}

	s := 0
	for i, b := range boxes {
		s += b.SumFocussingPowers(i + 1)
	}
	return s
}

type Lens struct {
	label       string
	focalLength int
}

type Box []Lens

func (b Box) SumFocussingPowers(boxNum int) int {
	var s int

	for i, lens := range b {
		s += boxNum * (i + 1) * lens.focalLength
	}

	return s
}

func (b Box) AddLens(label string, focalLength int) Box {
	lens := Lens{label, focalLength}

	existingInd := -1
	for i, l := range b {
		if l.label == label {
			existingInd = i
			break
		}
	}

	if existingInd == -1 {
		b = append(b, lens)
	} else {
		b[existingInd] = lens
	}

	return b
}

func (b Box) RemoveLens(label string) Box {
	var newBox Box

	for _, lens := range b {
		if lens.label != label {
			newBox = append(newBox, lens)
		}
	}

	return newBox
}

func parseStep(in string) (string, string, int) {
	var (
		label, op, pow string
		offset         int
	)

	for i, c := range in {
		if c == '-' || c == '=' {
			op = string(c)
			offset = i
			break
		}

		label += string(c)
	}

	pow, _ = strings.CutSuffix(string(in[offset+1:]), "\n")

	var powi int
	var err error
	if op == "=" {
		powi, err = strconv.Atoi(pow)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		powi = -1
	}

	return label, op, powi
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
