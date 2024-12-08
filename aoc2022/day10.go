package aoc2022

import (
	"advent/loader"
	"fmt"
)

func Day10Part1() {
	loader, err := loader.NewLoader("2022/day10.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 10 Part 1: %d\n", 0)
}

func Day10Part2() {
	loader, err := loader.NewLoader("2022/day10.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 10 Part 2: %d\n", 0)
}
