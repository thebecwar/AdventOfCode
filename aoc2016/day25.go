package aoc2016

import (
	"advent/loader"
	"fmt"
)

func Day25Part1() {
	loader, err := loader.NewLoader("2016/day25.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 25 Part 1: %d\n", 0)
}

func Day25Part2() {
	loader, err := loader.NewLoader("2016/day25.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 25 Part 2: %d\n", 0)
}
