package aoc2015

import (
	"advent/containers"
	"advent/loader"
	"fmt"
	"math"
	"strings"
)

type LightInstruction struct {
	instruction string
	x1          int
	y1          int
	x2          int
	y2          int
}

var instructionSetPart1 = map[string]func(int) int{
	"turn on":  func(_ int) int { return 1 },
	"turn off": func(_ int) int { return 0 },
	"toggle":   func(i int) int { return 1 - i },
}

var instructionSetPart2 = map[string]func(int) int{
	"turn on":  func(i int) int { return i + 1 },
	"turn off": func(i int) int { return int(math.Max(float64(i-1), 0)) },
	"toggle":   func(i int) int { return i + 2 },
}

func parseLightInstruction(instruction string) *LightInstruction {
	result := &LightInstruction{}
	if strings.HasPrefix(instruction, "turn on") {
		result.instruction = "turn on"
	} else if strings.HasPrefix(instruction, "turn off") {
		result.instruction = "turn off"
	} else if strings.HasPrefix(instruction, "toggle") {
		result.instruction = "toggle"
	}
	rest := strings.Replace(instruction, result.instruction, "", 1)
	fmt.Sscanf(rest, "%d,%d through %d,%d", &result.x1, &result.y1, &result.x2, &result.y2)

	return result
}

func Day6Part1() {
	loader, err := loader.NewLoader("2015/day6.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	grid := containers.NewGrid[int](1000, 1000)

	for _, line := range loader.Lines {
		instruction := parseLightInstruction(line)
		for x := instruction.x1; x <= instruction.x2; x++ {
			for y := instruction.y1; y <= instruction.y2; y++ {
				light := grid.Get(x, y)
				*light = instructionSetPart1[instruction.instruction](*light)
			}
		}
	}

	lightsOn := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if *grid.Get(x, y) == 1 {
				lightsOn++
			}
		}
	}

	fmt.Printf("Day 6 Part 1: %d\n", lightsOn)
}

func Day6Part2() {
	loader, err := loader.NewLoader("2015/day6.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	grid := containers.NewGrid[int](1000, 1000)

	for _, line := range loader.Lines {
		instruction := parseLightInstruction(line)
		for x := instruction.x1; x <= instruction.x2; x++ {
			for y := instruction.y1; y <= instruction.y2; y++ {
				light := grid.Get(x, y)
				*light = instructionSetPart2[instruction.instruction](*light)
			}
		}
	}

	brightness := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			brightness += *grid.Get(x, y)
		}
	}

	fmt.Printf("Day 6 Part 2: %d\n", brightness)
}
