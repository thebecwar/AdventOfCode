package aoc2021

import (
	"advent/loader"
	"fmt"
)

func Day24Part1() {
	loader, err := loader.NewLoader("2021/day24.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 24 Part 1: %d\n", 0)
}

func Day24Part2() {
	loader, err := loader.NewLoader("2021/day24.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 24 Part 2: %d\n", 0)
}
