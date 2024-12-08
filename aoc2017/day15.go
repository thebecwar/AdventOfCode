package aoc2017

import (
	"advent/loader"
	"fmt"
)

func Day15Part1() {
	loader, err := loader.NewLoader("2017/day15.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 15 Part 1: %d\n", 0)
}

func Day15Part2() {
	loader, err := loader.NewLoader("2017/day15.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 15 Part 2: %d\n", 0)
}
