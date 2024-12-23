package aoc2015

import (
	"advent/containers"
	"advent/loader"
	"fmt"
	"strconv"
	"strings"
)

type City struct {
	Name         string
	Destinations map[string]int
}

// 0         1  2         3 4
// something to something = distance

func parseCityData(data []string) containers.List[City] {
	result := containers.List[City]{
		Items: []City{},
	}
	currentCity := City{}
	for _, line := range data {
		nextLine := strings.Fields(line)
		if nextLine[0] != currentCity.Name {
			result.Add(currentCity)
			currentCity = City{}
		}
		if currentCity.Name == "" {
			currentCity.Name = nextLine[0]
			currentCity.Destinations = make(map[string]int)
		}
		distance, err := strconv.Atoi(nextLine[4])
		if err != nil {
			panic(err)
		}
		currentCity.Destinations[nextLine[2]] = distance
	}
	return result
}

func Day9Part1() {
	loader, err := loader.NewLoader("2015/day9.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{
		"London to Dublin = 464",
		"London to Belfast = 518",
		"Dublin to Belfast = 141",
	}
	cities := parseCityData(loader.Lines)

	routes := map[string]int{}
	for a := range cities.PermutionIterator {
		distance := 0
		for i := 0; i < len(a)-1; i++ {
			// change this
			distance += cities.Items[i].Destinations[a[i+1].Name]
		}
		routes[""] = distance
	}

	fmt.Printf("Day 9 Part 1: %d\n", 0)
}

func Day9Part2() {
	loader, err := loader.NewLoader("2015/day9.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day 9 Part 2: %d\n", 0)
}
