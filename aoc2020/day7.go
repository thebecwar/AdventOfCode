package aoc2020

import (
	"advent/loader"
	"fmt"
)

func Day7Part1() {
	loader, err := loader.NewLoader("2020/day7.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 7 Part 1: %d\n", 0)
}

func Day7Part2() {
	loader, err := loader.NewLoader("2020/day7.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 7 Part 2: %d\n", 0)
}
