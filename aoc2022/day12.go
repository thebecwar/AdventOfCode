package aoc2022

import (
	"advent/loader"
	"fmt"
)

func Day12Part1() {
	loader, err := loader.NewLoader("2022/day12.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 12 Part 1: %d\n", 0)
}

func Day12Part2() {
	loader, err := loader.NewLoader("2022/day12.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 12 Part 2: %d\n", 0)
}
