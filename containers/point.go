package containers

import "math"

type Point struct {
	X, Y int
}

func NewPoint(x, y int) *Point {
	return &Point{X: x, Y: y}
}
func (p *Point) Add(other *Point) *Point {
	return NewPoint(p.X+other.X, p.Y+other.Y)
}
func (p *Point) Subtract(other *Point) *Point {
	return NewPoint(p.X-other.X, p.Y-other.Y)
}
func (p *Point) Multiply(scalar int) *Point {
	return NewPoint(p.X*scalar, p.Y*scalar)
}
func (p *Point) Divide(scalar int) *Point {
	return NewPoint(p.X/scalar, p.Y/scalar)
}
func (p *Point) Equals(other *Point) bool {
	return p.X == other.X && p.Y == other.Y
}
func (p *Point) Copy() *Point {
	return NewPoint(p.X, p.Y)
}
func (p *Point) ManhattanDistance(other *Point) int {
	return int(math.Abs(float64(p.X-other.X)) + math.Abs(float64(p.Y-other.Y)))
}
