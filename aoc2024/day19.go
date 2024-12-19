package aoc2024

import (
	"advent/loader"
	"fmt"
	"strings"
)

func canBeConstructed(s string, prefixList []string, memo map[string]bool) bool {
	if _, ok := memo[s]; ok {
		return memo[s]
	}
	if s == "" {
		return true
	}
	for _, prefix := range prefixList {
		if len(prefix) > len(s) {
			continue
		}
		if s[:len(prefix)] == prefix {
			if canBeConstructed(s[len(prefix):], prefixList, memo) {
				memo[s] = true
				return true
			}
		}
	}
	memo[s] = false
	return false
}
func sumOfWaysToConstruct(s string, prefixList []string, memo map[string]int) int {
	if _, ok := memo[s]; ok {
		return memo[s]
	}
	if s == "" {
		return 1
	}
	result := 0
	for _, prefix := range prefixList {
		if len(prefix) > len(s) {
			continue
		}
		if s[:len(prefix)] == prefix {
			result += sumOfWaysToConstruct(s[len(prefix):], prefixList, memo)
		}
	}
	memo[s] = result
	return result
}

func Day19Part1() {
	loader, err := loader.NewLoader("2024/day19.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"r, wr, b, g, bwu, rb, gb, br",
		"",
		"brwrr",
		"bggr",
		"gbbr",
		"rrbgbr",
		"ubwu",
		"bwurrg",
		"brgr",
		"bbrgwb",
	}*/
	prefixes := strings.Split(loader.Lines[0], ", ")
	memo := map[string]bool{}
	result := 0
	for _, line := range loader.Lines[2:] {
		if canBeConstructed(line, prefixes, memo) {
			result++
		}
	}

	fmt.Printf("Day 19 Part 1: %d\n", result)
}

func Day19Part2() {
	loader, err := loader.NewLoader("2024/day19.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"r, wr, b, g, bwu, rb, gb, br",
		"",
		"brwrr",
		"bggr",
		"gbbr",
		"rrbgbr",
		"ubwu",
		"bwurrg",
		"brgr",
		"bbrgwb",
	}*/

	prefixes := strings.Split(loader.Lines[0], ", ")
	memo := map[string]int{}
	result := 0

	for _, line := range loader.Lines[2:] {
		next := sumOfWaysToConstruct(line, prefixes, memo)
		//fmt.Printf("Line: %s, Ways: %d\n", line, next)
		result += next
	}

	fmt.Printf("Day 19 Part 2: %d\n", result)
}
