package main

import (
	"slices"
	"strings"
	"testing"
)

const input = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func TestLowestLocation(t *testing.T) {
	result := lowestLocation(strings.NewReader(input))
	expected := 35

	if result != expected {
		t.Fatalf("Expected seeds to be %d but was %d", expected, result)
	}
}

func Test(t *testing.T) {
	seeds, mappings := parse(strings.NewReader(input))

	expSeeds := []int{79, 14, 55, 13}

	if !slices.Equal(seeds, expSeeds) {
		t.Fatalf("Expected seeds to be %v but was %v", expSeeds, seeds)
	}

	if len(mappings) != 7 {
		t.Fatalf("Expected mappings length to be 7 but was %d", len(mappings))
	}

	m := mappings[5]
	expName := "temperature-to-humidity"
	expRangeOne := rng{destStart: 0, sourceStart: 69, length: 1}
	expRangeTwo := rng{destStart: 1, sourceStart: 0, length: 69}

	if expName != m.name {
		t.Fatalf("Expected name to be %s but was %s", expName, m.name)
	}

	if expRangeOne != m.ranges[0] {
		t.Fatalf("Expected range to be %v but was %v", expRangeOne, m.ranges[0])
	}

	if expRangeTwo != m.ranges[1] {
		t.Fatalf("Expected range to be %v but was %v", expRangeTwo, m.ranges[1])
	}
}
