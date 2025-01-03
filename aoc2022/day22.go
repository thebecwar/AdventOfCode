package aoc2022

import (
	"advent/loader"
	"fmt"
)

func Day22Part1() {
	loader, err := loader.NewLoader("2022/day22.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 22 Part 1: %d\n", 0)
}

func Day22Part2() {
	loader, err := loader.NewLoader("2022/day22.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 22 Part 2: %d\n", 0)
}
