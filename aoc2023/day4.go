package aoc2023

import (
	"advent/loader"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	Id            int
	CardNumbers   []int
	ChoiceNumbers map[int]bool
	Copies        int
}

func (c *Card) Score() int {
	score := 0
	for _, num := range c.CardNumbers {
		if _, ok := c.ChoiceNumbers[num]; ok {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}
	return score
}

func (c *Card) Matches() int {
	matches := 0
	for _, num := range c.CardNumbers {
		if _, ok := c.ChoiceNumbers[num]; ok {
			matches++
		}
	}
	return matches
}

func ParseCard(line string) *Card {
	card := &Card{Copies: 1}

	cardIdRegex := regexp.MustCompile(`Card\s+(\d+):`)
	cardIdMatch := cardIdRegex.FindStringSubmatch(line)
	id, _ := strconv.Atoi(cardIdMatch[1])
	card.Id = id

	line = strings.Replace(line, cardIdMatch[0], "", 1)
	parts := strings.Split(line, "|")

	numbersRegex := regexp.MustCompile(`\d+`)
	matches := numbersRegex.FindAllStringSubmatch(parts[0], -1)

	for _, match := range matches {
		val, _ := strconv.Atoi(match[0])
		card.CardNumbers = append(card.CardNumbers, val)
	}

	matches = numbersRegex.FindAllStringSubmatch(parts[1], -1)
	card.ChoiceNumbers = make(map[int]bool)
	for _, match := range matches {
		val, _ := strconv.Atoi(match[0])
		card.ChoiceNumbers[val] = true
	}

	return card
}

func ParseCards(lines []string) []*Card {
	cards := []*Card{}

	for _, line := range lines {
		cards = append(cards, ParseCard(line))
	}

	return cards
}

func Day4Part1() {
	loader, err := loader.NewLoader("2023/day4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
		}
	*/

	cards := ParseCards(loader.Lines)

	sum := 0
	for _, card := range cards {
		sum += card.Score()
	}
	fmt.Printf("Day 4 Part 1: %d\n", sum)
}
func Day4Part2() {
	loader, err := loader.NewLoader("2023/day4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
		}
	*/

	cards := ParseCards(loader.Lines)

	for i, card := range cards {
		matches := card.Matches()
		if card.Score() > 0 {
			for j := i + 1; j < len(cards) && j <= i+matches; j++ {
				cards[j].Copies += card.Copies
			}
		}
	}

	count := 0
	for _, card := range cards {
		count += card.Copies
	}

	fmt.Printf("Day 4 Part 2: %d\n", count)
}
