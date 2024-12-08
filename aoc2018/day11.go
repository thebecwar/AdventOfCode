package aoc2018

import (
	"advent/loader"
	"fmt"
)

func Day11Part1() {
	loader, err := loader.NewLoader("2018/day11.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 11 Part 1: %d\n", 0)
}

func Day11Part2() {
	loader, err := loader.NewLoader("2018/day11.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 11 Part 2: %d\n", 0)
}
