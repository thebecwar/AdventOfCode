package aoc2024

import (
	"advent/loader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func checkLinePossible(line string) (bool, int) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return false, 0
	}

	target, err := strconv.Atoi(parts[0])
	if err != nil {
		return false, 0
	}

	temp := strings.Fields(parts[1])
	numbers := make([]int, len(temp))
	for i := 0; i < len(temp); i++ {
		numbers[i], err = strconv.Atoi(temp[i])
		if err != nil {
			return false, 0
		}
	}

	gaps := len(numbers) - 1
	max := int(math.Pow(2, float64(gaps)))
	for i := 0; i < max; i++ {
		current := numbers[0]
		n := i
		for j := 1; j < len(numbers); j++ {
			if n%2 == 0 {
				current += numbers[j]
			} else {
				current *= numbers[j]
			}
			n = n / 2
		}
		if current == target {
			return true, current
		}
	}

	return false, 0
}

func checkLinePossibleWithConcat(line string) (bool, int) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return false, 0
	}

	target, err := strconv.Atoi(parts[0])
	if err != nil {
		return false, 0
	}

	temp := strings.Fields(parts[1])
	numbers := make([]int, len(temp))
	for i := 0; i < len(temp); i++ {
		numbers[i], err = strconv.Atoi(temp[i])
		if err != nil {
			return false, 0
		}
	}

	gaps := len(numbers) - 1
	max := int(math.Pow(3, float64(gaps)))
	for i := 0; i < max; i++ {
		current := numbers[0]
		n := i
		for j := 1; j < len(numbers); j++ {
			if n%3 == 0 {
				current += numbers[j]
			} else if n%3 == 1 {
				// Concat number
				length := math.Floor(math.Log10(float64(numbers[j]))) + 1
				current = current*int(math.Pow(10, float64(length))) + numbers[j]
			} else {
				current *= numbers[j]
			}
			n = n / 3
		}
		if current == target {
			return true, current
		}
	}

	return false, 0
}

func Day7Part1() {
	loader, err := loader.NewLoader("2024/day7.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	/*loader.Lines = []string{
		"190: 10 19",
		"3267: 81 40 27",
		"83: 17 5",
		"156: 15 6",
		"7290: 6 8 6 15",
		"161011: 16 10 13",
		"192: 17 8 14",
		"21037: 9 7 18 13",
		"292: 11 6 16 20",
	}*/

	total := 0
	for _, line := range loader.Lines {
		possible, value := checkLinePossible(line)
		if possible {
			total += value
		}
	}

	fmt.Printf("Day 7 Part 1: %d\n", total)

}
func Day7Part2() {
	loader, err := loader.NewLoader("2024/day7.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	/*loader.Lines = []string{
		"190: 10 19",
		"3267: 81 40 27",
		"83: 17 5",
		"156: 15 6",
		"7290: 6 8 6 15",
		"161011: 16 10 13",
		"192: 17 8 14",
		"21037: 9 7 18 13",
		"292: 11 6 16 20",
	}*/

	total := 0
	for _, line := range loader.Lines {
		possible, value := checkLinePossibleWithConcat(line)
		if possible {
			total += value
		}
	}

	fmt.Printf("Day 7 Part 2: %d\n", total)
}
