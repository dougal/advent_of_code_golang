package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(LavaVolume(f))
}

func LavaVolume(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var steps []Step

	for scanner.Scan() {
		steps = append(steps, NewStep(scanner.Text()))
	}

	// Work out the matrix size, and the start position
	x, y := 1, 1

	var vertices []Point
	var trenchLen int

	for _, s := range steps {
		switch s.direction {
		case "U":
			y -= s.distance
		case "D":
			y += s.distance
		case "L":
			x -= s.distance
		case "R":
			x += s.distance
		}
		vertices = append(vertices, Point{x, y})
		trenchLen += s.distance
	}

	return shoeLace(vertices) + trenchLen/2 + 1
}

func shoeLace(vertices []Point) int {
	var area int

	j := len(vertices) - 1

	for i := 0; i < len(vertices); i++ {
		area += (vertices[j].x + vertices[i].x) * (vertices[j].y - vertices[i].y)
		j = i
	}

	return Abs(area / 2)
}

func Abs(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}

type Point struct {
	x int
	y int
}

type Step struct {
	direction string
	distance  int
	// colour    string
}

var directions = map[int]string{
	0: "R",
	1: "D",
	2: "L",
	3: "U",
}

func NewStep(in string) Step {
	s := Step{}
	inLen := len(in)

	dirI, err := strconv.ParseInt(in[inLen-2:inLen-1], 16, 0)
	if err != nil {
		log.Fatal(err)
	}
	s.direction = directions[int(dirI)]

	distI, err := strconv.ParseInt(in[inLen-7:inLen-2], 16, 0)
	if err != nil {
		log.Fatal(err)
	}
	s.distance = int(distI)

	return s
}
