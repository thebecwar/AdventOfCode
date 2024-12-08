package aoc2020

import (
	"advent/loader"
	"fmt"
)

func Day6Part1() {
	loader, err := loader.NewLoader("2020/day6.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 6 Part 1: %d\n", 0)
}

func Day6Part2() {
	loader, err := loader.NewLoader("2020/day6.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 6 Part 2: %d\n", 0)
}
