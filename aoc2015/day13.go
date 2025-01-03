package aoc2015

import (
	"advent/loader"
	"fmt"
)

func Day13Part1() {
	loader, err := loader.NewLoader("2015/day13.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 13 Part 1: %d\n", 0)
}

func Day13Part2() {
	loader, err := loader.NewLoader("2015/day13.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 13 Part 2: %d\n", 0)
}
