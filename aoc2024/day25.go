package aoc2024

import (
	"advent/loader"
	"fmt"
)

type Lock [5]int
type Key [5]int

type TumblerLocks struct {
	Locks  []Lock
	Keys   []Key
	Height int
}

func (t *TumblerLocks) KeyFits(lock, key int) bool {
	l := t.Locks[lock]
	k := t.Keys[key]

	for i := 0; i < len(l); i++ {
		if l[i]+k[i] > t.Height {
			return false
		}
	}
	return true
}

func parseTumblerLocks(data []string) *TumblerLocks {
	result := &TumblerLocks{}
	for i := 0; i < len(data); i++ {
		if data[i] == "" {
			result.Height = i
			break
		}
	}

	for i := 0; i < len(data); i += result.Height + 1 {
		item := [5]int{}
		for r := 0; r < result.Height; r++ {
			for j := 0; j < len(data[i]); j++ {
				if data[i+r][j] == '#' {
					item[j]++
				}
			}
		}
		if data[i][0] == '#' {
			result.Locks = append(result.Locks, item)
		} else {
			result.Keys = append(result.Keys, item)
		}
	}

	return result
}

func Day25Part1() {
	loader, err := loader.NewLoader("2024/day25.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"#####",
		".####",
		".####",
		".####",
		".#.#.",
		".#...",
		".....",
		"",
		"#####",
		"##.##",
		".#.##",
		"...##",
		"...#.",
		"...#.",
		".....",
		"",
		".....",
		"#....",
		"#....",
		"#...#",
		"#.#.#",
		"#.###",
		"#####",
		"",
		".....",
		".....",
		"#.#..",
		"###..",
		"###.#",
		"###.#",
		"#####",
		"",
		".....",
		".....",
		".....",
		"#....",
		"#.#..",
		"#.#.#",
		"#####",
	}*/

	tumblerLocks := parseTumblerLocks(loader.Lines)
	if tumblerLocks.Height == 0 {
		fmt.Println("No tumbler locks found")
	}
	fit := 0
	for lock := range tumblerLocks.Locks {
		for key := range tumblerLocks.Keys {
			if tumblerLocks.KeyFits(lock, key) {
				fit++
			}
		}
	}

	fmt.Printf("Day 25 Part 1: %d\n", fit)
}

func Day25Part2() {
	loader, err := loader.NewLoader("2024/day25.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 25 Part 2: %d\n", 0)
}
