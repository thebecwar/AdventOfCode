package aoc2023

import (
	"advent/loader"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func computeFirstLastSum(lines []string) int {
	dual := regexp.MustCompile(`(\d).*(\d)`)
	single := regexp.MustCompile(`(\d)`)

	sum := 0
	for _, line := range lines {
		//fmt.Println(line)
		digits := dual.FindStringSubmatch(line)
		if len(digits) == 0 {
			digits = single.FindStringSubmatch(line)
			if len(digits) == 0 {
				fmt.Println("Invalid line: ", line)
				return -1
			}
		}

		var val int
		var err error
		if len(digits) == 2 {
			val, err = strconv.Atoi(fmt.Sprintf("%s%s", digits[1], digits[1]))
		} else {
			val, err = strconv.Atoi(fmt.Sprintf("%s%s", digits[1], digits[2]))
		}

		if err != nil {
			fmt.Println("Invalid line: ", line)
			return -1
		}
		sum += val
	}
	return sum
}

func Day1Part1() {
	loader, err := loader.NewLoader("2023/day1.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	sum := computeFirstLastSum(loader.Lines)
	fmt.Printf("Day 1 Part 1: %d\n", sum)
}

func Day1Part2() {
	loader, err := loader.NewLoader("2023/day1.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	regex := regexp.MustCompile(`zero|one|two|three|four|five|six|seven|eight|nine`)
	toNumber := func(s string) string {
		switch s {
		case "zero":
			return "z0o"
		case "one":
			return "o1e"
		case "two":
			return "t2o"
		case "three":
			return "t3e"
		case "four":
			return "f4r"
		case "five":
			return "f5e"
		case "six":
			return "s6x"
		case "seven":
			return "s7n"
		case "eight":
			return "e8t"
		case "nine":
			return "n9e"
		default:
			panic("Invalid number")
		}
	}

	for i, line := range loader.Lines {
		for {
			if !regex.MatchString(line) {
				break
			}
			firstMatch := regex.FindString(line)
			line = strings.Replace(line, firstMatch, toNumber(firstMatch), 1)
		}

		loader.Lines[i] = line
	}

	sum := computeFirstLastSum(loader.Lines)
	fmt.Printf("Day 1 Part 2: %d\n", sum)
}
