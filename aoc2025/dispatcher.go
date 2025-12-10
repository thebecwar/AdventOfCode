package aoc2025

import (
	"fmt"
	"time"
)

type Aoc25Dispatcher struct {
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func (d *Aoc25Dispatcher) Run(day int) {
	defer timer(fmt.Sprintf("Day %d", day))()
	switch day {
	case 1:
		Day1Part1()
		Day1Part2()
	case 2:
		Day2Part1()
		Day2Part2()
	case 3:
		Day3Part1()
		Day3Part2()
	case 4:
		Day4Part1()
		Day4Part2()
	case 5:
		Day5Part1()
		Day5Part2()
	case 6:
		Day6Part1()
		Day6Part2()
	case 7:
		Day7Part1()
		Day7Part2()
	case 8:
		Day8Part1()
		Day8Part2()
	case 9:
		Day9Part1()
		Day9Part2()
	case 10:
		Day10Part1()
		Day10Part2()
	case 11:
		Day11Part1()
		Day11Part2()
	case 12:
		Day12Part1()
		Day12Part2()
	default:
		fmt.Println("Day not implemented")
		return
	}
}
