package aoc2024

import (
	"advent/containers"
	"advent/loader"
	"fmt"
	"sync"
)

var directions = []containers.Point{
	{X: 0, Y: -1}, // Up
	{X: 1, Y: 0},  // Right
	{X: 0, Y: 1},  // Down
	{X: -1, Y: 0}, // Left
}
var dirMarker = []string{"^", ">", "v", "<"}

func Day6Part1() {
	loader, err := loader.NewLoader("2024/day6.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}*/

	grid := containers.NewStringGrid(loader.Lines)

	cX, cY := grid.Find(func(x string) bool { return x == "^" })
	if cX == -1 || cY == -1 {
		fmt.Println("No starting point found")
		return
	}

	doneMarker := "X"
	direction := 0 // up

	for cX >= 0 && cX < grid.Width() && cY >= 0 && cY < grid.Height() {
		grid.Set(cX, cY, &doneMarker)
		nX := cX + directions[direction%4].X
		nY := cY + directions[direction%4].Y
		if nX < 0 || nX >= grid.Width() || nY < 0 || nY >= grid.Height() {
			// Exit Point
			//fmt.Printf("Exit at %d, %d\n", cX, cY)
			break
		}
		cell := grid.Get(nX, nY)
		if *cell == "#" {
			direction++ // Turn right
		} else {
			cX = nX
			cY = nY
		}
	}

	visitCount := 0
	for _, row := range grid.Cells {
		for _, cell := range row {
			if cell == doneMarker {
				visitCount++
			}
		}
	}
	//grid.PrintGrid()
	fmt.Printf("Day 6, Part 1: %d\n", visitCount)

}

func hasCycle(grid *containers.Grid[string], startX, startY int) bool {
	direction := 0
	visited := make(map[string]map[string]bool) // ["x,y"][direction]
	cX, cY := startX, startY
	for cX >= 0 && cX < grid.Width() && cY >= 0 && cY < grid.Height() {
		current := fmt.Sprintf("%d,%d", cX, cY)
		if _, ok := visited[current]; !ok {
			visited[current] = make(map[string]bool)
		}
		if _, ok := visited[current][dirMarker[direction%4]]; ok {
			return true
		}
		visited[current][dirMarker[direction%4]] = true

		nX := cX + directions[direction%4].X
		nY := cY + directions[direction%4].Y
		if nX < 0 || nX >= grid.Width() || nY < 0 || nY >= grid.Height() {
			// Exit Point
			//fmt.Printf("Exit at %d, %d\n", cX, cY)
			break
		}
		cell := grid.Get(nX, nY)
		if *cell == "#" {
			direction++ // Turn right
		} else {
			cX = nX
			cY = nY
		}
	}
	return false
}

func Day6Part2() {
	loader, err := loader.NewLoader("2024/day6.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}*/

	grid := containers.NewStringGrid(loader.Lines)

	startX, startY := grid.Find(func(x string) bool { return x == "^" })
	if startX == -1 || startY == -1 {
		fmt.Println("No starting point found")
		return
	}

	direction := 0
	visited := make(map[string]map[string]bool) // ["x,y"][direction]
	cX, cY := startX, startY
	for cX >= 0 && cX < grid.Width() && cY >= 0 && cY < grid.Height() {
		current := fmt.Sprintf("%d,%d", cX, cY)
		if _, ok := visited[current]; !ok {
			visited[current] = make(map[string]bool)
		}
		visited[current][dirMarker[direction%4]] = true

		nX := cX + directions[direction%4].X
		nY := cY + directions[direction%4].Y
		if nX < 0 || nX >= grid.Width() || nY < 0 || nY >= grid.Height() {
			// Exit Point
			//fmt.Printf("Exit at %d, %d\n", cX, cY)
			break
		}
		cell := grid.Get(nX, nY)
		if *cell == "#" {
			direction++ // Turn right
		} else {
			cX = nX
			cY = nY
		}
	}

	// Call the brute Squad.... I am the brute squad
	hash := "#"
	wg := &sync.WaitGroup{}
	i := 0
	results := make([]bool, len(visited))
	for k := range visited {
		x, y := 0, 0
		fmt.Sscanf(k, "%d,%d", &x, &y)
		wg.Add(1)
		go func(x, y, i int) {
			defer wg.Done()
			myCopy := grid.Copy()
			myCopy.Set(x, y, &hash)
			results[i] = hasCycle(myCopy, startX, startY)
		}(x, y, i)
		i++
	}

	wg.Wait()

	cycleOptions := 0
	for _, r := range results {
		if r {
			cycleOptions++
		}
	}

	fmt.Printf("Day 6, Part 2: %d\n", cycleOptions)

}
