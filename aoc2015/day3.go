package aoc2015

import (
	"advent/loader"
	"fmt"
)

var directions = map[rune][2]int{
	'^': {0, 1},
	'v': {0, -1},
	'>': {1, 0},
	'<': {-1, 0},
}

func Day3Part1() {
	loader, err := loader.NewLoader("2015/day3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	loc := [2]int{0, 0}
	visited := make(map[string]bool)
	visited["0,0"] = true

	for _, r := range loader.Lines[0] {
		dir := directions[r]
		loc[0] += dir[0]
		loc[1] += dir[1]
		visited[fmt.Sprintf("%d,%d", loc[0], loc[1])] = true
	}

	fmt.Printf("Day 3 Part 1: %d\n", len(visited))
}

func Day3Part2() {
	loader, err := loader.NewLoader("2015/day3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	loc := [2]int{0, 0}
	rsLoc := [2]int{0, 0}
	visited := make(map[string]bool)
	visited["0,0"] = true

	for i := 0; i < len(loader.Lines); i++ {
		for j := 0; j < len(loader.Lines[i])-1; j += 2 {
			santa := directions[rune(loader.Lines[i][j])]
			rs := directions[rune(loader.Lines[i][j+1])]
			loc[0] += santa[0]
			loc[1] += santa[1]
			visited[fmt.Sprintf("%d,%d", loc[0], loc[1])] = true
			rsLoc[0] += rs[0]
			rsLoc[1] += rs[1]
			visited[fmt.Sprintf("%d,%d", rsLoc[0], rsLoc[1])] = true
		}
	}

	fmt.Printf("Day 3 Part 2: %d\n", len(visited))
}
