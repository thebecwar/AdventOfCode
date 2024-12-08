package aoc2015

import (
	"advent/loader"
	"fmt"
)

func Day18Part1() {
	loader, err := loader.NewLoader("2015/day18.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 18 Part 1: %d\n", 0)
}

func Day18Part2() {
	loader, err := loader.NewLoader("2015/day18.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 18 Part 2: %d\n", 0)
}
