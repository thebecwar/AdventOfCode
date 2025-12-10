package aoc2025

import (
	"advent/loader"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day8Part1() {
	loader, err := loader.NewLoader("2025/day8.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"162,817,812",
			"57,618,57",
			"906,360,560",
			"592,479,940",
			"352,342,300",
			"466,668,158",
			"542,29,236",
			"431,825,988",
			"739,650,466",
			"52,470,668",
			"216,146,977",
			"819,987,18",
			"117,168,530",
			"805,96,715",
			"346,949,466",
			"970,615,88",
			"941,993,340",
			"862,61,35",
			"984,92,344",
			"425,690,689",
		}
	*/

	type JunctionBox struct {
		X, Y, Z int
		Circuit int
	}
	type JunctionBoxPair struct {
		A, B     *JunctionBox
		Distance float64
	}
	parseJunction := func(line string) *JunctionBox {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		return &JunctionBox{
			X:       x,
			Y:       y,
			Z:       z,
			Circuit: -1,
		}
	}
	parseJunctions := func(lines []string) []*JunctionBox {
		junctions := []*JunctionBox{}
		for _, l := range lines {
			junctions = append(junctions, parseJunction(l))
		}
		return junctions
	}
	distance := func(a, b *JunctionBox) float64 {
		dx := b.X - a.X
		dy := b.Y - a.Y
		dz := b.Z - a.Z

		return math.Sqrt(
			float64(dx*dx + dy*dy + dz*dz),
		)
	}
	makePairs := func(boxes []*JunctionBox) []JunctionBoxPair {
		pairs := []JunctionBoxPair{}
		for i := range boxes {
			for j := i + 1; j < len(boxes); j++ {
				pairs = append(pairs, JunctionBoxPair{
					A:        boxes[i],
					B:        boxes[j],
					Distance: distance(boxes[i], boxes[j]),
				})
			}
		}
		return pairs
	}
	maxCircuit := 0
	connectBoxes := func(a, b *JunctionBox, boxes []*JunctionBox) {
		newCircuitId := max(a.Circuit, b.Circuit)
		if newCircuitId == -1 {
			// easy
			a.Circuit = maxCircuit + 1
			b.Circuit = maxCircuit + 1
			maxCircuit++
		} else {
			if a.Circuit == -1 {
				a.Circuit = newCircuitId
			} else if a.Circuit != newCircuitId {
				old := a.Circuit
				for _, box := range boxes {
					if box.Circuit == old {
						box.Circuit = newCircuitId
					}
				}
			}

			if b.Circuit == -1 {
				b.Circuit = newCircuitId
			} else if b.Circuit != newCircuitId {
				old := b.Circuit
				for _, box := range boxes {
					if box.Circuit == old {
						box.Circuit = newCircuitId
					}
				}
			}
		}
	}
	type Circuit struct {
		Id, Nodes int
	}
	circuitSizes := func(boxes []*JunctionBox) []Circuit {
		counts := map[int]int{}
		for _, box := range boxes {
			counts[box.Circuit]++
		}
		result := []Circuit{}
		for id, count := range counts {
			result = append(result, Circuit{Id: id, Nodes: count})
		}
		return result
	}

	const n = 1000

	boxes := parseJunctions(loader.Lines)
	pairs := makePairs(boxes)
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].Distance < pairs[b].Distance
	})
	for i := range n {
		connectBoxes(pairs[i].A, pairs[i].B, boxes)
	}
	circuits := circuitSizes(boxes)
	sort.Slice(circuits, func(a, b int) bool {
		return circuits[a].Nodes > circuits[b].Nodes
	})
	first := 0
	if circuits[0].Id == -1 {
		first = 1
	}
	product := circuits[first].Nodes * circuits[first+1].Nodes * circuits[first+2].Nodes

	fmt.Printf("Day 8 Part 1: %d\n", product)
}

func Day8Part2() {
	loader, err := loader.NewLoader("2025/day8.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"162,817,812",
			"57,618,57",
			"906,360,560",
			"592,479,940",
			"352,342,300",
			"466,668,158",
			"542,29,236",
			"431,825,988",
			"739,650,466",
			"52,470,668",
			"216,146,977",
			"819,987,18",
			"117,168,530",
			"805,96,715",
			"346,949,466",
			"970,615,88",
			"941,993,340",
			"862,61,35",
			"984,92,344",
			"425,690,689",
		}
	*/

	type JunctionBox struct {
		X, Y, Z int
		Circuit int
	}
	type JunctionBoxPair struct {
		A, B     *JunctionBox
		Distance float64
	}
	parseJunction := func(line string) *JunctionBox {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		return &JunctionBox{
			X:       x,
			Y:       y,
			Z:       z,
			Circuit: -1,
		}
	}
	parseJunctions := func(lines []string) []*JunctionBox {
		junctions := []*JunctionBox{}
		for _, l := range lines {
			junctions = append(junctions, parseJunction(l))
		}
		return junctions
	}
	distance := func(a, b *JunctionBox) float64 {
		dx := b.X - a.X
		dy := b.Y - a.Y
		dz := b.Z - a.Z

		return math.Sqrt(
			float64(dx*dx + dy*dy + dz*dz),
		)
	}
	makePairs := func(boxes []*JunctionBox) []JunctionBoxPair {
		pairs := []JunctionBoxPair{}
		for i := range boxes {
			for j := i + 1; j < len(boxes); j++ {
				pairs = append(pairs, JunctionBoxPair{
					A:        boxes[i],
					B:        boxes[j],
					Distance: distance(boxes[i], boxes[j]),
				})
			}
		}
		return pairs
	}
	maxCircuit := 0
	connectBoxes := func(a, b *JunctionBox, boxes []*JunctionBox) {
		newCircuitId := max(a.Circuit, b.Circuit)
		if newCircuitId == -1 {
			// easy
			a.Circuit = maxCircuit + 1
			b.Circuit = maxCircuit + 1
			maxCircuit++
		} else {
			if a.Circuit == -1 {
				a.Circuit = newCircuitId
			} else if a.Circuit != newCircuitId {
				old := a.Circuit
				for _, box := range boxes {
					if box.Circuit == old {
						box.Circuit = newCircuitId
					}
				}
			}

			if b.Circuit == -1 {
				b.Circuit = newCircuitId
			} else if b.Circuit != newCircuitId {
				old := b.Circuit
				for _, box := range boxes {
					if box.Circuit == old {
						box.Circuit = newCircuitId
					}
				}
			}
		}
	}

	distinctCircuits := func(boxes []*JunctionBox) int {
		circuits := map[int]bool{}
		for _, b := range boxes {
			circuits[b.Circuit] = true
		}
		return len(circuits)
	}

	boxes := parseJunctions(loader.Lines)
	pairs := makePairs(boxes)
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].Distance < pairs[b].Distance
	})
	unconnectedNodes := len(boxes)
	current := 0
	keepConnecting := true
	for keepConnecting {
		next := pairs[current]

		if next.A.Circuit == -1 {
			unconnectedNodes--
		}
		if next.B.Circuit == -1 {
			unconnectedNodes--
		}
		connectBoxes(next.A, next.B, boxes)

		if unconnectedNodes == 0 {
			distinct := distinctCircuits(boxes)
			if distinct == 1 {
				keepConnecting = false
			}
		}

		current++
		if current >= len(pairs) {
			panic("oops")
		}
	}

	xProduct := pairs[current-1].A.X * pairs[current-1].B.X

	fmt.Printf("Day 8 Part 2: %d (Iterations: %d)\n", xProduct, current)
}
