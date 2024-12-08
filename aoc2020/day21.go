package aoc2020

import (
	"advent/loader"
	"fmt"
)

func Day21Part1() {
	loader, err := loader.NewLoader("2020/day21.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 21 Part 1: %d\n", 0)
}

func Day21Part2() {
	loader, err := loader.NewLoader("2020/day21.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 21 Part 2: %d\n", 0)
}
