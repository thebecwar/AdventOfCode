package aoc2019

import (
	"advent/loader"
	"fmt"
)

func Day19Part1() {
	loader, err := loader.NewLoader("2019/day19.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 19 Part 1: %d\n", 0)
}

func Day19Part2() {
	loader, err := loader.NewLoader("2019/day19.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 19 Part 2: %d\n", 0)
}
