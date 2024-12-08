package aoc2015

import (
	"advent/loader"
	"fmt"
	"sort"
)

func parseWrapping(s string) (int, int, int) {
	l, w, h := 0, 0, 0
	fmt.Sscanf(s, "%dx%dx%d", &l, &w, &h)
	return l, w, h
}

func Day2Part1() {
	loader, err := loader.NewLoader("2015/day2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	totalArea := 0

	for _, line := range loader.Lines {
		l, w, h := parseWrapping(line)
		area := 2*l*w + 2*w*h + 2*h*l
		extra := min(l*w, min(w*h, h*l))
		totalArea += area + extra
	}

	fmt.Printf("Day 2 Part 1: %d\n", totalArea)
}

func Day2Part2() {
	loader, err := loader.NewLoader("2015/day2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	ribbon := 0
	for _, line := range loader.Lines {
		l, w, h := parseWrapping(line)
		sides := []int{l, w, h}
		sort.Ints(sides)
		ribbon += 2*sides[0] + 2*sides[1] + l*w*h
	}

	fmt.Printf("Day 2 Part 2: %d\n", ribbon)
}
