package aoc2024

import (
	"advent/loader"
	"fmt"
	"sort"
	"strings"
)

type SafetyManualPageOrder map[int]map[int]bool

var predMap SafetyManualPageOrder

type SafetyPageSort []int

func (s SafetyPageSort) Len() int      { return len(s) }
func (s SafetyPageSort) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s SafetyPageSort) Less(i, j int) bool {
	rule, ok := predMap[s[i]]
	if !ok {
		return false
	}
	_, ok = rule[s[j]]
	return ok
}

func parseSafetyManual(lines []string) (SafetyManualPageOrder, [][]int) {
	predMap = make(SafetyManualPageOrder)
	pages := [][]int{}

	rules := true
	for _, line := range lines {
		if line == "" {
			rules = false
			continue
		}
		if rules {
			var a, b int
			fmt.Sscanf(line, "%d|%d", &a, &b)
			if _, ok := predMap[a]; !ok {
				predMap[a] = make(map[int]bool)
			}
			predMap[a][b] = true
		} else {
			raw := strings.Split(line, ",")
			current := []int{}
			for _, r := range raw {
				var a int
				fmt.Sscanf(r, "%d", &a)
				current = append(current, a)
			}
			pages = append(pages, current)
		}
	}
	return predMap, pages
}

func Day5Part1() {
	loader, err := loader.NewLoader("2024/day5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
		"",
		"75,47,61,53,29",
		"97,61,53,29,13",
		"75,29,13",
		"75,97,47,61,53",
		"61,13,29",
		"97,13,75,29,47",
	}*/

	_, pages := parseSafetyManual(loader.Lines)

	correct := 0
	centerSum := 0
	for _, page := range pages {
		before := strings.Join(strings.Fields(fmt.Sprint(page)), ",")
		sort.Sort(SafetyPageSort(page))
		after := strings.Join(strings.Fields(fmt.Sprint(page)), ",")
		if before != after {
			continue
		}
		correct++
		middle := page[len(page)/2]
		centerSum += middle
	}

	fmt.Printf("Day 5, Part 1: %d (%d pages correct)\n", centerSum, correct)
}
func Day5Part2() {
	loader, err := loader.NewLoader("2024/day5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
		"",
		"75,47,61,53,29",
		"97,61,53,29,13",
		"75,29,13",
		"75,97,47,61,53",
		"61,13,29",
		"97,13,75,29,47",
	}*/

	_, pages := parseSafetyManual(loader.Lines)

	incorrect := 0
	centerSum := 0
	for _, page := range pages {
		before := strings.Join(strings.Fields(fmt.Sprint(page)), ",")
		sort.Sort(SafetyPageSort(page))
		after := strings.Join(strings.Fields(fmt.Sprint(page)), ",")
		if before == after {
			continue
		}
		incorrect++
		middle := page[len(page)/2]
		centerSum += middle
	}

	fmt.Printf("Day 5, Part 1: %d (%d pages incorrect)\n", centerSum, incorrect)
}
