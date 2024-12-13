package aoc2024

import (
	"advent/containers"
	"advent/loader"
	"fmt"
)

type ClawMachine struct {
	A      containers.Point
	B      containers.Point
	Prize  containers.Point
	Winner [2]int
}

func NewClawMachine(lineA, lineB, linePrize string) *ClawMachine {
	a := containers.Point{}
	b := containers.Point{}
	prize := containers.Point{}

	fmt.Sscanf(lineA, "Button A: X+%d, Y+%d", &a.X, &a.Y)
	fmt.Sscanf(lineB, "Button B: X+%d, Y+%d", &b.X, &b.Y)
	fmt.Sscanf(linePrize, "Prize: X=%d, Y=%d", &prize.X, &prize.Y)

	return &ClawMachine{
		A:     a,
		B:     b,
		Prize: prize,
	}
}
func (c *ClawMachine) String() string {
	return fmt.Sprintf("ClawMachine{A: %+v, B: %+v, Prize: %+v, Winner: %+v, Cost: %d}", c.A, c.B, c.Prize, c.Winner, c.Cost())
}
func (c *ClawMachine) Solve() (int, int) {
	// Solve the system of equations
	// a*AX + b*BX = PX
	// a*AY + b*BY = PY
	// Using cramer's rule

	// | AX BX |
	// | AY BY |
	det := c.A.X*c.B.Y - c.A.Y*c.B.X

	// | PX BX |
	// | PY BY |
	detX := c.Prize.X*c.B.Y - c.Prize.Y*c.B.X

	// | AX PX |
	// | AY PY |
	detY := c.A.X*c.Prize.Y - c.A.Y*c.Prize.X

	a := detX / det
	b := detY / det

	// check if we have a valid solution
	if a*c.A.X+b*c.B.X == c.Prize.X && a*c.A.Y+b*c.B.Y == c.Prize.Y {
		c.Winner = [2]int{a, b}
		return a, b
	}

	c.Winner = [2]int{0, 0}
	return 0, 0
}
func (c *ClawMachine) HasWinner() bool {
	return c.Winner[0] != 0 || c.Winner[1] != 0
}
func (c *ClawMachine) Cost() int {
	return c.Winner[0]*3 + c.Winner[1]
}

func Day13Part1() {
	loader, err := loader.NewLoader("2024/day13.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"Button A: X+94, Y+34",
		"Button B: X+22, Y+67",
		"Prize: X=8400, Y=5400",
		"",
		"Button A: X+26, Y+66",
		"Button B: X+67, Y+21",
		"Prize: X=12748, Y=12176",
		"",
		"Button A: X+17, Y+86",
		"Button B: X+84, Y+37",
		"Prize: X=7870, Y=6450",
		"",
		"Button A: X+69, Y+23",
		"Button B: X+27, Y+71",
		"Prize: X=18641, Y=10279",
		"",
		"Button A: X+1, Y+1", // Added case for >100 presses of a button to get to solution
		"Button B: X+1, Y+0",
		"Prize: X=10000, Y=10000",
	}*/

	clawMachines := []*ClawMachine{}
	for i := 0; i < len(loader.Lines)+1; i += 4 {
		clawMachines = append(clawMachines, NewClawMachine(loader.Lines[i], loader.Lines[i+1], loader.Lines[i+2]))
	}

	winners := 0
	totalCost := 0
	for _, clawMachine := range clawMachines {
		clawMachine.Solve()
		if clawMachine.HasWinner() {
			winners++
			totalCost += clawMachine.Cost()
		}
	}

	fmt.Printf("Day 13 Part 1: %d (%d prizes)\n", totalCost, winners)
}

func Day13Part2() {
	loader, err := loader.NewLoader("2024/day13.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"Button A: X+94, Y+34",
		"Button B: X+22, Y+67",
		"Prize: X=8400, Y=5400",
		"",
		"Button A: X+26, Y+66",
		"Button B: X+67, Y+21",
		"Prize: X=12748, Y=12176",
		"",
		"Button A: X+17, Y+86",
		"Button B: X+84, Y+37",
		"Prize: X=7870, Y=6450",
		"",
		"Button A: X+69, Y+23",
		"Button B: X+27, Y+71",
		"Prize: X=18641, Y=10279",
		"",
		"Button A: X+1, Y+1", // Added case for >100 presses of a button to get to solution
		"Button B: X+1, Y+0",
		"Prize: X=10000, Y=10000",
	}*/

	offset := 10000000000000
	clawMachines := []*ClawMachine{}
	for i := 0; i < len(loader.Lines)+1; i += 4 {
		machine := NewClawMachine(loader.Lines[i], loader.Lines[i+1], loader.Lines[i+2])
		machine.Prize.X += offset
		machine.Prize.Y += offset
		clawMachines = append(clawMachines, machine)
	}

	winners := 0
	totalCost := 0
	for _, clawMachine := range clawMachines {
		clawMachine.Solve()
		if clawMachine.HasWinner() {
			winners++
			totalCost += clawMachine.Cost()
		}
	}

	fmt.Printf("Day 13 Part 2: %d (%d prizes)\n", totalCost, winners)
}
