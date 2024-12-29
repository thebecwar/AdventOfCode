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
type Cities map[string]City
type Day9State struct {
	AllDestinations *containers.List[string]
	Cities          Cities
}

// 0         1  2         3 4
// something to something = distance

func parseCityData(data []string) Day9State {
	result := Day9State{
		AllDestinations: containers.NewList[string](),
		Cities:          make(Cities),
	}
	allDestinations := map[string]bool{}
	for _, line := range data {
		nextLine := strings.Fields(line)
		from := nextLine[0]
		to := nextLine[2]
		allDestinations[from] = true
		allDestinations[to] = true
		distance, err := strconv.Atoi(nextLine[4])
		if err != nil {
			panic(err)
		}
		if _, ok := result.Cities[from]; !ok {
			result.Cities[from] = City{
				Name:         from,
				Destinations: map[string]int{},
			}
		}
		result.Cities[from].Destinations[to] = distance

		if _, ok := result.Cities[to]; !ok {
			result.Cities[to] = City{
				Name:         to,
				Destinations: map[string]int{},
			}
		}
		result.Cities[to].Destinations[from] = distance
	}
	for destination := range allDestinations {
		result.AllDestinations.Add(destination)
	}
	return result
}

func Day9Part1() {
	loader, err := loader.NewLoader("2015/day9.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"London to Dublin = 464",
		"London to Belfast = 518",
		"Dublin to Belfast = 141",
	}*/
	cities := parseCityData(loader.Lines)

	count := 0
	minRoute := ""
	minDistance := 10000000000
	maxRoute := ""
	maxDistance := -1
	for a := range cities.AllDestinations.PermutionIterator {
		count++
		distance := 0
		completedRoute := true
		for i := 0; i < len(a)-1; i++ {
			if from, ok := cities.Cities[a[i]]; !ok {
				completedRoute = false
				break
			} else {
				if to, ok := from.Destinations[a[i+1]]; !ok {
					completedRoute = false
					break
				} else {
					distance += to
				}
			}
		}
		if !completedRoute {
			continue
		}
		route := strings.Join(a, " -> ")
		if distance < minDistance {
			minDistance = distance
			minRoute = route
		}
		if distance > maxDistance {
			maxDistance = distance
			maxRoute = route
		}
	}

	fmt.Printf("Day 9 Part 1: %d (%s) [Checked: %d]\n", minDistance, minRoute, count)
	fmt.Printf("Day 9 Part 2: %d (%s)\n", maxDistance, maxRoute)
}

func Day9Part2() {
	return
}
