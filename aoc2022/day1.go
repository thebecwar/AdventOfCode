package aoc2022

import (
	"advent/loader"
	"fmt"
	"sort"
	"strconv"
)

func Day1Part1() {
	loader, err := loader.NewLoader("2022/day1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	/*loader.Lines = []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}*/

	max := 0
	current := 0

	for _, line := range loader.Lines {
		if line == "" {
			if current > max {
				max = current
			}
			current = 0
		}
		val, _ := strconv.Atoi(line)
		current += val
	}
	fmt.Printf("Day 1 Part 1: %d\n", max)

}
func Day1Part2() {
	loader, err := loader.NewLoader("2022/day1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	/*loader.Lines = []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}*/

	calories := []int{}

	current := 0
	for _, line := range loader.Lines {
		if line == "" {
			calories = append(calories, current)
			current = 0
		}
		val, _ := strconv.Atoi(line)
		current += val
	}
	if current > 0 {
		calories = append(calories, current)
	}

	sort.Ints(calories)

	topThree := calories[len(calories)-3:]
	total := topThree[0] + topThree[1] + topThree[2]

	fmt.Printf("Day 1 Part 2: %d (%+v)\n", total, topThree)
}
