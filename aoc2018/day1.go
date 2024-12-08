package aoc2018

import (
	"advent/loader"
	"fmt"
)

func Day1Part1() {
	loader, err := loader.NewLoader("2018/day1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 1 Part 1: %d\n", 0)
}

func Day1Part2() {
	loader, err := loader.NewLoader("2018/day1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 1 Part 2: %d\n", 0)
}
