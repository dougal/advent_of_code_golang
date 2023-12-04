#!/bin/bash -xe

dir=$1

mkdir "$dir"

touch "$dir/input.txt"

# Part A
########
mkdir "$dir/part_a"
cd "$dir/part_a"
go mod init dougal/advent_of_code_golang/2023/part_a
echo "package main

import (
  \"bufio\"
  \"fmt\"
  \"io\"
  \"log\"
  \"os\"
)

func main() {
  f, err := os.Open(\"../input.txt\")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(doSomething(f))
}

func doSomething(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	s := 0

	for scanner.Scan() {
		// line := scanner.Text()
  }

  return s
}
" > "main.go"

echo "package main

import (
  \"testing\"
)

// func TestSomething(t *testing.T) {
// }
" > "main_test.go"

go fmt .

cd -

# Part B
########
mkdir "$dir/part_b"
cd "$dir/part_b"
go mod init dougal/advent_of_code_golang/2023/part_b
echo "package main

import (
  \"bufio\"
  \"fmt\"
  \"io\"
  \"log\"
  \"os\"
)

func main() {
  f, err := os.Open(\"../input.txt\")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(doSomething(f))
}

func doSomething(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	s := 0

	for scanner.Scan() {
		// line := scanner.Text()
  }

  return s
}
" > "main.go"

echo "package main

import (
  \"testing\"
)

// func TestSomething(t *testing.T) {
// }
" > "main_test.go"

go fmt .

cd -
