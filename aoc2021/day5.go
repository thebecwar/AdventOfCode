package aoc2021

import (
	"advent/loader"
	"fmt"
)

func Day5Part1() {
	loader, err := loader.NewLoader("2021/day5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 5 Part 1: %d\n", 0)
}

func Day5Part2() {
	loader, err := loader.NewLoader("2021/day5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 5 Part 2: %d\n", 0)
}
