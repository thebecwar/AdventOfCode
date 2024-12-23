package aoc2024

import (
	"advent/loader"
	"fmt"
	"sort"
	"strings"
)

type LanParty struct {
	ComputerA string
	ComputerB string
	ComputerC string
}

func NewLanParty(computerA string, computerB string, computerC string) LanParty {
	machines := []string{computerA, computerB, computerC}
	sort.Strings(machines)
	return LanParty{
		ComputerA: machines[0],
		ComputerB: machines[1],
		ComputerC: machines[2],
	}
}

func findCycles(adjacency map[string]map[string]bool) map[LanParty]bool {
	cycles := map[LanParty]bool{}
	for computerA := range adjacency {
		for computerB := range adjacency[computerA] {
			for computerC := range adjacency[computerB] {
				if adjacency[computerC][computerA] {
					cycles[NewLanParty(computerA, computerB, computerC)] = true
				}
			}
		}
	}
	return cycles
}

func allCliques(adjacency map[string]map[string]bool) []map[string]bool {
	networks := []map[string]bool{}
	for computer := range adjacency {
		networks = append(networks, map[string]bool{computer: true})
	}

	for _, network := range networks {
		for candidate := range adjacency {
			all := true
			for c := range network {
				if _, ok := adjacency[candidate][c]; !ok {
					all = false
					break
				}
			}
			if all {
				network[candidate] = true
			}
		}
	}
	return networks
}

func Day23Part1() {
	loader, err := loader.NewLoader("2024/day23.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"kh-tc",
		"qp-kh",
		"de-cg",
		"ka-co",
		"yn-aq",
		"qp-ub",
		"cg-tb",
		"vc-aq",
		"tb-ka",
		"wh-tc",
		"yn-cg",
		"kh-ub",
		"ta-co",
		"de-co",
		"tc-td",
		"tb-wq",
		"wh-td",
		"ta-ka",
		"td-qp",
		"aq-cg",
		"wq-ub",
		"ub-vc",
		"de-ta",
		"wq-aq",
		"wq-vc",
		"wh-yn",
		"ka-de",
		"kh-ta",
		"co-tc",
		"wh-qp",
		"tb-vc",
		"td-yn",
	}*/

	adjacency := map[string]map[string]bool{}
	for _, line := range loader.Lines {
		computers := strings.Split(line, "-")
		if _, ok := adjacency[computers[0]]; !ok {
			adjacency[computers[0]] = map[string]bool{}
		}
		adjacency[computers[0]][computers[1]] = true
		if _, ok := adjacency[computers[1]]; !ok {
			adjacency[computers[1]] = map[string]bool{}
		}
		adjacency[computers[1]][computers[0]] = true
	}

	parties := findCycles(adjacency)

	count := 0
	for p := range parties {
		if p.ComputerA[0] == 't' || p.ComputerB[0] == 't' || p.ComputerC[0] == 't' {
			count++
		}
	}

	fmt.Printf("Day 23 Part 1: %d (Of %d lan parties)\n", count, len(parties))
}

func Day23Part2() {
	loader, err := loader.NewLoader("2024/day23.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"ka-co",
		"ta-co",
		"de-co",
		"ta-ka",
		"de-ta",
		"ka-de",
	}*/

	adjacency := map[string]map[string]bool{}
	for _, line := range loader.Lines {
		computers := strings.Split(line, "-")
		if _, ok := adjacency[computers[0]]; !ok {
			adjacency[computers[0]] = map[string]bool{}
		}
		adjacency[computers[0]][computers[1]] = true
		if _, ok := adjacency[computers[1]]; !ok {
			adjacency[computers[1]] = map[string]bool{}
		}
		adjacency[computers[1]][computers[0]] = true
	}

	cliques := allCliques(adjacency)
	maxSize := 0
	maxClique := map[string]bool{}
	for _, clique := range cliques {
		if len(clique) > maxSize {
			maxSize = len(clique)
			maxClique = clique
		}
	}

	keys := []string{}
	for k := range maxClique {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	password := strings.Join(keys, ",")

	fmt.Printf("Day 23 Part 2: %s\n", password)
}
