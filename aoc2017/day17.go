package aoc2017

import (
	"advent/loader"
	"fmt"
)

func Day17Part1() {
	loader, err := loader.NewLoader("2017/day17.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 17 Part 1: %d\n", 0)
}

func Day17Part2() {
	loader, err := loader.NewLoader("2017/day17.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 17 Part 2: %d\n", 0)
}
