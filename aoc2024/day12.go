package aoc2024

import (
	"advent/containers"
	"advent/loader"
	"fmt"
)

type PlantedPlot struct {
	Plant   rune
	Visited bool
}
type PlantedRegion struct {
	Plant     rune
	Area      int
	Perimeter int
	Corners   int
}

func NewPlantedRegion(plant rune) *PlantedRegion {
	return &PlantedRegion{Plant: plant, Area: 0, Perimeter: 0}
}

func makePlantGrid(lines []string) *containers.Grid[PlantedPlot] {
	grid := containers.NewGrid[PlantedPlot](len(lines), len(lines[0]))
	for y, line := range lines {
		for x, r := range line {
			grid.Set(x, y, &PlantedPlot{Plant: rune(r), Visited: false})
		}
	}
	return grid
}
func identifyPlantedRegion(grid *containers.Grid[PlantedPlot], x, y int, region *PlantedRegion) {
	if x < 0 || y < 0 || x >= grid.Width() || y >= grid.Height() {
		return
	}
	if grid.Get(x, y).Visited {
		return
	}
	if grid.Get(x, y).Plant != region.Plant {
		return
	}
	grid.Get(x, y).Visited = true
	region.Area++

	perimeter := 4
	fenceUp, fenceDown, fenceLeft, fenceRight := false, false, false, false
	fenceUL, fenceUR, fenceDL, fenceDR, fenceLU, fenceRU, fenceLD, fenceRD := false, false, false, false, false, false, false, false
	// Left
	if x > 0 && grid.Get(x-1, y).Plant == region.Plant {
		perimeter--
		if y > 0 && grid.Get(x-1, y-1).Plant != region.Plant {
			fenceLU = true
		}
		if y < grid.Height()-1 && grid.Get(x-1, y+1).Plant != region.Plant {
			fenceLD = true
		}
	} else {
		fenceLeft = true
	}
	// Right
	if x < grid.Width()-1 && grid.Get(x+1, y).Plant == region.Plant {
		perimeter--
		if y > 0 && grid.Get(x+1, y-1).Plant != region.Plant {
			fenceRU = true
		}
		if y < grid.Height()-1 && grid.Get(x+1, y+1).Plant != region.Plant {
			fenceRD = true
		}
	} else {
		fenceRight = true
	}
	// Up
	if y > 0 && grid.Get(x, y-1).Plant == region.Plant {
		perimeter--
		if x > 0 && grid.Get(x-1, y-1).Plant != region.Plant {
			fenceUL = true
		}
		if x < grid.Width()-1 && grid.Get(x+1, y-1).Plant != region.Plant {
			fenceUR = true
		}
	} else {
		fenceUp = true
	}
	// Down
	if y < grid.Height()-1 && grid.Get(x, y+1).Plant == region.Plant {
		perimeter--
		if x > 0 && grid.Get(x-1, y+1).Plant != region.Plant {
			fenceDL = true
		}
		if x < grid.Width()-1 && grid.Get(x+1, y+1).Plant != region.Plant {
			fenceDR = true
		}
	} else {
		fenceDown = true
	}
	region.Perimeter += perimeter

	// Exterior corner count
	// +---+
	// | A |
	// +---+
	if fenceUp && fenceRight {
		region.Corners++
	}
	if fenceRight && fenceDown {
		region.Corners++
	}
	if fenceDown && fenceLeft {
		region.Corners++
	}
	if fenceLeft && fenceUp {
		region.Corners++
	}
	// Interior corner count
	//  1 | A | 2
	// ---+   +---
	//  A   A   A
	// ---+   +---
	//  4 | A | 3
	if fenceUL && fenceLU { // 1
		region.Corners++
	}
	if fenceUR && fenceRU { // 2
		region.Corners++
	}
	if fenceDR && fenceRD { // 3
		region.Corners++
	}
	if fenceDL && fenceLD { // 4
		region.Corners++
	}

	identifyPlantedRegion(grid, x-1, y, region)
	identifyPlantedRegion(grid, x+1, y, region)
	identifyPlantedRegion(grid, x, y-1, region)
	identifyPlantedRegion(grid, x, y+1, region)
}

func Day12Part1() {
	loader, err := loader.NewLoader("2024/day12.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"RRRRIICCFF",
		"RRRRIICCCF",
		"VVRRRCCFFF",
		"VVRCCCJFFF",
		"VVVVCJJCFE",
		"VVIVCCJJEE",
		"VVIIICJJEE",
		"MIIIIIJJEE",
		"MIIISIJEEE",
		"MMMISSJEEE",
	}*/
	/*loader.Lines = []string{
		"AAAA",
		"BBCD",
		"BBCC",
		"EEEC",
	}*/
	/*loader.Lines = []string{
		"EEEEE",
		"EXXXX",
		"EEEEE",
		"EXXXX",
		"EEEEE",
	}*/
	/*loader.Lines = []string{
		"AAAAAA",
		"AAABBA",
		"AAABBA",
		"ABBAAA",
		"ABBAAA",
		"AAAAAA",
	}*/

	regions := make([]PlantedRegion, 0)

	grid := makePlantGrid(loader.Lines)
	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			if grid.Get(x, y).Visited {
				continue
			}
			region := NewPlantedRegion(grid.Get(x, y).Plant)
			identifyPlantedRegion(grid, x, y, region)
			regions = append(regions, *region)
		}
	}

	totalValue := 0
	totalValueP2 := 0
	for _, region := range regions {
		totalValue += region.Area * region.Perimeter
		totalValueP2 += region.Area * region.Corners
	}

	fmt.Printf("Day 12 Part 1: %d\n", totalValue)
	fmt.Printf("Day 12 Part 2: %d\n", totalValueP2)
}

func Day12Part2() {
	return
}
