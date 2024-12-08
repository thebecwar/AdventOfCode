package aoc2015

import (
	"advent/loader"
	"fmt"
)

func Day16Part1() {
	loader, err := loader.NewLoader("2015/day16.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 16 Part 1: %d\n", 0)
}

func Day16Part2() {
	loader, err := loader.NewLoader("2015/day16.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 16 Part 2: %d\n", 0)
}
