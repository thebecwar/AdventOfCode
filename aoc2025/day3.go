package aoc2025

import (
	"advent/loader"
	"fmt"
)

func Day3Part1() {
	loader, err := loader.NewLoader("2025/day3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"987654321111111",
			"811111111111119",
			"234234234234278",
			"818181911112111",
		}
	*/
	joltage := func(s string) int {
		max := 0
		maxIdx := 0
		for i := 0; i < len(s)-1; i++ {
			if int(s[i]-'0') > max {
				max = int(s[i] - '0')
				maxIdx = i
				if max == 9 {
					break
				}
			}
		}
		secondDigit := 0
		for i := maxIdx + 1; i < len(s); i++ {
			if int(s[i]-'0') > secondDigit {
				secondDigit = int(s[i] - '0')
				if secondDigit == 9 {
					break
				}
			}
		}
		return max*10 + secondDigit
	}
	total := 0
	for _, s := range loader.Lines {
		jolt := joltage(s)
		total += jolt
	}

	fmt.Printf("Day 3 Part 1: %d\n", total)
}

func Day3Part2() {
	loader, err := loader.NewLoader("2025/day3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"987654321111111",
			"811111111111119",
			"234234234234278",
			"818181911112111",
		}
	*/
	maxInRange := func(s string, start, end int) (int, int) {
		max := 0
		maxIdx := 0
		for i := start; i <= end; i++ {
			if int(s[i]-'0') > max {
				max = int(s[i] - '0')
				maxIdx = i
				if max == 9 {
					return max, maxIdx
				}
			}
		}
		return max, maxIdx
	}
	total := 0
	for _, l := range loader.Lines {
		multiplier := 100000000000
		jolt := 0
		start := 0
		for i := 0; i < 12; i++ {
			end := len(l) - (12 - i)
			max, idx := maxInRange(l, start, end)
			start = idx + 1
			jolt += max * multiplier
			multiplier = multiplier / 10
		}
		total += jolt
	}

	fmt.Printf("Day 3 Part 2: %d\n", total)
}
