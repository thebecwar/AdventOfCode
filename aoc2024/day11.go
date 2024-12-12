package aoc2024

import (
	"advent/loader"
	"fmt"
	"strconv"
	"strings"
)

type StoneSplit struct {
	Value  int
	Blinks int
}

var stoneMemo = make(map[StoneSplit]int)

func blinkStone(s int) (int, int) {
	str := strconv.Itoa(s)
	if s == 0 {
		return 1, -1
	} else if len(str)%2 == 0 {
		left, _ := strconv.Atoi(str[:len(str)/2])
		right, _ := strconv.Atoi(str[len(str)/2:])
		return left, right
	} else {
		return s * 2024, -1
	}
}

func blinkRecursive(s int, blinks int) int {
	if v, ok := stoneMemo[StoneSplit{Value: s, Blinks: blinks}]; ok {
		return v
	}
	if s < 0 {
		return 0
	}
	if blinks == 0 {
		stoneMemo[StoneSplit{Value: s, Blinks: blinks}] = 1
		return 1
	}
	left, right := blinkStone(s)
	current := blinkRecursive(left, blinks-1) + blinkRecursive(right, blinks-1)

	stoneMemo[StoneSplit{Value: s, Blinks: blinks}] = current
	return current
}

func Day11Part1() {
	loader, err := loader.NewLoader("2024/day11.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//loader.Lines = []string{"125 17"}

	raw := strings.Fields(loader.Lines[0])
	stones := make([]int, len(raw))
	for i, r := range raw {
		val, _ := strconv.Atoi(r)
		stones[i] = val
	}

	total := 0
	for _, s := range stones {
		next := blinkRecursive(s, 25)
		total += next
	}

	fmt.Printf("Day 11 Part 1: %d (cache keys: %d)\n", total, len(stoneMemo))
}

func Day11Part2() {
	loader, err := loader.NewLoader("2024/day11.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//loader.Lines = []string{"125 17"}

	raw := strings.Fields(loader.Lines[0])
	stones := make([]int, len(raw))
	for i, r := range raw {
		val, _ := strconv.Atoi(r)
		stones[i] = val
	}

	total := 0
	for _, s := range stones {
		total += blinkRecursive(s, 75)
	}

	fmt.Printf("Day 11 Part 2: %d (cache keys: %d)\n", total, len(stoneMemo))
}
