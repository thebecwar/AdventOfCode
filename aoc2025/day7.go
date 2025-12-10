package aoc2025

import (
	"advent/loader"
	"fmt"
	"strings"
)

func Day7Part1() {
	loader, err := loader.NewLoader("2025/day7.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			".......S.......",
			"...............",
			".......^.......",
			"...............",
			"......^.^......",
			"...............",
			".....^.^.^.....",
			"...............",
			"....^.^...^....",
			"...............",
			"...^.^...^.^...",
			"...............",
			"..^...^.....^..",
			"...............",
			".^.^.^.^.^...^.",
			"...............",
		}
	*/
	splits := 0
	beams := map[int]bool{
		strings.Index(loader.Lines[0], "S"): true,
	}
	for row := 1; row < len(loader.Lines); row++ {
		adds := []int{}
		drops := []int{}
		for col, _ := range beams {
			if loader.Lines[row][col] == '^' {
				splits++
				drops = append(drops, col)
				adds = append(adds, col-1)
				adds = append(adds, col+1)
			}
		}
		for _, col := range adds {
			beams[col] = true
		}
		for _, col := range drops {
			delete(beams, col)
		}
	}

	fmt.Printf("Day 7 Part 1: %d\n", splits)
}

func Day7Part2() {
	loader, err := loader.NewLoader("2025/day7.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			".......S.......",
			"...............",
			".......^.......",
			"...............",
			"......^.^......",
			"...............",
			".....^.^.^.....",
			"...............",
			"....^.^...^....",
			"...............",
			"...^.^...^.^...",
			"...............",
			"..^...^.....^..",
			"...............",
			".^.^.^.^.^...^.",
			"...............",
		}
	*/

	type visitedMemo map[int]map[int]int
	getMemo := func(m visitedMemo, r, c int) int {
		if row, ok := m[r]; ok {
			if col, ok := row[c]; ok {
				return col
			}
		}
		return -1
	}
	setMemo := func(m visitedMemo, r, c, count int) {
		if _, ok := m[r]; !ok {
			m[r] = map[int]int{}
		}
		m[r][c] = count
	}
	m := visitedMemo{}

	var dfs func(int, int) int
	dfs = func(row, inCol int) int {
		if row >= len(loader.Lines) {
			return 1
		}
		if val := getMemo(m, row, inCol); val >= 0 {
			return val
		}
		if loader.Lines[row][inCol] == '^' {
			count := dfs(row+2, inCol-1) + dfs(row+2, inCol+1)
			setMemo(m, row, inCol, count)
			return count
		} else {
			count := dfs(row+2, inCol)
			setMemo(m, row, inCol, count)
			return count
		}
	}
	col := strings.Index(loader.Lines[0], "S")
	res := dfs(0, col)

	fmt.Printf("Day 7 Part 2: %d\n", res)
}
