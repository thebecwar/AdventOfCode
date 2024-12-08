package aoc2020

import (
	"advent/loader"
	"fmt"
)

func Day11Part1() {
	loader, err := loader.NewLoader("2020/day11.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 11 Part 1: %d\n", 0)
}

func Day11Part2() {
	loader, err := loader.NewLoader("2020/day11.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 11 Part 2: %d\n", 0)
}
