package main

import (
	"bufio"
	"fmt"
	"os"
  "strconv"
  "strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

  eclipsedSections := 0

  for scanner.Scan() {
    elvesSections := strings.Split(scanner.Text(), ",")
    elfALimits := strings.Split(elvesSections[0], "-")
    elfBLimits := strings.Split(elvesSections[1], "-")

    a, _ := strconv.Atoi(elfALimits[0])
    b, _ := strconv.Atoi(elfALimits[1])

    c, _ := strconv.Atoi(elfBLimits[0])
    d, _ := strconv.Atoi(elfBLimits[1])

    if a <= c && b >= d || c <= a && d >= b {
      eclipsedSections++
    }
  }

  fmt.Println(eclipsedSections)
}
