package aoc2015

import (
	"advent/loader"
	"fmt"
	"regexp"
)

func inMemorySize(s string) int {
	withoutQuotes := s[1 : len(s)-1]

	re := regexp.MustCompile(`\\\\|\\"|\\x[0-9a-f]{2}`)
	withoutEscapes := re.ReplaceAllString(withoutQuotes, "X")

	return len(withoutEscapes)
}
func encodedSize(s string) int {
	total := 0

	for _, c := range s {
		if c == '\\' || c == '"' {
			total += 2
		} else {
			total++
		}
	}

	return total + 2 // include start and end quote
}

func Day8Part1() {
	loader, err := loader.NewLoader("2015/day8.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	total := 0
	for _, line := range loader.Lines {
		raw := len(line)
		mem := inMemorySize(line)
		total += raw - mem
	}

	fmt.Printf("Day 8 Part 1: %d\n", total)
}

func Day8Part2() {
	loader, err := loader.NewLoader("2015/day8.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	total := 0
	for _, line := range loader.Lines {
		raw := len(line)
		enc := encodedSize(line)
		total += enc - raw

	}

	fmt.Printf("Day 8 Part 2: %d\n", total)
}
