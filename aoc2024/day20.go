package aoc2024

import (
	"advent/loader"
	"fmt"
)

func Day20Part1() {
	loader, err := loader.NewLoader("2024/day20.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 20 Part 1: %d\n", 0)
}

func Day20Part2() {
	loader, err := loader.NewLoader("2024/day20.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 20 Part 2: %d\n", 0)
}
