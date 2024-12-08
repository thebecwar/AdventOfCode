package aoc2015

import (
	"advent/loader"
	"fmt"
)

func hasNVowels(s string, n int) bool {
	count := 0
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			count++
			if count >= n {
				return true
			}
		}
	}
	return false
}
func hasDoubleLetter(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			return true
		}
	}
	return false
}
func hasForbiddenStrings(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i-1] == 'a' && s[i] == 'b' {
			return true
		}
		if s[i-1] == 'c' && s[i] == 'd' {
			return true
		}
		if s[i-1] == 'p' && s[i] == 'q' {
			return true
		}
		if s[i-1] == 'x' && s[i] == 'y' {
			return true
		}
	}
	return false
}

func isNiceStringPart1(s string) bool {
	// three vowels
	if !hasNVowels(s, 3) {
		return false
	}
	if !hasDoubleLetter(s) {
		return false
	}
	if hasForbiddenStrings(s) {
		return false
	}

	return true
}

func hasPairOfLetters(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		a := s[i]
		b := s[i+1]
		for j := i + 2; j < len(s)-1; j++ {
			c := s[j]
			d := s[j+1]
			if a == c && b == d {
				return true
			}
		}
	}
	return false
}
func hasABA(s string) bool {
	for i := 2; i < len(s); i++ {
		if s[i-2] == s[i] {
			return true
		}
	}
	return false
}

func isNiceStringPart2(s string) bool {
	if !hasPairOfLetters(s) {
		return false
	}
	if !hasABA(s) {
		return false
	}
	return true
}

func Day5Part1() {
	loader, err := loader.NewLoader("2015/day5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	niceStrings := 0
	for _, line := range loader.Lines {
		if isNiceStringPart1(line) {
			niceStrings++
		}
	}

	fmt.Printf("Day 5 Part 1: %d\n", niceStrings)
}

func Day5Part2() {
	loader, err := loader.NewLoader("2015/day5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	niceStrings := 0
	for _, line := range loader.Lines {
		if isNiceStringPart2(line) {
			niceStrings++
		}
	}

	fmt.Printf("Day 5 Part 2: %d\n", niceStrings)
}
