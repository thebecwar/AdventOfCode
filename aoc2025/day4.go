package aoc2025

import (
	"advent/loader"
	"fmt"
)

func Day4Part1() {
	loader, err := loader.NewLoader("2025/day4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"..@@.@@@@.",
			"@@@.@.@.@@",
			"@@@@@.@.@@",
			"@.@@@@..@.",
			"@@.@@@@.@@",
			".@@@@@@@.@",
			".@.@.@.@@@",
			"@.@@@.@@@@",
			".@@@@@@@@.",
			"@.@.@@@.@.",
		}
	*/
	getContrib := func(grid []string, r, c int) int {
		if r < 0 || r >= len(grid) {
			return 0
		}
		if c < 0 || c >= len(grid[r]) {
			return 0
		}
		if grid[r][c] == '@' {
			return 1
		} else {
			return 0
		}
	}
	neighbors := func(grid []string, row, col int) int {
		count := 0
		for r := row - 1; r <= row+1; r++ {
			for c := col - 1; c <= col+1; c++ {
				if r == row && c == col {
					continue
				}
				count += getContrib(grid, r, c)
			}
		}
		return count
	}

	count := 0
	for row := 0; row < len(loader.Lines); row++ {
		for col := 0; col < len(loader.Lines[row]); col++ {
			if loader.Lines[row][col] == '.' {
				continue
			}
			if neighbors(loader.Lines, row, col) < 4 {
				count++
			}
		}
	}

	fmt.Printf("Day 4 Part 1: %d\n", count)
}

func Day4Part2() {
	loader, err := loader.NewLoader("2025/day4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"..@@.@@@@.",
			"@@@.@.@.@@",
			"@@@@@.@.@@",
			"@.@@@@..@.",
			"@@.@@@@.@@",
			".@@@@@@@.@",
			".@.@.@.@@@",
			"@.@@@.@@@@",
			".@@@@@@@@.",
			"@.@.@@@.@.",
		}
	*/
	getContrib := func(grid [][]byte, r, c int) int {
		if r < 0 || r >= len(grid) {
			return 0
		}
		if c < 0 || c >= len(grid[r]) {
			return 0
		}
		if grid[r][c] == '@' {
			return 1
		} else {
			return 0
		}
	}
	neighbors := func(grid [][]byte, row, col int) int {
		count := 0
		for r := row - 1; r <= row+1; r++ {
			for c := col - 1; c <= col+1; c++ {
				if r == row && c == col {
					continue
				}
				count += getContrib(grid, r, c)
			}
		}
		return count
	}
	type point struct {
		row, col int
	}
	findPositions := func(grid [][]byte) []point {
		points := []point{}
		for row := 0; row < len(grid); row++ {
			for col := 0; col < len(grid[row]); col++ {
				if grid[row][col] == '.' {
					continue
				}
				if neighbors(grid, row, col) < 4 {
					points = append(points, point{row: row, col: col})
				}
			}
		}
		return points
	}
	removePoints := func(grid [][]byte, points []point) {
		for _, pt := range points {
			grid[pt.row][pt.col] = '.'
		}
	}

	grid := [][]byte{}
	for _, r := range loader.Lines {
		grid = append(grid, []byte(r))
	}

	count := 0
	removed := true
	for removed {
		nextIter := findPositions(grid)
		count += len(nextIter)
		removePoints(grid, nextIter)
		removed = len(nextIter) > 0
	}

	fmt.Printf("Day 4 Part 2: %d\n", count)
}
