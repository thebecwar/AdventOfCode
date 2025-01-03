package aoc2018

import (
	"advent/loader"
	"fmt"
)

func Day14Part1() {
	loader, err := loader.NewLoader("2018/day14.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 14 Part 1: %d\n", 0)
}

func Day14Part2() {
	loader, err := loader.NewLoader("2018/day14.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 14 Part 2: %d\n", 0)
}
