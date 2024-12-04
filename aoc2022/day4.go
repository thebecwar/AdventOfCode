package aoc2022

import (
	"advent/loader"
	"fmt"
)

func parseElfPairs(line string) (elf1 [2]int, elf2 [2]int) {
	fmt.Sscanf(line, "%d-%d,%d-%d", &elf1[0], &elf1[1], &elf2[0], &elf2[1])
	return
}

func Day4Part1() {
	loader, err := loader.NewLoader("2022/day4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}*/

	containSets := 0
	for _, line := range loader.Lines {
		elf1, elf2 := parseElfPairs(line)

		if elf1[0] <= elf2[0] && elf1[1] >= elf2[1] {
			containSets++
		} else if elf2[0] <= elf1[0] && elf2[1] >= elf1[1] {
			containSets++
		}
	}
	fmt.Printf("Day 4, Part 1: %d\n", containSets)
}
func Day4Part2() {
	loader, err := loader.NewLoader("2022/day4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}*/

	overlaps := 0
	for _, line := range loader.Lines {
		elf1, elf2 := parseElfPairs(line)

		if elf1[0] <= elf2[1] && elf2[0] <= elf1[1] {
			overlaps++
		}
	}
	fmt.Printf("Day 4, Part 2: %d\n", overlaps)
}
