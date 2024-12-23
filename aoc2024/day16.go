package aoc2024

import (
	"advent/containers"
	"advent/loader"
	"container/heap"
	"fmt"
)

type Maze struct {
	containers.Grid[string]
	StartPoint   containers.Point
	EndPoint     containers.Point
	EmptyTile    string
	WallTile     string
	StartTile    string
	EndTile      string
	DirectionMap map[string]containers.Point
}

type MazePointWithDirection struct {
	Point     containers.Point
	Direction string
}

func NewMaze(lines []string) *Maze {
	result := &Maze{
		Grid:      *containers.NewStringGrid(lines),
		EmptyTile: ".",
		WallTile:  "#",
		StartTile: "S",
		EndTile:   "E",
		DirectionMap: map[string]containers.Point{
			"^": {X: 0, Y: -1},
			">": {X: 1, Y: 0},
			"v": {X: 0, Y: 1},
			"<": {X: -1, Y: 0},
		},
	}
	x, y := result.Find(func(value string) bool { return value == result.StartTile })
	result.StartPoint = containers.Point{X: x, Y: y}
	x, y = result.Find(func(value string) bool { return value == result.EndTile })
	result.EndPoint = containers.Point{X: x, Y: y}

	return result
}

func (m Maze) IsMoveTile(x, y int) bool {
	return *m.Get(x, y) == m.EmptyTile || *m.Get(x, y) == m.EndTile || *m.Get(x, y) == m.StartTile
}
func (m Maze) IsMoveTilePoint(point containers.Point) bool {
	return m.IsMoveTile(point.X, point.Y)
}

func (m Maze) GetAvailableMoves(x, y int) []MazePointWithDirection {
	result := []MazePointWithDirection{}
	for direction, point := range m.DirectionMap {
		if m.IsMoveTile(x+point.X, y+point.Y) {
			result = append(result, MazePointWithDirection{
				Point:     containers.Point{X: x + point.X, Y: y + point.Y},
				Direction: direction,
			})
		}
	}
	return result
}

func (m Maze) GetRotationCount(from, to string) int {
	directions := [4]string{"^", ">", "v", "<"}
	fromIndex := -1
	toIndex := -1
	for i, direction := range directions {
		if direction == from {
			fromIndex = i
		}
		if direction == to {
			toIndex = i
		}
	}
	distance := toIndex - fromIndex
	if distance < 0 {
		distance += 4
	}
	if distance == 3 {
		distance -= 2 // Account for wrapping around
	}
	return distance
}

func (m Maze) GetRotationCost(from, to string) int {
	return m.GetRotationCount(from, to) * 1000 // 1000 per 90 degrees
}
func (m Maze) CountAdjacent(x, y int) int {
	adjacent := 0
	if x > 0 && m.IsMoveTile(x-1, y) {
		adjacent++
	}
	if x < m.Width()-1 && m.IsMoveTile(x+1, y) {
		adjacent++
	}
	if y > 0 && m.IsMoveTile(x, y-1) {
		adjacent++
	}
	if y < m.Height()-1 && m.IsMoveTile(x, y+1) {
		adjacent++
	}
	return adjacent
}
func (m Maze) CountAdjacentPoint(point containers.Point) int {
	return m.CountAdjacent(point.X, point.Y)
}

func (m Maze) IsIntersection(x, y int) bool {
	return m.CountAdjacent(x, y) > 1
}
func (m Maze) IsIntersectionPoint(point containers.Point) bool {
	return m.IsIntersection(point.X, point.Y)
}

func (m Maze) FindDeadEnd() *containers.Point {
	for y := 0; y < m.Height(); y++ {
		for x := 0; x < m.Width(); x++ {
			if *m.Get(x, y) == m.EmptyTile {
				if (x == m.StartPoint.X && y == m.StartPoint.Y) || (x == m.EndPoint.X && y == m.EndPoint.Y) {
					// Start and end don't count as dead ends
					continue
				}
				if m.CountAdjacent(x, y) == 1 {
					return &containers.Point{X: x, Y: y}
				}
			}
		}
	}
	return nil
}
func (m Maze) FillDeadEnds() {
	// Not super efficient method, but it doesn't need any logic to walk
	for end := m.FindDeadEnd(); end != nil; end = m.FindDeadEnd() {
		m.Set(end.X, end.Y, &m.WallTile)
	}
}

type mazePqItem struct {
	value    MazePointWithDirection
	priority int
	index    int
}
type mazePriorityQueue []*mazePqItem

func (pq mazePriorityQueue) Len() int { return len(pq) }
func (pq mazePriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq mazePriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *mazePriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*mazePqItem)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *mazePriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
func (pq *mazePriorityQueue) update(item *mazePqItem, value MazePointWithDirection, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
func (pq *mazePriorityQueue) Contains(value MazePointWithDirection) bool {
	for _, item := range *pq {
		if item.value == value {
			return true
		}
	}
	return false
}
func (pq *mazePriorityQueue) ContainsPoint(value containers.Point) bool {
	for _, item := range *pq {
		if item.value.Point == value {
			return true
		}
	}
	return false
}

func (m Maze) aStarSearch() int {
	openSet := mazePriorityQueue{
		{
			value:    MazePointWithDirection{Point: m.StartPoint, Direction: ">"},
			priority: m.StartPoint.ManhattanDistance(&m.EndPoint),
		},
	}
	heap.Init(&openSet)

	gScore := make(map[containers.Point]int)
	gScore[m.StartPoint] = 0

	// For node n, fScore[n] := gScore[n] + h(n). fScore[n] represents our current best guess as to
	// how cheap a path could be from start to finish if it goes through n.
	fScore := make(map[containers.Point]int)
	fScore[m.StartPoint] = m.StartPoint.ManhattanDistance(&m.EndPoint) + 1000

	for len(openSet) > 0 {
		// This operation can occur in O(Log(N)) time if openSet is a min-heap or a priority queue
		current := heap.Pop(&openSet).(*mazePqItem)
		if current.value.Point == m.EndPoint {
			return gScore[current.value.Point]
		}

		for _, neighbor := range m.GetAvailableMoves(current.value.Point.X, current.value.Point.Y) {
			estimatedCost := gScore[current.value.Point] + m.GetRotationCost(current.value.Direction, neighbor.Direction) + 1
			if cost, ok := gScore[neighbor.Point]; !ok || estimatedCost < cost {
				gScore[neighbor.Point] = estimatedCost
				fScore[neighbor.Point] = estimatedCost + m.GetRotationCost(current.value.Direction, neighbor.Direction)

				if !openSet.Contains(neighbor) {
					heap.Push(&openSet, &mazePqItem{
						value:    neighbor,
						priority: fScore[neighbor.Point],
					})
				}
			}
		}

	}

	return -1
}

func (m Maze) FindMinCostPath() int {
	return m.aStarSearch()
}

func (m Maze) NextMinCostStep(
	minRemaining int,
	current MazePointWithDirection,
	parentStack *containers.Stack[containers.Point],
	paths *[][]containers.Point,
) {
	if current.Point == m.StartPoint && minRemaining == 0 {
		// Reconstruct path
		nextPath := []containers.Point{}
		for _, point := range parentStack.Items {
			nextPath = append(nextPath, *point)
		}
		*paths = append(*paths, nextPath)
	}
	if minRemaining == 0 {
		return
	}

	parentStack.Push(&current.Point)
	defer parentStack.Pop()
	for _, neighbor := range m.GetAvailableMoves(current.Point.X, current.Point.Y) {
		nextStepCost := minRemaining - 1
		if current.Direction != "" {
			nextStepCost -= m.GetRotationCost(current.Direction, neighbor.Direction)
		}
		if nextStepCost < 0 {
			continue
		}
		m.NextMinCostStep(nextStepCost, neighbor, parentStack, paths)
	}
}

func (m Maze) FindAllMinCostPathPoints(min int) []containers.Point {
	parents := containers.NewStack[containers.Point]()
	allPaths := [][]containers.Point{}
	m.NextMinCostStep(min, MazePointWithDirection{Point: m.EndPoint, Direction: ""}, parents, &allPaths)

	points := map[containers.Point]bool{}
	for _, path := range allPaths {
		for _, point := range path {
			points[point] = true
		}
	}
	result := []containers.Point{}
	for point := range points {
		result = append(result, point)
	}
	return result
}

func Day16Part1() {
	loader, err := loader.NewLoader("2024/day16.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 7036
	/*loader.Lines = []string{
		"###############",
		"#.......#....E#",
		"#.#.###.#.###.#",
		"#.....#.#...#.#",
		"#.###.#####.#.#",
		"#.#.#.......#.#",
		"#.#.#####.###.#",
		"#...........#.#",
		"###.#.#####.#.#",
		"#...#.....#.#.#",
		"#.#.#.###.#.#.#",
		"#.....#...#.#.#",
		"#.###.#.#.#.#.#",
		"#S..#.....#...#",
		"###############",
	}*/
	// 11048
	/*loader.Lines = []string{
		"#################",
		"#...#...#...#..E#",
		"#.#.#.#.#.#.#.#.#",
		"#.#.#.#...#...#.#",
		"#.#.#.#.###.#.#.#",
		"#...#.#.#.....#.#",
		"#.#.#.#.#.#####.#",
		"#.#...#.#.#.....#",
		"#.#.#####.#.###.#",
		"#.#.#.......#...#",
		"#.#.###.#####.###",
		"#.#.#...#.....#.#",
		"#.#.#.#####.###.#",
		"#.#.#.........#.#",
		"#.#.#.#########.#",
		"#S#.............#",
		"#################",
	}*/

	maze := NewMaze(loader.Lines)
	//maze.PrintGrid()
	maze.FillDeadEnds()
	//maze.PrintGrid()
	min := maze.FindMinCostPath()

	fmt.Printf("Day 16 Part 1: %d\n", min)
}

func Day16Part2() {
	loader, err := loader.NewLoader("2024/day16.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 7036
	loader.Lines = []string{
		"###############",
		"#.......#....E#",
		"#.#.###.#.###.#",
		"#.....#.#...#.#",
		"#.###.#####.#.#",
		"#.#.#.......#.#",
		"#.#.#####.###.#",
		"#...........#.#",
		"###.#.#####.#.#",
		"#...#.....#.#.#",
		"#.#.#.###.#.#.#",
		"#.....#...#.#.#",
		"#.###.#.#.#.#.#",
		"#S..#.....#...#",
		"###############",
	}
	// 11048
	/*loader.Lines = []string{
		"#################",
		"#...#...#...#..E#",
		"#.#.#.#.#.#.#.#.#",
		"#.#.#.#...#...#.#",
		"#.#.#.#.###.#.#.#",
		"#...#.#.#.....#.#",
		"#.#.#.#.#.#####.#",
		"#.#...#.#.#.....#",
		"#.#.#####.#.###.#",
		"#.#.#.......#...#",
		"#.#.###.#####.###",
		"#.#.#...#.....#.#",
		"#.#.#.#####.###.#",
		"#.#.#.........#.#",
		"#.#.#.#########.#",
		"#S#.............#",
		"#################",
	}*/

	maze := NewMaze(loader.Lines)
	//maze.PrintGrid()
	//maze.FillDeadEnds()
	//maze.PrintGrid()

	min := maze.FindMinCostPath()
	minPoints := maze.FindAllMinCostPathPoints(min)

	oh := "O"
	for _, m := range minPoints {
		maze.Set(m.X, m.Y, &oh)
	}
	maze.PrintGrid()

	fmt.Printf("Day 16 Part 2: %d\n", len(minPoints))

}
