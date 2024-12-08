package aoc2021

import (
	"advent/loader"
	"fmt"
)

func Day4Part1() {
	loader, err := loader.NewLoader("2021/day4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 4 Part 1: %d\n", 0)
}

func Day4Part2() {
	loader, err := loader.NewLoader("2021/day4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 4 Part 2: %d\n", 0)
}
