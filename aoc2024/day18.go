package aoc2024

import (
	"advent/containers"
	"advent/loader"
	"container/heap"
	"fmt"
)

func parseByteRain(lines []string) []containers.Point {
	result := []containers.Point{}
	for _, line := range lines {
		x, y := 0, 0
		fmt.Sscanf(line, "%d,%d", &x, &y)
		result = append(result, containers.Point{X: x, Y: y})
	}
	return result
}
func adjacentMemoryPoints(point containers.Point, grid *containers.Grid[string]) []containers.Point {
	result := []containers.Point{}

	// up, down, left, right
	if point.Y > 0 {
		if *grid.Get(point.X, point.Y-1) != "#" {
			result = append(result, containers.Point{X: point.X, Y: point.Y - 1})
		}
	}
	if point.Y < grid.Height()-1 {
		if *grid.Get(point.X, point.Y+1) != "#" {
			result = append(result, containers.Point{X: point.X, Y: point.Y + 1})
		}
	}
	if point.X > 0 {
		if *grid.Get(point.X-1, point.Y) != "#" {
			result = append(result, containers.Point{X: point.X - 1, Y: point.Y})
		}
	}
	if point.X < grid.Width()-1 {
		if *grid.Get(point.X+1, point.Y) != "#" {
			result = append(result, containers.Point{X: point.X + 1, Y: point.Y})
		}
	}

	return result
}
func aStarMemory(start, end containers.Point, grid *containers.Grid[string]) int {
	// Borrow maze priority queue from day 16
	openSet := mazePriorityQueue{
		{
			value:    MazePointWithDirection{Point: start, Direction: ""},
			priority: start.ManhattanDistance(&end),
		},
	}
	heap.Init(&openSet)

	gScore := make(map[containers.Point]int)
	gScore[start] = 0

	fScore := make(map[containers.Point]int)
	fScore[start] = start.ManhattanDistance(&end)

	for len(openSet) > 0 {
		current := heap.Pop(&openSet).(*mazePqItem)
		if current.value.Point == end {
			return gScore[current.value.Point]
		}

		for _, neighbor := range adjacentMemoryPoints(current.value.Point, grid) {
			estimatedCost := gScore[current.value.Point] + 1
			if cost, ok := gScore[neighbor]; !ok || estimatedCost < cost {
				gScore[neighbor] = estimatedCost
				fScore[neighbor] = estimatedCost + neighbor.ManhattanDistance(&end)

				if !openSet.ContainsPoint(neighbor) {
					heap.Push(&openSet, &mazePqItem{
						value:    MazePointWithDirection{Point: neighbor, Direction: ""},
						priority: fScore[neighbor],
					})
				}
			}
		}

	}

	return -1
}

func Day18Part1() {
	loader, err := loader.NewLoader("2024/day18.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	width := 71
	height := 71
	limit := 1024

	/*loader.Lines = []string{
		"5,4",
		"4,2",
		"4,5",
		"3,0",
		"2,1",
		"6,3",
		"2,4",
		"1,5",
		"0,6",
		"3,3",
		"2,6",
		"5,1",
		"1,2",
		"5,5",
		"2,5",
		"6,5",
		"1,4",
		"0,4",
		"6,4",
		"1,1",
		"6,1",
		"1,0",
		"0,5",
		"1,6",
		"2,0",
	}
	width = 7
	height = 7
	limit = 12*/

	raindrops := parseByteRain(loader.Lines)
	corrupted := "#"
	free := "."

	memory := containers.NewGrid[string](width, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			memory.Set(x, y, &free)
		}
	}
	for i := 0; i < limit; i++ {
		memory.Set(raindrops[i].X, raindrops[i].Y, &corrupted)
	}
	minPath := aStarMemory(containers.Point{X: 0, Y: 0}, containers.Point{X: width - 1, Y: height - 1}, memory)

	fmt.Printf("Day 18 Part 1: %d\n", minPath)
}

func Day18Part2() {
	loader, err := loader.NewLoader("2024/day18.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	width := 71
	height := 71
	limit := 1024

	/*loader.Lines = []string{
		"5,4",
		"4,2",
		"4,5",
		"3,0",
		"2,1",
		"6,3",
		"2,4",
		"1,5",
		"0,6",
		"3,3",
		"2,6",
		"5,1",
		"1,2",
		"5,5",
		"2,5",
		"6,5",
		"1,4",
		"0,4",
		"6,4",
		"1,1",
		"6,1",
		"1,0",
		"0,5",
		"1,6",
		"2,0",
	}
	width = 7
	height = 7
	limit = 12*/

	raindrops := parseByteRain(loader.Lines)
	corrupted := "#"
	free := "."

	memory := containers.NewGrid[string](width, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			memory.Set(x, y, &free)
		}
	}
	for i := 0; i < limit; i++ {
		memory.Set(raindrops[i].X, raindrops[i].Y, &corrupted)
	}

	for i := limit; i < len(raindrops); i++ {
		memory.Set(raindrops[i].X, raindrops[i].Y, &corrupted)
		minPath := aStarMemory(containers.Point{X: 0, Y: 0}, containers.Point{X: width - 1, Y: height - 1}, memory)
		if minPath == -1 {
			fmt.Printf("Day 18 Part 2: %d,%d (i: %d)\n", raindrops[i].X, raindrops[i].Y, i)
			return
		}
	}

	fmt.Printf("Day 18 Part 1: Not Found")

}
