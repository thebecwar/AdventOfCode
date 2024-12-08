package aoc2024

import (
	"advent/loader"
	"fmt"
)

func Day8Part1() {
	loader, err := loader.NewLoader("2024/day8.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 8 Part 1: %d\n", 0)
}

func Day8Part2() {
	loader, err := loader.NewLoader("2024/day8.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 8 Part 2: %d\n", 0)
}
