package aoc2024

import (
	"advent/containers"
	"advent/loader"
	"fmt"
)

type SecurityRobot struct {
	Position containers.Point
	Velocity containers.Point
}

func SecurityScore(robots *[]SecurityRobot, height, width int) int {
	halfWidth := width / 2
	halfHeight := height / 2
	quadrants := [4]int{0, 0, 0, 0}

	for _, r := range *robots {
		if r.Position.X < halfWidth && r.Position.Y < halfHeight {
			// top left
			quadrants[0]++
		} else if r.Position.X < halfWidth && r.Position.Y > halfHeight {
			// bottom left
			quadrants[1]++
		} else if r.Position.X > halfWidth && r.Position.Y < halfHeight {
			// top right
			quadrants[2]++
		} else if r.Position.X > halfWidth && r.Position.Y > halfHeight {
			// bottom right
			quadrants[3]++
		}
	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func parseRobot(line string) SecurityRobot {
	result := SecurityRobot{}
	fmt.Sscanf(line, "p=%d,%d v=%d,%d", &result.Position.X, &result.Position.Y, &result.Velocity.X, &result.Velocity.Y)
	return result
}

func Day14Part1() {
	loader, err := loader.NewLoader("2024/day14.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	width := 101
	height := 103

	/*loader.Lines = []string{
		"p=0,4 v=3,-3",
		"p=6,3 v=-1,-3",
		"p=10,3 v=-1,2",
		"p=2,0 v=2,-1",
		"p=0,0 v=1,3",
		"p=3,0 v=-2,-2",
		"p=7,6 v=-1,-3",
		"p=3,0 v=-1,-2",
		"p=9,3 v=2,3",
		"p=7,3 v=-1,2",
		"p=2,4 v=2,-3",
		"p=9,5 v=-3,-3",
	}
	width := 11
	height := 7*/

	quadrants := [4]int{0, 0, 0, 0}
	halfWidth := width / 2
	halfHeight := height / 2

	for _, line := range loader.Lines {
		robot := parseRobot(line)
		x := (robot.Position.X + (robot.Velocity.X * 100)) % width
		y := (robot.Position.Y + (robot.Velocity.Y * 100)) % height

		for y > height {
			y -= height
		}
		for y < 0 {
			y += height
		}
		for x > width {
			x -= width
		}
		for x < 0 {
			x += width
		}

		if x < halfWidth && y < halfHeight {
			// top left
			quadrants[0]++
		} else if x < halfWidth && y > halfHeight {
			// bottom left
			quadrants[1]++
		} else if x > halfWidth && y < halfHeight {
			// top right
			quadrants[2]++
		} else if x > halfWidth && y > halfHeight {
			// bottom right
			quadrants[3]++
		}
	}

	fmt.Printf("Day 14 Part 1: %d, (%+v)\n", quadrants[0]*quadrants[1]*quadrants[2]*quadrants[3], quadrants)
}

// ...#... - Center X, 0
// ..#.#..
// .#...#.
// #.....#

func nextStep(robot *SecurityRobot, width, height int) {
	robot.Position.X += robot.Velocity.X
	if robot.Position.X < 0 {
		robot.Position.X += width
	}
	if robot.Position.X >= width {
		robot.Position.X -= width
	}
	robot.Position.Y += robot.Velocity.Y
	if robot.Position.Y < 0 {
		robot.Position.Y += height
	}
	if robot.Position.Y >= height {
		robot.Position.Y -= height
	}
}

func Day14Part2() {
	loader, err := loader.NewLoader("2024/day14.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	width := 101
	height := 103

	/*loader.Lines = []string{
		"p=0,4 v=3,-3",
		"p=6,3 v=-1,-3",
		"p=10,3 v=-1,2",
		"p=2,0 v=2,-1",
		"p=0,0 v=1,3",
		"p=3,0 v=-2,-2",
		"p=7,6 v=-1,-3",
		"p=3,0 v=-1,-2",
		"p=9,3 v=2,3",
		"p=7,3 v=-1,2",
		"p=2,4 v=2,-3",
		"p=9,5 v=-3,-3",
	}
	width = 11
	height = 7*/

	robots := []SecurityRobot{}
	for _, line := range loader.Lines {
		robot := parseRobot(line)
		robots = append(robots, robot)
	}

	minScore := 10000000000000
	minIteration := 0
	var minGrid *containers.Grid[bool]
	t := true

	for i := 0; i < 10000; i++ {
		locations := map[int]map[int]bool{}
		currentScore := 0
		for j := 0; j < len(robots); j++ {
			nextStep(&robots[j], width, height)
			if _, ok := locations[robots[j].Position.Y]; !ok {
				locations[robots[j].Position.Y] = map[int]bool{}
			}
			locations[robots[j].Position.Y][robots[j].Position.X] = true
		}
		currentScore = SecurityScore(&robots, height, width)
		if currentScore < minScore {
			minIteration = i
			minScore = currentScore
			minGrid = containers.NewGrid[bool](width, height)
			for y, row := range locations {
				for x := range row {
					minGrid.Set(x, y, &t)
				}
			}
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if minGrid.Get(x, y) != nil && *minGrid.Get(x, y) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Printf("Day 14 Part 2: %d\n", minIteration+1)

}
