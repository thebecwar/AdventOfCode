package aoc2024

import (
	"advent/loader"
	"fmt"
)

func xmasSearch(lines []string) int {
	xmasCount := 0
	for ri, row := range lines {
		for ci, col := range row {
			if col == 'X' {
				if ri >= 3 && lines[ri-1][ci] == 'M' && lines[ri-2][ci] == 'A' && lines[ri-3][ci] == 'S' {
					// Up
					xmasCount++
				}
				if ri >= 3 && ci <= len(row)-4 && lines[ri-1][ci+1] == 'M' && lines[ri-2][ci+2] == 'A' && lines[ri-3][ci+3] == 'S' {
					// Up Right
					xmasCount++
				}
				if ci <= len(row)-4 && lines[ri][ci+1] == 'M' && lines[ri][ci+2] == 'A' && lines[ri][ci+3] == 'S' {
					// Right
					xmasCount++
				}
				if ri <= len(lines)-4 && ci <= len(row)-4 && lines[ri+1][ci+1] == 'M' && lines[ri+2][ci+2] == 'A' && lines[ri+3][ci+3] == 'S' {
					// Down Right
					xmasCount++
				}
				if ri <= len(lines)-4 && lines[ri+1][ci] == 'M' && lines[ri+2][ci] == 'A' && lines[ri+3][ci] == 'S' {
					// Down
					xmasCount++
				}
				if ri <= len(lines)-4 && ci >= 3 && lines[ri+1][ci-1] == 'M' && lines[ri+2][ci-2] == 'A' && lines[ri+3][ci-3] == 'S' {
					// Down Left
					xmasCount++
				}
				if ci >= 3 && lines[ri][ci-1] == 'M' && lines[ri][ci-2] == 'A' && lines[ri][ci-3] == 'S' {
					// Left
					xmasCount++
				}
				if ri >= 3 && ci >= 3 && lines[ri-1][ci-1] == 'M' && lines[ri-2][ci-2] == 'A' && lines[ri-3][ci-3] == 'S' {
					// Up Left
					xmasCount++
				}
			}
		}
	}
	return xmasCount
}
func match(lines []string, r, c int) bool {
	if r < 0 || r >= len(lines) || c < 0 || c+3 >= len(lines[0]) {
		return false
	}
	/*
		X  0  1  2       X -2 -1  0
		0  M  .  S       0  S  .  M
		1  .  A  .       1  .  A  .
		2  M  .  S       2  S  .  M

		X  0  1  2       X -2 -1  0
		0  M  .  M      -2  S  .  S
		1  .  A  .      -1  .  A  .
		2  S  .  S       0  M  .  M

	*/

	raw := []rune(lines[r][c : c+3])
	raw[1] = '.'
	line1 := string(raw)

	raw = []rune(lines[r+1][c : c+3])
	raw[0] = '.'
	raw[2] = '.'
	line2 := string(raw)

	raw = []rune(lines[r+2][c : c+3])
	raw[1] = '.'
	line3 := string(raw)

	if line1 == "M.S" &&
		line2 == ".A." &&
		line3 == "M.S" {
		return true
	}
	if line1 == "M.M" &&
		line2 == ".A." &&
		line3 == "S.S" {
		return true
	}
	if line1 == "S.M" &&
		line2 == ".A." &&
		line3 == "S.M" {
		return true
	}
	if line1 == "S.S" &&
		line2 == ".A." &&
		line3 == "M.M" {
		return true
	}
	return false
}

func xmasSearchLol(lines []string) int {
	xmasCount := 0
	for r, row := range lines {
		for c, _ := range row {
			if r+2 < len(lines) && c+2 < len(row) {
				if match(lines, r, c) {
					xmasCount++
				}
			}
		}
	}
	return xmasCount
}

func Day4Part1() {
	loader, err := loader.NewLoader("2024/day4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}*/

	xmasCount := xmasSearch(loader.Lines)
	fmt.Printf("Day 4, Part 1: %d\n", xmasCount)

}
func Day4Part2() {
	loader, err := loader.NewLoader("2024/day4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}*/

	for x := range loader.Lines {
		loader.Lines[x] = loader.Lines[x] + "..."
	}
	emptyRow := []rune{}
	for range loader.Lines[0] {
		emptyRow = append(emptyRow, '.')
	}
	loader.Lines = append(loader.Lines, string(emptyRow))
	loader.Lines = append(loader.Lines, string(emptyRow))
	loader.Lines = append(loader.Lines, string(emptyRow))

	xmasCount := xmasSearchLol(loader.Lines)
	fmt.Printf("Day 4, Part 1: %d\n", xmasCount)
}
