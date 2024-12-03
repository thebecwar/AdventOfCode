package aoc2023

import (
	"advent/containers"
	"advent/loader"
	"fmt"
)

func parseAndExpandUniverse(lines []string, factor int) []*containers.Point {
	rawGalaxies := []*containers.Point{}
	columns := make([]bool, len(lines[0]))
	emptyRows := []int{}

	for y, row := range lines {
		galaxyRow := false
		for x, cell := range row {
			if cell == '#' {
				columns[x] = true
				galaxyRow = true
				rawGalaxies = append(rawGalaxies, containers.NewPoint(x, y))
			}
		}
		if !galaxyRow {
			emptyRows = append(emptyRows, y)
		}
	}

	columnMap := make(map[int]int)
	rowMap := make(map[int]int)

	current := -1
	for x := range columns {
		if !columns[x] {
			current += factor
		} else {
			current++
		}
		columnMap[x] = current
	}
	idx := 0
	current = -1
	for y := 0; y < len(lines); y++ {
		if idx < len(emptyRows) && emptyRows[idx] == y {
			current += factor
			idx++
		} else {
			current++
		}
		rowMap[y] = current
	}

	expandedGalaxies := []*containers.Point{}
	for _, galaxy := range rawGalaxies {
		expandedGalaxies = append(expandedGalaxies, containers.NewPoint(columnMap[galaxy.X], rowMap[galaxy.Y]))
	}
	return expandedGalaxies
}

func Day11Part1() {
	loader, err := loader.NewLoader("2023-day11-part1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	/*loader.Lines = []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}*/

	galaxies := parseAndExpandUniverse(loader.Lines, 2)

	totalDistance := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			totalDistance += galaxies[i].ManhattanDistance(galaxies[j])
		}
	}

	fmt.Printf("Day 11, Part 1: Total distance: %d\n", totalDistance)

}
func Day11Part2() {
	loader, err := loader.NewLoader("2023-day11-part1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	/*loader.Lines = []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}*/

	galaxies := parseAndExpandUniverse(loader.Lines, 1000000)

	totalDistance := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			totalDistance += galaxies[i].ManhattanDistance(galaxies[j])
		}
	}

	fmt.Printf("Day 11, Part 1: Total distance: %d\n", totalDistance)
}
