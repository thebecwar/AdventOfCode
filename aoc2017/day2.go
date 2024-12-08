package aoc2017

import (
	"advent/loader"
	"fmt"
)

func Day2Part1() {
	loader, err := loader.NewLoader("2017/day2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 2 Part 1: %d\n", 0)
}

func Day2Part2() {
	loader, err := loader.NewLoader("2017/day2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 2 Part 2: %d\n", 0)
}
