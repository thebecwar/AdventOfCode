
package aoc2017

import (
	"fmt"
	"time"
)

type Aoc17Dispatcher struct {
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func (d *Aoc17Dispatcher) Run(day int) {
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
	case 13:
		Day13Part1()
		Day13Part2()
	case 14:
		Day14Part1()
		Day14Part2()
	case 15:
		Day15Part1()
		Day15Part2()
	case 16:
		Day16Part1()
		Day16Part2()
	case 17:
		Day17Part1()
		Day17Part2()
	case 18:
		Day18Part1()
		Day18Part2()
	case 19:
		Day19Part1()
		Day19Part2()
	case 20:
		Day20Part1()
		Day20Part2()
	case 21:
		Day21Part1()
		Day21Part2()
	case 22:
		Day22Part1()
		Day22Part2()
	case 23:
		Day23Part1()
		Day23Part2()
	case 24:
		Day24Part1()
		Day24Part2()
	case 25:
		Day25Part1()
		Day25Part2()
	default:
		fmt.Println("Day not implemented")
		return
	}
}
