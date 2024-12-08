package aoc2017

import (
	"advent/loader"
	"fmt"
)

func Day4Part1() {
	loader, err := loader.NewLoader("2017/day4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 4 Part 1: %d\n", 0)
}

func Day4Part2() {
	loader, err := loader.NewLoader("2017/day4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 4 Part 2: %d\n", 0)
}
