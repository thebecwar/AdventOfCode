package aoc2025

import (
	"advent/loader"
	"fmt"
	"strconv"
	"strings"
)

func Day6Part1() {
	loader, err := loader.NewLoader("2025/day6.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"123 328  51 64 ",
			" 45 64  387 23 ",
			"  6 98  215 314",
			"*   +   *   +  ",
		}
	*/
	parseNumbersLine := func(line string) []int {
		nums := []int{}
		items := strings.Split(line, " ")
		for _, item := range items {
			if item == "" {
				continue
			}
			num, err := strconv.Atoi(item)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			nums = append(nums, num)
		}
		return nums
	}
	parseOperatorsLine := func(line string) []rune {
		ops := []rune{}
		items := strings.Split(line, " ")
		for _, item := range items {
			if item == "" {
				continue
			}
			op := []rune(item)[0]
			ops = append(ops, op)
		}
		return ops
	}
	parseGrid := func(lines []string) ([][]int, []rune) {
		numbers := [][]int{}
		for i := 0; i < len(lines)-1; i++ {
			nums := parseNumbersLine(lines[i])
			numbers = append(numbers, nums)
		}
		operators := parseOperatorsLine(lines[len(lines)-1])
		return numbers, operators
	}

	numbers, operators := parseGrid(loader.Lines)

	sum := 0
	for col := 0; col < len(numbers[0]); col++ {
		total := 0
		if operators[col] == '*' {
			total = 1
			for row := 0; row < len(numbers); row++ {
				total *= numbers[row][col]
			}
		} else if operators[col] == '+' {
			for row := 0; row < len(numbers); row++ {
				total += numbers[row][col]
			}
		}
		sum += total
	}

	fmt.Printf("Day 6 Part 1: %d\n", sum)
}

func Day6Part2() {
	loader, err := loader.NewLoader("2025/day6.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"123 328  51 64 ",
			" 45 64  387 23 ",
			"  6 98  215 314",
			"*   +   *   +  ",
		}
	*/
	type ColumnOperation struct {
		Numbers  []int
		Operator rune
	}
	parseGrid := func(lines []string) []ColumnOperation {
		operations := []ColumnOperation{}
		currentOp := ColumnOperation{}
		for col := len(lines[0]) - 1; col >= 0; col-- {
			num := 0
			multiplier := 1
			for row := len(lines) - 2; row >= 0; row-- {
				if lines[row][col] == ' ' {
					continue
				}
				digit := int(lines[row][col] - '0')
				num += digit * multiplier
				multiplier *= 10
			}
			currentOp.Numbers = append(currentOp.Numbers, num)
			if lines[len(lines)-1][col] != ' ' {
				// Operator, we're done with the column
				currentOp.Operator = rune(lines[len(lines)-1][col])
				operations = append(operations, currentOp)
				currentOp = ColumnOperation{}
				col--
			}
		}
		return operations
	}

	grid := parseGrid(loader.Lines)

	sum := 0
	for _, op := range grid {
		current := op.Numbers[0]
		for i := 1; i < len(op.Numbers); i++ {
			switch op.Operator {
			case '+':
				current += op.Numbers[i]
			case '*':
				current *= op.Numbers[i]
			}
		}
		sum += current
	}

	fmt.Printf("Day 6 Part 2: %d\n", sum)
}
