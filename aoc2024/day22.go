package aoc2024

import (
	"advent/loader"
	"fmt"
	"strconv"
)

type PriceSequence struct {
	DeltaMinus0 int
	DeltaMinus1 int
	DeltaMinus2 int
	DeltaMinus3 int
}

func (p *PriceSequence) ShiftIn(priceDelta int) {
	p.DeltaMinus3 = p.DeltaMinus2
	p.DeltaMinus2 = p.DeltaMinus1
	p.DeltaMinus1 = p.DeltaMinus0
	p.DeltaMinus0 = priceDelta
}
func (p *PriceSequence) String() string {
	return fmt.Sprintf("{%d %d %d %d}", p.DeltaMinus0, p.DeltaMinus1, p.DeltaMinus2, p.DeltaMinus3)
}

func nextSecret(secret int) int {
	// Step 1:
	secret = (secret ^ (secret * 64)) % 16777216
	// Step 2:
	secret = (secret ^ (secret / 32)) % 16777216
	// Step 3:
	secret = (secret ^ (secret * 2048)) % 16777216
	return secret
}

func price(secret int) int {
	return secret % 10
}

func Day22Part1() {
	loader, err := loader.NewLoader("2024/day22.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"1",
		"10",
		"100",
		"2024",
	}*/

	sum := 0
	for _, line := range loader.Lines {
		secret, _ := strconv.Atoi(line)
		for i := 0; i < 2000; i++ {
			secret = nextSecret(secret)
		}
		sum += secret
	}

	fmt.Printf("Day 22 Part 1: %d\n", sum)
}

func Day22Part2() {
	loader, err := loader.NewLoader("2024/day22.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"1",
		"2",
		"3",
		"2024",
	}*/

	// For each line, the first time we see a sequence of deltas, we add the price to the current value here
	sequencePrices := map[PriceSequence]int{}
	for _, line := range loader.Lines {
		secret, _ := strconv.Atoi(line)
		previousPrice := price(secret)
		sequence := PriceSequence{}
		seenSequences := map[PriceSequence]bool{}
		for i := 0; i < 2000; i++ {
			secret = nextSecret(secret)
			delta := price(secret) - previousPrice
			previousPrice = price(secret)
			sequence.ShiftIn(delta)
			if i >= 3 {
				if _, ok := seenSequences[sequence]; ok {
					// We've already seen this sequence of deltas
					continue
				}
				seenSequences[sequence] = true
				sequencePrices[sequence] += price(secret)
			}
		}
	}

	maxPrice := -1
	var maxSequence PriceSequence
	for s, p := range sequencePrices {
		if p > maxPrice {
			maxPrice = p
			maxSequence = s
		}
	}
	if maxPrice == -1 {
		fmt.Println("No sequence found")
		return
	}

	fmt.Printf("Day 22 Part 1: %d, (Sequence: %v)\n", maxPrice, maxSequence)
}
