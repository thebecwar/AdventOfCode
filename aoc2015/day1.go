package aoc2015

import (
	"advent/loader"
	"fmt"
)

func Day1Part1() {
	loader, err := loader.NewLoader("2015/day1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	floor := 0
	for _, line := range loader.Lines {
		for _, c := range line {
			if c == '(' {
				floor++
			} else if c == ')' {
				floor--
			}
		}
	}

	fmt.Printf("Day 1 Part 1: %d\n", floor)
}

func Day1Part2() {
	loader, err := loader.NewLoader("2015/day1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	floor := 0
	position := 0
	for _, line := range loader.Lines {
		for _, c := range line {
			position++
			if c == '(' {
				floor++
			} else if c == ')' {
				floor--
			}
			if floor < 0 {
				break
			}
		}
	}

	fmt.Printf("Day 1 Part 2: %d\n", position)
}
