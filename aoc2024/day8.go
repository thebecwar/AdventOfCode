package aoc2024

import (
	"advent/containers"
	"advent/loader"
	"fmt"
)

func Day8Part1() {
	loader, err := loader.NewLoader("2024/day8.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}*/

	antennas := map[rune][]containers.Point{}
	for y, line := range loader.Lines {
		for x, c := range line {
			if c != '.' {
				antennas[c] = append(antennas[c], containers.Point{x, y})
			}
		}
	}

	nodes := map[containers.Point]bool{}
	for _, v := range antennas {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				dx := v[i].X - v[j].X
				dy := v[i].Y - v[j].Y
				nodes[containers.Point{X: v[i].X + dx, Y: v[i].Y + dy}] = true
				nodes[containers.Point{X: v[j].X - dx, Y: v[j].Y - dy}] = true
			}
		}
	}

	count := 0
	for v := range nodes {
		if v.X < 0 || v.Y < 0 || v.X >= len(loader.Lines[0]) || v.Y >= len(loader.Lines) {
			continue
		}
		count++
	}

	fmt.Printf("Day 8 Part 1: %d\n", count)
}

func Day8Part2() {
	loader, err := loader.NewLoader("2024/day8.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}*/

	antennas := map[rune][]containers.Point{}
	for y, line := range loader.Lines {
		for x, c := range line {
			if c != '.' {
				antennas[c] = append(antennas[c], containers.Point{X: x, Y: y})
			}
		}
	}

	nodes := map[containers.Point]bool{}
	for _, v := range antennas {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				dx := v[i].X - v[j].X
				dy := v[i].Y - v[j].Y

				currentX := v[i].X
				currentY := v[i].Y

				for currentX >= 0 && currentY >= 0 && currentX <= len(loader.Lines[0]) && currentY <= len(loader.Lines) {
					nodes[containers.Point{X: currentX, Y: currentY}] = true
					currentX += dx
					currentY += dy
				}

				currentX = v[j].X
				currentY = v[j].Y
				for currentX >= 0 && currentY >= 0 && currentX <= len(loader.Lines[0]) && currentY <= len(loader.Lines) {
					nodes[containers.Point{X: currentX, Y: currentY}] = true
					currentX -= dx
					currentY -= dy
				}
			}
		}
	}

	// kludge to make sure that the antenna locations are definitely included
	for _, a := range antennas {
		for _, v := range a {
			nodes[v] = true
		}
	}

	count := 0
	for v := range nodes {
		if v.X < 0 || v.Y < 0 || v.X >= len(loader.Lines[0]) || v.Y >= len(loader.Lines) {
			continue
		}
		count++
	}

	fmt.Printf("Day 8 Part 2: %d\n", count)
}
