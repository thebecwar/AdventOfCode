package aoc2015

import (
	"fmt"
)

func lookAndSay(s string) string {
	output := ""
	count := 1
	current := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != current {
			output += fmt.Sprintf("%d%c", count, current)
			count = 1
			current = s[i]
		} else {
			count++
		}
	}
	output += fmt.Sprintf("%d%c", count, current)
	return output
}

func Day10Part1() {
	input := "1113222113"
	for i := 0; i < 40; i++ {
		input = lookAndSay(input)
	}

	fmt.Printf("Day 10 Part 1: %d\n", len(input))
}

func Day10Part2() {
	input := "1113222113"
	for i := 0; i < 50; i++ {
		input = lookAndSay(input)
	}

	fmt.Printf("Day 10 Part 2: %d\n", len(input))
}
