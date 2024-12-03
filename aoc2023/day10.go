package aoc2023

import (
	"advent/containers"
	"advent/loader"
	"fmt"
)

type Connection struct {
	Top, Right, Bottom, Left bool
}
type GridCell struct {
	Value   string
	Count   int
	Visited bool
	Region  int
}
type Region struct {
	Id       int
	Location containers.Point
	Edge     bool
	Size     int
}

func (g *GridCell) Connections() Connection {
	conn, ok := connections[g.Value]
	if !ok {
		return Connection{}
	}
	return conn
}

var connections = map[string]Connection{
	"L": Connection{Top: true, Right: true, Bottom: false, Left: false},
	"J": Connection{Top: true, Right: false, Bottom: false, Left: true},
	"F": Connection{Top: false, Right: true, Bottom: true, Left: false},
	"7": Connection{Top: false, Right: false, Bottom: true, Left: true},
	"|": Connection{Top: true, Right: false, Bottom: true, Left: false},
	"-": Connection{Top: false, Right: true, Bottom: false, Left: true},
	"S": Connection{Top: true, Right: true, Bottom: true, Left: true},
}

func parseGrid(lines []string) *containers.Grid[GridCell] {
	rawGrid := containers.NewStringGrid(lines)
	finalGrid := containers.NewGrid[GridCell](rawGrid.Width(), rawGrid.Height())
	for y, row := range rawGrid.Cells {
		for x, cell := range row {
			finalGrid.Cells[y][x] = GridCell{Value: cell, Count: -1}
		}
	}
	return finalGrid
}

func getConnections(grid *containers.Grid[GridCell], x, y int) []containers.Point {
	result := []containers.Point{}

	current := grid.Get(x, y)
	conn := current.Connections()

	if conn.Top && y-1 >= 0 {
		above := grid.Get(x, y-1)
		if above != nil && above.Connections().Bottom {
			result = append(result, containers.Point{X: x, Y: y - 1})
		}
	}
	if conn.Right && x+1 < grid.Width() {
		right := grid.Get(x+1, y)
		if right != nil && right.Connections().Left {
			result = append(result, containers.Point{X: x + 1, Y: y})
		}
	}
	if conn.Bottom && y+1 < grid.Height() {
		below := grid.Get(x, y+1)
		if below != nil && below.Connections().Top {
			result = append(result, containers.Point{X: x, Y: y + 1})
		}
	}
	if conn.Left && x-1 >= 0 {
		left := grid.Get(x-1, y)
		if left != nil && left.Connections().Right {
			result = append(result, containers.Point{X: x - 1, Y: y})
		}
	}
	return result
}

func printGrid(grid *containers.Grid[GridCell]) {
	for _, row := range grid.Cells {
		for _, cell := range row {
			fmt.Printf("%s ", cell.Value)
		}
		fmt.Println()
	}
}

func traversePath(grid *containers.Grid[GridCell], startX, startY int) containers.Point {
	grid.Get(startX, startY).Count = 0

	maxX, maxY, max := 0, 0, 0

	s := containers.NewQueue[containers.Point]()
	s.Enqueue(containers.NewPoint(startX, startY))

	var current *containers.Point
	for !s.IsEmpty() {
		current = s.Dequeue()
		value := grid.Get(current.X, current.Y)
		//fmt.Printf("Visiting (%d,%d) %s visited: %t\n", current.X, current.Y, value.Value, value.Visited)
		value.Visited = true

		conns := getConnections(grid, current.X, current.Y)
		for _, conn := range conns {
			connected := grid.Get(conn.X, conn.Y)
			if connected.Visited {
				continue
			}
			connected.Count = value.Count + 1
			if connected.Count > max {
				max = connected.Count
				maxX = conn.X
				maxY = conn.Y
			}
			s.Enqueue(&conn)
		}
	}
	return containers.Point{X: maxX, Y: maxY}
}

func identifyRegions(grid *containers.Grid[GridCell]) []Region {
	result := []Region{}

	queue := containers.NewQueue[containers.Point]()

	currentRegion := 0
	regionName := "0"

	rx, ry := grid.Find(func(v GridCell) bool { return v.Value == "." })
	for rx != -1 && ry != -1 {
		regionData := Region{Id: currentRegion, Edge: false, Size: 0}
		queue.Enqueue(containers.NewPoint(rx, ry))
		fmt.Printf("Starting region at (%d,%d)\n", rx, ry)
		for !queue.IsEmpty() {
			current := queue.Dequeue()
			value := grid.Get(current.X, current.Y)
			if value.Visited {
				continue
			}

			value.Value = regionName
			value.Region = currentRegion
			value.Visited = true
			regionData.Edge = regionData.Edge || grid.IsEdge(current.X, current.Y)
			regionData.Size++

			// Above
			// Left
			if current.X-1 >= 0 && current.Y-1 >= 0 && !grid.Get(current.X-1, current.Y-1).Visited {
				queue.Enqueue(containers.NewPoint(current.X-1, current.Y-1))
			}
			// Center
			if current.Y-1 >= 0 && !grid.Get(current.X, current.Y-1).Visited {
				queue.Enqueue(containers.NewPoint(current.X, current.Y-1))
			}
			// Right
			if current.X+1 < grid.Width() && current.Y-1 >= 0 && grid.Get(current.X+1, current.Y-1).Visited {
				queue.Enqueue(containers.NewPoint(current.X+1, current.Y-1))
			}
			// L&R Neighbors
			if current.X-1 >= 0 && !grid.Get(current.X-1, current.Y).Visited {
				queue.Enqueue(containers.NewPoint(current.X-1, current.Y))
			}
			if current.X+1 < grid.Width() && !grid.Get(current.X+1, current.Y).Visited {
				queue.Enqueue(containers.NewPoint(current.X+1, current.Y))
			}
			// Below
			// Left
			if current.X-1 >= 0 && current.Y+1 < grid.Height() && !grid.Get(current.X-1, current.Y+1).Visited {
				queue.Enqueue(containers.NewPoint(current.X-1, current.Y+1))
			}
			// Center
			if current.Y+1 < grid.Height() && !grid.Get(current.X, current.Y+1).Visited {
				queue.Enqueue(containers.NewPoint(current.X, current.Y+1))
			}
			// Right
			if current.X+1 < grid.Width() && current.Y+1 < grid.Height() && !grid.Get(current.X+1, current.Y+1).Visited {
				queue.Enqueue(containers.NewPoint(current.X+1, current.Y+1))
			}

		}

		result = append(result, regionData)
		currentRegion++
		regionName = fmt.Sprintf("%d", currentRegion)
		rx, ry = grid.Find(func(v GridCell) bool { return v.Value == "." })
	}

	return result
}

func expandGrid(grid *containers.Grid[GridCell]) *containers.Grid[GridCell] {
	newWidth := grid.Width() * 2
	newHeight := grid.Height() * 2
	newGrid := containers.NewGrid[GridCell](newWidth, newHeight)

	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			newGrid.Cells[y*2][x*2] = *grid.Get(x, y)
			// Down
			if y+1 < grid.Height() {
				if grid.Get(x, y).Connections().Bottom && grid.Get(x, y+1).Connections().Top {
					newGrid.Cells[y*2+1][x*2] = *grid.Get(x, y+1)
				}
			}
			// Right
			if x+1 < grid.Width() {
				if grid.Get(x, y).Connections().Right && grid.Get(x+1, y).Connections().Left {
					newGrid.Cells[y*2][x*2+1] = *grid.Get(x+1, y)
				}
			}
		}
	}

	return newGrid
}

func Day10Part1() {
	loader, err := loader.NewLoader("2023-day10-part1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"7-F7-",
			".FJ|7",
			"SJLL7",
			"|F--J",
			"LJ.LJ",
		}
	*/
	// L top<->right, J top<->left, F bottom<->right, 7 bottom<->left, | bottom<->top, - left<->right
	// S start
	grid := parseGrid(loader.Lines)
	sX, sY := grid.Find(func(v GridCell) bool { return v.Value == "S" })
	maxPoint := traversePath(grid, sX, sY)
	max := grid.Get(maxPoint.X, maxPoint.Y).Count

	fmt.Printf("Day 10 Part 1: (%d, %d): %d\n", maxPoint.X, maxPoint.Y, max)
}

func Day10Part2() {
	loader, err := loader.NewLoader("2023-day10-part1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"...........",
		".S-------7.",
		".|F-----7|.",
		".||.....||.",
		".||.....||.",
		".|L-7.F-J|.",
		".|..|.|..|.",
		".L--J.L--J.",
		"...........",
	}*/
	/*loader.Lines = []string{
		".F----7F7F7F7F-7....",
		".|F--7||||||||FJ....",
		".||.FJ||||||||L7....",
		"FJL7L7LJLJ||LJ.L-7..",
		"L--J.L7...LJS7F-7L7.",
		"....F-J..F7FJ|L7L7L7",
		"....L7.F7||L7|.L7L7|",
		".....|FJLJ|FJ|F7|.LJ",
		"....FJL-7.||.||||...",
		"....L---J.LJ.LJLJ...",
	}*/
	/*loader.Lines = []string{
		"FF7FSF7F7F7F7F7F---7",
		"L|LJ||||||||||||F--J",
		"FL-7LJLJ||||||LJL-77",
		"F--JF--7||LJLJ7F7FJ-",
		"L---JF-JLJ.||-FJLJJ7",
		"|F|F-JF---7F7-L7L|7|",
		"|FFJF7L7F-JF7|JL---7",
		"7-L-JL7||F7|L7F-7F7|",
		"L.L7LFJ|||||FJL7||LJ",
		"L7JLJL-JLJLJL--JLJ.L",
	}*/
	// L top<->right, J top<->left, F bottom<->right, 7 bottom<->left, | bottom<->top, - left<->right
	// S start
	grid := parseGrid(loader.Lines)
	sX, sY := grid.Find(func(v GridCell) bool { return v.Value == "S" })
	_ = traversePath(grid, sX, sY)

	expandedGrid := expandGrid(grid)
	_ = traversePath(expandedGrid, sX*2, sY*2)

	for y, row := range grid.Cells {
		for x, cell := range row {
			if cell.Visited {
				grid.Cells[y][x].Value = "X"
			} else {
				grid.Cells[y][x].Value = "."
			}
			cell.Visited = false
		}
	}
	for y, row := range expandedGrid.Cells {
		for x, cell := range row {
			if cell.Visited {
				expandedGrid.Cells[y][x].Value = "X"
			} else {
				expandedGrid.Cells[y][x].Value = "."
			}
			cell.Visited = false
		}
	}

	regions := identifyRegions(grid)
	expandedRegions := identifyRegions(expandedGrid)
	regionLookup := make(map[int]Region)
	for _, region := range expandedRegions {
		regionLookup[region.Id] = region
	}

	//printGrid(grid)

	external := 0
	internal := 0
	for y, row := range grid.Cells {
		for x := range row {
			if grid.Cells[y][x].Value == "X" {
				continue
			}
			expandedRegionId := expandedGrid.Get(x*2, y*2).Region
			region := regionLookup[expandedRegionId]
			if region.Edge {
				grid.Cells[y][x].Value = "O"
				external++
			} else {
				grid.Cells[y][x].Value = "I"
				internal++
			}
		}
	}
	//printGrid(grid)

	fmt.Printf(
		"Day 10 Part 2: Enclosed Area: %d (regions: %d, external: %d, internal: %d)\n",
		internal,
		len(regions),
		external,
		internal,
	)
}
