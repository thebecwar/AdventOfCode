package aoc2024

import "fmt"

type Aoc24Dispatcher struct {
}

func (d *Aoc24Dispatcher) Run(day int) {
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
	default:
		fmt.Println("Day not implemented")
		return
	}
}
