#!/bin/bash -xe

dir=$1

mkdir "$dir"

mkdir "$dir/part_a"
cd "$dir/part_a"
go mod init dougal/advent_of_code_golang/2023/part_a
cd -
echo "package main

func main() {
}
" > "$dir/part_a/main.go"
echo "package main

// func TestSomething(t *testing.T) {
// }
" > "$dir/part_a/main_test.go"

mkdir "$dir/part_b"
cd "$dir/part_b"
go mod init dougal/advent_of_code_golang/2023/part_b
cd -
echo "package main

func main() {
}
" > "$dir/part_b/main.go"
echo "package main

// func TestSomething(t *testing.T) {
// }
" > "$dir/part_b/main_test.go"
