package aoc2021

import (
	"advent/loader"
	"fmt"
)

func Day23Part1() {
	loader, err := loader.NewLoader("2021/day23.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 23 Part 1: %d\n", 0)
}

func Day23Part2() {
	loader, err := loader.NewLoader("2021/day23.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 23 Part 2: %d\n", 0)
}
