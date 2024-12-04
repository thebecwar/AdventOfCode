package aoc2024

import (
	"advent/loader"
	"fmt"
	"math"
	"sort"
)

func parseLists(lines []string) (left []int, right []int) {
	for _, line := range lines {
		var l, r int
		fmt.Sscanf(line, "%d %d", &l, &r)
		left = append(left, l)
		right = append(right, r)
	}
	return
}

func Day1Part1() {
	loader, err := loader.NewLoader("2024/day1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}*/

	left, right := parseLists(loader.Lines)
	sort.Ints(left)
	sort.Ints(right)

	sum := 0

	for i := 0; i < len(left); i++ {
		diff := int(math.Abs(float64(left[i] - right[i])))
		sum += diff
	}

	fmt.Printf("Day 1 Part 1: %d\n", sum)
}
func Day1Part2() {
	loader, err := loader.NewLoader("2024/day1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}*/

	left, right := parseLists(loader.Lines)

	frequencies := make(map[int]int)
	for i := 0; i < len(right); i++ {
		frequencies[right[i]]++
	}

	sum := 0
	for i := 0; i < len(left); i++ {
		score := left[i] * frequencies[left[i]]
		sum += score
	}

	fmt.Printf("Day 1 Part 2: %d\n", sum)

}
