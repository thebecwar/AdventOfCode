package aoc2017

import (
	"advent/loader"
	"fmt"
)

func Day3Part1() {
	loader, err := loader.NewLoader("2017/day3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 3 Part 1: %d\n", 0)
}

func Day3Part2() {
	loader, err := loader.NewLoader("2017/day3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 3 Part 2: %d\n", 0)
}
