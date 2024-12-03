package aoc2024

import (
	"advent/loader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func isSafe(numbers []int) bool {
	if len(numbers) == 0 {
		return false
	}
	ascending := false
	for i := 0; i < len(numbers)-1; i++ {
		// Condition 1 - all ascending or all descending
		if i == 0 {
			ascending = numbers[i] < numbers[i+1]
		} else {
			if ascending && numbers[i] > numbers[i+1] {
				return false
			} else if !ascending && numbers[i] < numbers[i+1] {
				return false
			}
		}
		// condition 2 - Max delta 3
		delta := math.Abs(float64(numbers[i] - numbers[i+1]))
		if delta == 0 || delta > 3 {
			return false
		}
	}
	return true
}

func Day2Part1() {
	loader, err := loader.NewLoader("2024-day2-part1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}*/

	safe := 0
	for _, line := range loader.Lines {
		raw := strings.Split(line, " ")
		numbers := []int{}
		for i := 0; i < len(raw); i++ {
			n, _ := strconv.Atoi(raw[i])
			numbers = append(numbers, n)
		}
		if isSafe(numbers) {
			safe++
		}
	}
	fmt.Printf("Day 2 Part 1: %d\n", safe)
}
func Day2Part2() {
	loader, err := loader.NewLoader("2024-day2-part1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}*/

	safe := 0
	for _, line := range loader.Lines {
		raw := strings.Split(line, " ")
		numbers := []int{}
		for i := 0; i < len(raw); i++ {
			n, _ := strconv.Atoi(raw[i])
			numbers = append(numbers, n)
		}
		if isSafe(numbers) {
			safe++
		} else {
			for i := range numbers {
				copied := []int{}
				if i > 0 {
					copied = append(copied, numbers[:i]...)
				}
				copied = append(copied, numbers[i+1:]...)
				if isSafe(copied) {
					safe++
					break
				}
			}
		}
	}
	fmt.Printf("Day 2 Part 1: %d\n", safe)
}
