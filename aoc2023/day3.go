package aoc2023

import (
	"advent/loader"
	"fmt"
	"regexp"
	"strconv"
)

func isNumeric(c byte) bool {
	if c == '0' {
		return true
	} else if c == '1' {
		return true
	} else if c == '2' {
		return true
	} else if c == '3' {
		return true
	} else if c == '4' {
		return true
	} else if c == '5' {
		return true
	} else if c == '6' {
		return true
	} else if c == '7' {
		return true
	} else if c == '8' {
		return true
	} else if c == '9' {
		return true
	} else {
		return false
	}
}
func isSymbol(c byte) bool {
	if c == '.' {
		return false
	} else if c == '0' {
		return false
	} else if c == '1' {
		return false
	} else if c == '2' {
		return false
	} else if c == '3' {
		return false
	} else if c == '4' {
		return false
	} else if c == '5' {
		return false
	} else if c == '6' {
		return false
	} else if c == '7' {
		return false
	} else if c == '8' {
		return false
	} else if c == '9' {
		return false
	} else {
		return true
	}
}

func Day3Part1() {
	loader, err := loader.NewLoader("2023/day3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"467..114..",
			"...*......",
			"..35..633.",
			"......#...",
			"617*......",
			".....+.58.",
			"..592.....",
			"......755.",
			"...$.*....",
			".664.598..",
		}
	*/

	digitRegex := regexp.MustCompile(`(\d+)`)

	sum := 0

	for i, line := range loader.Lines {
		numbers := digitRegex.FindAllStringSubmatchIndex(line, -1)
		for _, number := range numbers {
			keep := false
			for x := number[0] - 1; x < number[1]+1; x++ {
				for y := i - 1; y <= i+1; y++ {
					if x < 0 || y < 0 || x >= len(line) || y >= len(loader.Lines) {
						continue
					}
					if isSymbol(loader.Lines[y][x]) {
						keep = true
						break
					}
				}
				if keep {
					break
				}
			}
			if keep {
				n, _ := strconv.Atoi(line[number[0]:number[1]])
				sum += n
			}
		}
	}

	fmt.Printf("Day 3 Part 1: %d\n", sum)

}

type NumberPosition struct {
	Left  int
	Right int // Last index inclusive
	Y     int
	Value int
}
type StarPosition struct {
	X int
	Y int
}
type Grid struct {
	Numbers []NumberPosition
	Stars   []StarPosition
}

func ParseGrid(lines []string) *Grid {
	grid := &Grid{}

	digitRegex := regexp.MustCompile(`(\d+)`)
	starRegex := regexp.MustCompile(`\*`)

	for i, line := range lines {
		numbers := digitRegex.FindAllStringSubmatchIndex(line, -1)
		for _, number := range numbers {
			value, _ := strconv.Atoi(line[number[0]:number[1]])
			grid.Numbers = append(grid.Numbers, NumberPosition{Left: number[0], Right: number[1] - 1, Y: i, Value: value})
		}
		stars := starRegex.FindAllStringSubmatchIndex(line, -1)
		for _, star := range stars {
			grid.Stars = append(grid.Stars, StarPosition{X: star[0], Y: i})
		}
	}

	return grid
}

func (n *NumberPosition) Contains(x int, y int) bool {
	if y == n.Y {
		if x >= n.Left && x <= n.Right {
			return true
		}
	}
	return false
}
func BetweenInclusive(n int, a int, b int) bool {
	return n >= a && n <= b
}
func (n *NumberPosition) RectContains(left int, right int, top int, bottom int) bool {
	if n.Y >= top && n.Y <= bottom {
		if BetweenInclusive(n.Left, left, right) {
			return true
		}
		if BetweenInclusive(n.Right, left, right) {
			return true
		}
		if n.Left < left && n.Right > right {
			return true
		}
	}
	return false
}

func Day3Part2() {
	loader, err := loader.NewLoader("2023/day3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"467..114..",
			"...*......",
			"..35..633.",
			"......#...",
			"617*......",
			".....+.58.",
			"..592.....",
			"......755.",
			"...$.*....",
			".664.598..",
		}
	*/

	grid := ParseGrid(loader.Lines)

	gearRatios := []int{}

	for _, star := range grid.Stars {
		gear := []int{}
		for _, number := range grid.Numbers {
			if number.RectContains(star.X-1, star.X+1, star.Y-1, star.Y+1) {
				gear = append(gear, number.Value)
			}
		}
		if len(gear) == 2 {
			// we have a gear
			gearRatios = append(gearRatios, gear[0]*gear[1])
		}
	}

	sum := 0
	for _, gear := range gearRatios {
		sum += gear
	}
	fmt.Printf("Day 3 Part 2: %d\n", sum)

}
