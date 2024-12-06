package aoc2024

import (
	"fmt"
	"time"
)

type Aoc24Dispatcher struct {
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func (d *Aoc24Dispatcher) Run(day int) {
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
	default:
		fmt.Println("Day not implemented")
		return
	}
}
