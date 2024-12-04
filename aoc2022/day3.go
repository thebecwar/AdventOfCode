package aoc2022

import (
	"advent/loader"
	"fmt"
)

func Day3Part1() {
	loader, err := loader.NewLoader("2022/day3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}*/

	sum := 0
	for _, line := range loader.Lines {
		firstHalf := line[:len(line)/2]
		secondHalf := line[len(line)/2:]
		first := make(map[rune]bool)
		for _, r := range firstHalf {
			first[r] = true
		}
		for _, r := range secondHalf {
			_, ok := first[r]
			if ok {
				if r >= 'A' && r <= 'Z' {
					value := int(r) - int('A') + 27
					sum += value
					break
				} else {
					value := int(r) - int('a') + 1
					sum += value
					break
				}
			}
		}
	}
	fmt.Printf("Day 3, Part 1: %d\n", sum)
}
func GetCharMap(line string) map[rune]bool {
	charMap := make(map[rune]bool)
	for _, r := range line {
		charMap[r] = true
	}
	return charMap
}
func Day3Part2() {
	loader, err := loader.NewLoader("2022/day3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}*/

	sum := 0
	for i := 0; i < len(loader.Lines); i += 3 {
		elf0 := GetCharMap(loader.Lines[i])
		elf1 := GetCharMap(loader.Lines[i+1])

		for _, r := range loader.Lines[i+2] {
			_, ok0 := elf0[r]
			_, ok1 := elf1[r]
			if ok0 && ok1 {
				if r >= 'A' && r <= 'Z' {
					value := int(r) - int('A') + 27
					sum += value
					break
				} else {
					value := int(r) - int('a') + 1
					sum += value
					break
				}
			}
		}
	}
	fmt.Printf("Day 3, Part 2: %d\n", sum)
}
