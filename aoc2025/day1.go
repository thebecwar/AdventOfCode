package aoc2025

import (
	"advent/loader"
	"fmt"
	"strconv"
)

type day1CombinationCommand struct {
	Direction rune
	Amount    int
}

func day1ParseLine(line string) day1CombinationCommand {
	rawAmount := line[1:]
	amount, _ := strconv.Atoi(rawAmount)
	return day1CombinationCommand{
		Direction: rune(line[0]),
		Amount:    amount,
	}
}
func day1ParseLines(lines []string) []day1CombinationCommand {
	result := []day1CombinationCommand{}
	for _, l := range lines {
		result = append(result, day1ParseLine(l))
	}
	return result
}

func Day1Part1() {
	loader, err := loader.NewLoader("2025/day1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"L68",
			"L30",
			"R48",
			"L5",
			"R60",
			"L55",
			"L1",
			"L99",
			"R14",
			"L82",
		}
	*/
	instructions := day1ParseLines(loader.Lines)

	zeroes := 0
	current := 50
	for _, i := range instructions {
		if i.Direction == 'L' {
			current -= i.Amount
			for current < 0 {
				current += 100
			}
		} else {
			current += i.Amount
			for current >= 100 {
				current -= 100
			}
		}
		if current == 0 {
			zeroes++
		}
	}

	fmt.Printf("Day 1 Part 1: %d\n", zeroes)
}

func Day1Part2() {
	loader, err := loader.NewLoader("2025/day1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"L68",
			"L30",
			"R48",
			"L5",
			"R60",
			"L55",
			"L1",
			"L99",
			"R14",
			"L82",
		}
	*/
	instructions := day1ParseLines(loader.Lines)

	//5505 low
	//6649 high
	zeroes := 0
	current := 50

	for _, i := range instructions {
		zeroes += int(i.Amount / 100)
		remainder := i.Amount % 100

		if i.Direction == 'L' {
			if current > 0 && (current-remainder) <= 0 {
				zeroes++
			}
			current -= remainder
			if current < 0 {
				current += 100
			}
		} else {
			if (current + remainder) >= 100 {
				zeroes++
			}
			current += remainder
			if current > 99 {
				current -= 100
			}
		}
	}

	fmt.Printf("Day 1 Part 2: %d\n", zeroes)
}
