package aoc2025

import (
	"advent/loader"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day5Part1() {
	loader, err := loader.NewLoader("2025/day5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"3-5",
			"10-14",
			"16-20",
			"12-18",
			"",
			"1",
			"5",
			"8",
			"11",
			"17",
			"32",
		}
	*/
	type FreshRange struct {
		Min int
		Max int
	}
	isFresh := func(fr FreshRange, id int) bool {
		return id >= fr.Min && id <= fr.Max
	}

	fresh := []FreshRange{}
	i := 0
	for loader.Lines[i] != "" {
		r := strings.Split(loader.Lines[i], "-")
		min, err := strconv.Atoi(r[0])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(r[1])
		if err != nil {
			panic(err)
		}
		fresh = append(fresh, FreshRange{Min: min, Max: max})
		i++
	}
	i++
	count := 0
	for ; i < len(loader.Lines); i++ {
		ingredient, _ := strconv.Atoi(loader.Lines[i])
		for _, fr := range fresh {
			if isFresh(fr, ingredient) {
				count++
				break
			}
		}
	}

	fmt.Printf("Day 5 Part 1: %d\n", count)
}

func Day5Part2() {
	loader, err := loader.NewLoader("2025/day5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"3-5",
			"10-14",
			"16-20",
			"12-18",
			"",
			"1",
			"5",
			"8",
			"11",
			"17",
			"32",
		}
	*/

	type FreshRange struct {
		Min int
		Max int
	}
	fresh := []FreshRange{}
	i := 0
	for loader.Lines[i] != "" {
		r := strings.Split(loader.Lines[i], "-")
		min, err := strconv.Atoi(r[0])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(r[1])
		if err != nil {
			panic(err)
		}
		fresh = append(fresh, FreshRange{Min: min, Max: max})
		i++
	}
	sort.Slice(fresh, func(i, j int) bool {
		return fresh[i].Min < fresh[j].Min
	})

	merged := []FreshRange{fresh[0]}
	for i := 1; i < len(fresh); i++ {
		if fresh[i].Min <= merged[len(merged)-1].Max {
			merged[len(merged)-1].Max = max(fresh[i].Max, merged[len(merged)-1].Max)
		} else {
			merged = append(merged, fresh[i])
		}
	}

	count := 0
	for _, i := range merged {
		count += (i.Max - i.Min) + 1
	}

	fmt.Printf("Day 5 Part 2: %d\n", count)
}
