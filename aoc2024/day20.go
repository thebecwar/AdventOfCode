package aoc2024

import (
	"advent/containers"
	"advent/loader"
	"fmt"
)

type CPUCheat struct {
	Start containers.Point
	End   containers.Point
}

func (c *CPUCheat) ManhattanDistance() int {
	return c.Start.ManhattanDistance(&c.End)
}

func isPointVisitable(grid *containers.Grid[string], point containers.Point) bool {
	return *grid.Get(point.X, point.Y) == "." || *grid.Get(point.X, point.Y) == "S" || *grid.Get(point.X, point.Y) == "E"
}

type WallCrossFill struct {
	Location    containers.Point
	CrossedWall bool
}

func findPoints(grid *containers.Grid[string], startPoint containers.Point, maxDistance int) map[containers.Point]int {
	if startPoint.X < 0 || startPoint.X >= grid.Width() || startPoint.Y < 0 || startPoint.Y >= grid.Height() {
		return map[containers.Point]int{}
	}

	result := map[containers.Point]int{}

	visited := map[containers.Point]bool{}
	queue := containers.NewStack[WallCrossFill]()
	queue.Push(&WallCrossFill{Location: startPoint, CrossedWall: false})

	// For each direction:
	// If the point is more than Manhattan Distance from the start, skip it
	// If we've crossed a wall to get to the current point:
	//   If the next point is not a wall, this is a cheat point. We're done.
	//   If the next point is a wall, continue
	// If we haven't crossed a wall to get to the current point:
	//   If the next point is a wall, enqueue the next point and continue.
	//   If the next point is not a wall, skip it.
	handleMovement := func(current WallCrossFill, next containers.Point) {
		// OOB
		if next.X < 0 || next.X >= grid.Width() || next.Y < 0 || next.Y >= grid.Height() {
			return
		}
		if _, ok := visited[next]; !ok && next.ManhattanDistance(&startPoint) <= maxDistance {
			// Have we already crossed a wall?
			if current.CrossedWall {
				if isPointVisitable(grid, next) {
					result[next]++
					visited[next] = true
				} else {
					// We're still in the wall
					queue.Push(&WallCrossFill{Location: next, CrossedWall: true})
				}
			} else {
				if isPointVisitable(grid, next) {
					// We haven't crossed a wall, so this is not a cheat, it's just a path on the course
					visited[next] = true
				} else {
					// We're in a wall, so let's cross it
					queue.Push(&WallCrossFill{Location: next, CrossedWall: true})
				}
			}
		}
	}

	for !queue.IsEmpty() {
		current := *queue.Pop()
		visited[current.Location] = true

		handleMovement(current, containers.Point{X: current.Location.X, Y: current.Location.Y - 1})
		handleMovement(current, containers.Point{X: current.Location.X, Y: current.Location.Y + 1})
		handleMovement(current, containers.Point{X: current.Location.X - 1, Y: current.Location.Y})
		handleMovement(current, containers.Point{X: current.Location.X + 1, Y: current.Location.Y})
	}

	return result
}

func FloodFillFindPoint(
	grid *containers.Grid[string],
	maxLength, remaining int,
	currentPoint containers.Point,
	reachable *map[containers.Point]int,
	fromPoint containers.Point,
	crossedWall bool) {

	if remaining < 0 {
		return
	}
	if currentPoint.X < 0 || currentPoint.X >= grid.Width() || currentPoint.Y < 0 || currentPoint.Y >= grid.Height() {
		return // Out of bounds
	}

	visitable := isPointVisitable(grid, currentPoint)
	if visitable && !crossedWall && maxLength != remaining {
		return
	}
	if visitable && crossedWall {
		(*reachable)[currentPoint]++
		//return
	}

	if currentPoint.X-1 != fromPoint.X {
		FloodFillFindPoint(grid, maxLength, remaining-1, containers.Point{X: currentPoint.X - 1, Y: currentPoint.Y}, reachable, currentPoint, true)
	}
	if currentPoint.X+1 != fromPoint.X {
		FloodFillFindPoint(grid, maxLength, remaining-1, containers.Point{X: currentPoint.X + 1, Y: currentPoint.Y}, reachable, currentPoint, true)
	}
	if currentPoint.Y-1 != fromPoint.Y {
		FloodFillFindPoint(grid, maxLength, remaining-1, containers.Point{X: currentPoint.X, Y: currentPoint.Y - 1}, reachable, currentPoint, true)
	}
	if currentPoint.Y+1 != fromPoint.Y {
		FloodFillFindPoint(grid, maxLength, remaining-1, containers.Point{X: currentPoint.X, Y: currentPoint.Y + 1}, reachable, currentPoint, true)
	}
}

func findCpuCheats(grid *containers.Grid[string], maxLength int, startPoint containers.Point) map[CPUCheat]int {
	cheats := map[CPUCheat]int{}

	// for each spot on the path, find all empty points within a manhattan distance of maxLength
	for x := 0; x < grid.Width(); x++ {
		for y := 0; y < grid.Height(); y++ {
			if isPointVisitable(grid, containers.Point{X: x, Y: y}) {
				next := findPoints(grid, containers.Point{X: x, Y: y}, maxLength)
				for k, v := range next {
					cheats[CPUCheat{Start: containers.Point{X: x, Y: y}, End: k}] += v
				}
			}
		}
	}

	result := map[CPUCheat]int{}
	for k, v := range cheats {
		d := k.Start.ManhattanDistance(&k.End)
		if d == 1 {
			continue
		}
		result[k] = v
	}
	return result
}

func Day20Part1() {
	loader, err := loader.NewLoader("2024/day20.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	part1Threshold := 100
	part2Threshold := 100

	/*loader.Lines = []string{
		"###############",
		"#...#...#.....#",
		"#.#.#.#.#.###.#",
		"#S#...#.#.#...#",
		"#######.#.#.###",
		"#######.#.#...#",
		"#######.#.###.#",
		"###..E#...#...#",
		"###.#######.###",
		"#...###...#...#",
		"#.#####.#.###.#",
		"#.#...#.#.#...#",
		"#.#.#.#.#.#.###",
		"#...#...#...###",
		"###############",
	}
	part1Threshold = 1
	part2Threshold = 50*/

	grid := containers.NewStringGrid(loader.Lines)

	sX, sY := grid.Find(func(s string) bool { return s == "S" })
	eX, eY := grid.Find(func(s string) bool { return s == "E" })

	start := containers.Point{X: sX, Y: sY}
	end := containers.Point{X: eX, Y: eY}

	// Trace out the existing route to get the distance from the start to every point on the course
	costs := map[containers.Point]int{}
	current := start
	costs[current] = 0
	for current != end {
		// Find the neighbor
		up := containers.Point{X: current.X, Y: current.Y - 1}
		down := containers.Point{X: current.X, Y: current.Y + 1}
		left := containers.Point{X: current.X - 1, Y: current.Y}
		right := containers.Point{X: current.X + 1, Y: current.Y}
		if up.Y > 0 {
			if _, ok := costs[up]; !ok && isPointVisitable(grid, up) {
				costs[up] = costs[current] + 1
				current = up
				continue
			}
		}
		if down.Y < grid.Height()-1 {
			if _, ok := costs[down]; !ok && isPointVisitable(grid, down) {
				costs[down] = costs[current] + 1
				current = down
				continue
			}
		}
		if left.X > 0 {
			if _, ok := costs[left]; !ok && isPointVisitable(grid, left) {
				costs[left] = costs[current] + 1
				current = left
				continue
			}
		}
		if right.X < grid.Width()-1 {
			if _, ok := costs[right]; !ok && isPointVisitable(grid, right) {
				costs[right] = costs[current] + 1
				current = right
				continue
			}
		}
	}

	cpuCheats := findCpuCheats(grid, 2, start)
	cheats := map[int]int{} // cheats[savings]count
	for p, _ := range cpuCheats {
		d := p.ManhattanDistance()
		if costs[p.Start]+d < costs[p.End] {
			shortening := costs[p.End] - (costs[p.Start] + d)
			cheats[shortening]++
		}
	}
	routesInThreshold := 0
	for saving, count := range cheats {
		if saving >= part1Threshold {
			routesInThreshold += count
		}
	}
	fmt.Printf("Day 20 Part 1: %d\n", routesInThreshold)

	cheatCount := 0
	for p, startCost := range costs {
		for next, endCost := range costs {
			if p == next {
				continue
			}
			d := p.ManhattanDistance(&next)
			if d <= 20 && endCost-startCost-d >= part2Threshold {
				cheatCount++
			}
		}
	}
	fmt.Printf("Day 20 Part 2: %d\n", cheatCount)
}

func Day20Part2() {
	return
}
