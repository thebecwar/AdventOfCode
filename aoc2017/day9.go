package aoc2017

import (
	"advent/loader"
	"fmt"
)

func Day9Part1() {
	loader, err := loader.NewLoader("2017/day9.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 9 Part 1: %d\n", 0)
}

func Day9Part2() {
	loader, err := loader.NewLoader("2017/day9.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 9 Part 2: %d\n", 0)
}
