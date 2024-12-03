package aoc2023

import (
	"advent/loader"
	"fmt"
	"strconv"
	"strings"
)

func parseInputLine(line string) []int {
	numbers := []int{}
	matches := strings.Split(strings.TrimSpace(line), " ")
	for _, match := range matches {
		n, _ := strconv.Atoi(match)
		numbers = append(numbers, n)
	}
	return numbers
}

func calculateDerivitive(numbers []int) ([]int, bool) {
	derivitive := []int{}
	allZero := true
	for i := 1; i < len(numbers); i++ {
		n := numbers[i] - numbers[i-1]
		if n != 0 {
			allZero = false
		}
		derivitive = append(derivitive, numbers[i]-numbers[i-1])
	}
	return derivitive, allZero
}

func predictRecursive(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	} else if len(numbers) == 1 {
		return numbers[0]
	}

	derivitive, allZero := calculateDerivitive(numbers)
	if allZero {
		return numbers[len(numbers)-1]
	}

	next := predictRecursive(derivitive)
	return numbers[len(numbers)-1] + next
}
func backPredictRecursive(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	} else if len(numbers) == 1 {
		return numbers[0]
	}

	derivitive, allZero := calculateDerivitive(numbers)
	if allZero {
		return numbers[0]
	}
	next := backPredictRecursive(derivitive)
	return numbers[0] - next
}

func Day9Part1() {
	loader, err := loader.NewLoader("2023-day9-part1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	/*
		loader.Lines = []string{
			"0 3 6 9 12 15",
			"1 3 6 10 15 21",
			"10 13 16 21 30 45",
			"0 1 1 0",
		}
	*/

	sum := 0
	for _, line := range loader.Lines {
		parsed := parseInputLine(line)
		step := predictRecursive(parsed)
		sum += step
	}

	fmt.Printf("Day 9, Part 1: %d\n", sum)
}
func Day9Part2() {
	loader, err := loader.NewLoader("2023-day9-part1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"0 3 6 9 12 15",
			"1 3 6 10 15 21",
			"10 13 16 21 30 45",
		}
	*/

	sum := 0
	for _, line := range loader.Lines {
		parsed := parseInputLine(line)
		step := backPredictRecursive(parsed)
		sum += step
	}

	fmt.Printf("Day 9, Part 2: %d\n", sum)
}
