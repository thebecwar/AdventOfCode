package aoc2024

import (
	"advent/containers"
	"advent/loader"
	"fmt"
	"strings"
)

func parseTopoGrid(lines []string) *containers.Grid[int] {
	grid := containers.NewGrid[int](len(lines), len(lines[0]))
	for y, line := range lines {
		for x, char := range line {
			value := int(char - '0')
			grid.Set(x, y, &value)
		}
	}
	return grid
}

func mergeMaps(a, b *map[string]int) {
	for k, v := range *b {
		(*a)[k] += v
	}
}

func getTrailEndpoint(s string) containers.Point {
	parts := strings.Split(s, "->")
	last := parts[len(parts)-1]
	x, y := 0, 0
	fmt.Sscanf(last, "(%d,%d)", &x, &y)
	return containers.Point{X: x, Y: y}
}
func getTrailEndpoints(paths map[string]int) map[containers.Point]int {
	endpoints := map[containers.Point]int{}
	for path := range paths {
		endpoints[getTrailEndpoint(path)]++
	}
	return endpoints
}

func findHikingPath(grid *containers.Grid[int], startPoint containers.Point, currentPath *containers.Stack[containers.Point]) map[string]int {
	startingValue := *grid.Get(startPoint.X, startPoint.Y)
	if startingValue == 9 {
		pathString := ""
		for _, point := range currentPath.Items {
			pathString += fmt.Sprintf("(%d,%d)->", point.X, point.Y)
		}
		pathString += fmt.Sprintf("(%d,%d)", startPoint.X, startPoint.Y)
		return map[string]int{pathString: 1}
	}

	paths := map[string]int{}
	currentPath.Push(&startPoint)
	// up
	if startPoint.Y > 0 {
		if *grid.Get(startPoint.X, startPoint.Y-1) == startingValue+1 {
			next := findHikingPath(grid, containers.Point{X: startPoint.X, Y: startPoint.Y - 1}, currentPath)
			mergeMaps(&paths, &next)
		}
	}
	// down
	if startPoint.Y < grid.Height()-1 {
		if *grid.Get(startPoint.X, startPoint.Y+1) == startingValue+1 {
			next := findHikingPath(grid, containers.Point{X: startPoint.X, Y: startPoint.Y + 1}, currentPath)
			mergeMaps(&paths, &next)
		}
	}
	// left
	if startPoint.X > 0 {
		if *grid.Get(startPoint.X-1, startPoint.Y) == startingValue+1 {
			next := findHikingPath(grid, containers.Point{X: startPoint.X - 1, Y: startPoint.Y}, currentPath)
			mergeMaps(&paths, &next)
		}
	}
	// right
	if startPoint.X < grid.Width()-1 {
		if *grid.Get(startPoint.X+1, startPoint.Y) == startingValue+1 {
			next := findHikingPath(grid, containers.Point{X: startPoint.X + 1, Y: startPoint.Y}, currentPath)
			mergeMaps(&paths, &next)
		}
	}
	currentPath.Pop()

	return paths
}

func Day10Part1() {
	loader, err := loader.NewLoader("2024/day10.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}*/

	trailheads := 0
	deadendTrailheads := 0
	endpoints := 0
	paths := map[string]int{}
	grid := parseTopoGrid(loader.Lines)
	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			if *grid.Get(x, y) == 0 {
				pathStack := containers.NewStack[containers.Point]()
				trailheads++
				next := findHikingPath(grid, containers.Point{X: x, Y: y}, pathStack)
				if len(next) == 0 {
					deadendTrailheads++
				}
				mergeMaps(&paths, &next)
				eps := getTrailEndpoints(next)
				endpoints += len(eps)
			}
		}
	}

	fmt.Printf("Day 10 Part 1: %d (paths: %d, trailheads: %d, deadends: %d)\n", endpoints, len(paths), trailheads, deadendTrailheads)
}

func Day10Part2() {
	// Exact same code, but we output the number of paths
	loader, err := loader.NewLoader("2024/day10.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}*/

	trailheads := 0
	deadendTrailheads := 0
	endpoints := 0
	paths := map[string]int{}
	grid := parseTopoGrid(loader.Lines)
	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			if *grid.Get(x, y) == 0 {
				pathStack := containers.NewStack[containers.Point]()
				trailheads++
				next := findHikingPath(grid, containers.Point{X: x, Y: y}, pathStack)
				if len(next) == 0 {
					deadendTrailheads++
				}
				mergeMaps(&paths, &next)
				eps := getTrailEndpoints(next)
				endpoints += len(eps)
			}
		}
	}

	fmt.Printf("Day 10 Part 1: %d (endpoints: %d, trailheads: %d, deadends: %d)\n", len(paths), endpoints, trailheads, deadendTrailheads)
}
