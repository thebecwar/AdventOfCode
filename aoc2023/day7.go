package aoc2023

import (
	"advent/loader"
	"fmt"
	"sort"
)

var cardValues = map[byte]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}
var cardValuesPartTwo = map[byte]int{
	'J': 2,
	'2': 3,
	'3': 4,
	'4': 5,
	'5': 6,
	'6': 7,
	'7': 8,
	'8': 9,
	'9': 10,
	'T': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

type Hand struct {
	Cards     string
	Wager     int
	HandType  int
	GroupRank int
	High      byte
	Low       byte
	counts    map[byte]int
}

func (h *Hand) FiveOfAKind() (byte, bool) {
	return h.Cards[0], len(h.counts) == 1
}
func (h *Hand) FourOfAKind() (byte, bool) {
	if len(h.counts) != 2 {
		return 'X', false
	}
	for c, count := range h.counts {
		if count == 4 {
			return c, true
		}
	}
	return 'X', false
}
func (h *Hand) FullHouse() (byte, byte, bool) {
	high, low := byte('X'), byte('X')
	if len(h.counts) != 2 {
		return high, low, false
	}
	two, three := false, false
	for c, count := range h.counts {
		if count == 2 {
			two = true
			low = c
		}
		if count == 3 {
			three = true
			high = c
		}
	}
	return high, low, two && three
}
func (h *Hand) ThreeOfAKind() (byte, bool) {
	_, _, fh := h.FullHouse()
	if fh {
		return 'X', false
	}
	for c, count := range h.counts {
		if count == 3 {
			return c, true
		}
	}
	return 'X', false
}
func (h *Hand) TwoPair() (byte, byte, bool) {
	_, toak := h.ThreeOfAKind()
	if toak {
		return 'X', 'X', false
	}
	pairs := []byte{}
	for c, count := range h.counts {
		if count == 2 {
			pairs = append(pairs, c)
		}
	}
	sort.Sort(SortCards(pairs))
	if len(pairs) == 2 {
		return pairs[1], pairs[0], true
	}
	if len(pairs) == 1 {
		return pairs[0], 'X', false
	}
	return 'X', 'X', false
}
func (h *Hand) OnePair() (byte, bool) {
	p, _, twopair := h.TwoPair()
	if twopair == false {
		if p != 'X' {
			return p, true
		}
	}
	return 'X', false
}
func (h *Hand) HighCard() (byte, bool) {
	_, ok := h.OnePair()
	return h.Cards[len(h.Cards)-1], !ok
}
func (h *Hand) GetGroupRank() (byte, byte, int) {
	c, ok := h.FiveOfAKind()
	if ok {
		return c, 'X', 7
	}
	c, ok = h.FourOfAKind()
	if ok {
		return c, 'X', 6
	}
	high, low, ok := h.FullHouse()
	if ok {
		return high, low, 5
	}
	c, ok = h.ThreeOfAKind()
	if ok {
		return c, 'X', 4
	}
	high, low, ok = h.TwoPair()
	if ok {
		return high, low, 3
	}
	c, ok = h.OnePair()
	if ok {
		return c, 'X', 2
	}
	c, ok = h.HighCard()
	if ok {
		return c, 'X', 1
	}
	return 'X', 'X', 0
}

type SortCards []byte

func (a SortCards) Len() int      { return len(a) }
func (a SortCards) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortCards) Less(i, j int) bool {
	return cardValues[a[i]] < cardValues[a[j]]
}

type SortHands []Hand

func (a SortHands) Len() int      { return len(a) }
func (a SortHands) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortHands) Less(i, j int) bool {
	if a[i].GroupRank == a[j].GroupRank {
		for k := 0; k < 5; k++ {
			if a[i].Cards[k] != a[j].Cards[k] {
				return cardValues[a[i].Cards[k]] < cardValues[a[j].Cards[k]]
			}
		}
	}
	return a[i].GroupRank < a[j].GroupRank
}

type SortHandsPartTwo []Hand

func (a SortHandsPartTwo) Len() int      { return len(a) }
func (a SortHandsPartTwo) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortHandsPartTwo) Less(i, j int) bool {
	if a[i].GroupRank == a[j].GroupRank {
		for k := 0; k < 5; k++ {
			if a[i].Cards[k] != a[j].Cards[k] {
				return cardValuesPartTwo[a[i].Cards[k]] < cardValuesPartTwo[a[j].Cards[k]]
			}
		}
	}
	return a[i].GroupRank < a[j].GroupRank
}

func ParseHand(line string) Hand {
	var h Hand
	fmt.Sscanf(line, "%s %d", &h.Cards, &h.Wager)
	cards := []byte(h.Cards)
	//sort.Sort(SortCards(cards))
	h.Cards = string(cards)
	h.counts = make(map[byte]int)
	for _, c := range cards {
		h.counts[c]++
	}
	h.High, h.Low, h.GroupRank = h.GetGroupRank()
	return h
}

func Day7Part1() {
	loader, err := loader.NewLoader("2023-day7-part1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"32T3K 765",
			"T55J5 684",
			"KK677 28",
			"KTJJT 220",
			"QQQJA 483",
		}
	*/

	hands := []Hand{}

	for _, line := range loader.Lines {
		hand := ParseHand(line)
		hands = append(hands, hand)
	}
	sort.Sort(SortHands(hands))

	score := 0
	for i, hand := range hands {
		score += hand.Wager * (i + 1)
	}
	fmt.Printf("Day 7, Part 1: %d\n", score)

}
func Day7Part2() {
	loader, err := loader.NewLoader("2023-day7-part1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"32T3K 765",
			"T55J5 684",
			"KK677 28",
			"KTJJT 220",
			"QQQJA 483",
		}
	*/

	hands := []Hand{}

	for _, line := range loader.Lines {
		hand := ParseHand(line)
		count, ok := hand.counts['J']
		if ok {
			max := 0
			charMax := byte('X')
			delete(hand.counts, 'J')
			for c, count := range hand.counts {
				if count > max && c != 'J' {
					max = count
					charMax = c
				}
			}
			hand.counts[charMax] += count
			hand.High, hand.Low, hand.GroupRank = hand.GetGroupRank()
		}
		hands = append(hands, hand)
	}
	sort.Sort(SortHandsPartTwo(hands))

	score := 0
	for i, hand := range hands {
		score += hand.Wager * (i + 1)
	}
	fmt.Printf("Day 7, Part 2: %d\n", score)
}
