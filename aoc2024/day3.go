package aoc2024

import (
	"advent/loader"
	"fmt"
	"regexp"
	"strconv"
)

func Day3Part1() {
	loader, err := loader.NewLoader("2024/day3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
		}
	*/

	sum := 0
	count := 0
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for _, line := range loader.Lines {
		matches := regex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			a, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Println(err)
				return
			}
			b, err := strconv.Atoi(match[2])
			if err != nil {
				fmt.Println(err)
				return
			}
			sum += a * b
			count++
		}
	}

	fmt.Printf("Day 3, Part 1: %d (Count: %d)\n", sum, count)
}
func Day3Part2() {
	loader, err := loader.NewLoader("2024/day3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
	}*/

	sum := 0
	count := 0
	skipped := 0
	do := true
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	for _, line := range loader.Lines {
		matches := regex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				do = true
				continue
			}
			if match[0] == "don't()" {
				do = false
				continue
			}
			if !do {
				skipped++
				continue
			}
			a, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Println(err)
				return
			}
			b, err := strconv.Atoi(match[2])
			if err != nil {
				fmt.Println(err)
				return
			}
			sum += a * b
			count++
		}
	}

	fmt.Printf("Day 3, Part 1: %d (Count: %d, Skipped: %d)\n", sum, count, skipped)
}
