package containers

import "fmt"

type Grid[T any] struct {
	Cells [][]T
}

func NewGrid[T any](width, height int) *Grid[T] {
	cells := make([][]T, height)
	for i := range cells {
		cells[i] = make([]T, width)
	}
	return &Grid[T]{Cells: cells}
}
func NewStringGrid(lines []string) *Grid[string] {
	grid := NewGrid[string](len(lines[0]), len(lines))
	for y, row := range lines {
		for x, cell := range row {
			grid.Cells[y][x] = string(cell)
		}
	}
	return grid
}

func (g *Grid[T]) Get(x, y int) *T {
	return &g.Cells[y][x]
}
func (g *Grid[T]) Set(x, y int, value *T) {
	g.Cells[y][x] = *value
}
func (g *Grid[T]) Swap(x1, y1, x2, y2 int) {
	g.Cells[y1][x1], g.Cells[y2][x2] = g.Cells[y2][x2], g.Cells[y1][x1]
}
func (g *Grid[T]) Width() int {
	return len(g.Cells[0])
}
func (g *Grid[T]) Height() int {
	return len(g.Cells)
}
func (g *Grid[T]) IsEdge(x, y int) bool {
	return x == 0 || y == 0 || x == g.Width()-1 || y == g.Height()-1
}
func (g *Grid[T]) Find(predicate func(v T) bool) (int, int) {
	for y, row := range g.Cells {
		for x, cell := range row {
			if predicate(cell) {
				return x, y
			}
		}
	}
	return -1, -1
}
func (g *Grid[T]) PrintGrid() {
	for _, row := range g.Cells {
		for _, cell := range row {
			fmt.Printf("%v", cell)
		}
		fmt.Println()
	}
}
func (g *Grid[T]) ToString() string {
	result := ""
	for _, row := range g.Cells {
		for _, cell := range row {
			result += fmt.Sprintf("%v", cell)
		}
		result += "\n"
	}
	return result
}
func (g *Grid[T]) Copy() *Grid[T] {
	next := NewGrid[T](g.Width(), g.Height())
	for y, row := range g.Cells {
		copy(next.Cells[y], row)
	}
	return next
}
