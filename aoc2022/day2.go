package aoc2022

import (
	"advent/loader"
	"fmt"
	"strings"
)

type rpsHand int
type outcome int

const (
	rock     rpsHand = 1
	paper    rpsHand = 2
	scissors rpsHand = 3

	lose outcome = 1
	draw outcome = 2
	win  outcome = 3
)

var rpsLookup = map[string]rpsHand{
	"A": rock,
	"B": paper,
	"C": scissors,
	"X": rock,
	"Y": paper,
	"Z": scissors,
}
var winNeed = map[rpsHand]rpsHand{
	scissors: rock,
	rock:     paper,
	paper:    scissors,
}
var loseNeed = map[rpsHand]rpsHand{
	scissors: paper,
	rock:     scissors,
	paper:    rock,
}

func IsWin(player, opponent rpsHand) bool {
	return (player == rock && opponent == scissors) ||
		(player == paper && opponent == rock) ||
		(player == scissors && opponent == paper)
}
func HandScore(player, opponent rpsHand) int {
	handScore := int(player)
	// 0 loss, 3 tie, 6 win
	if player == opponent {
		handScore += 3
	} else if IsWin(player, opponent) {
		handScore += 6
	}
	return handScore
}
func Part2HandScore(player, opponent rpsHand) int {
	oc := outcome(player)
	if oc == lose {
		player = loseNeed[opponent]
	} else if oc == win {
		player = winNeed[opponent]
	} else {
		player = opponent
	}
	return HandScore(player, opponent)
}

func Day2Part1() {
	loader, err := loader.NewLoader("2022/day2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	/*loader.Lines = []string{
		"A Y",
		"B X",
		"C Z",
	}*/

	totalScore := 0
	for _, line := range loader.Lines {
		fields := strings.Fields(line)
		opponentHand := rpsLookup[fields[0]]
		playerHand := rpsLookup[fields[1]]
		totalScore += HandScore(playerHand, opponentHand)
	}
	fmt.Printf("Day 2 Part 1: %d\n", totalScore)
}
func Day2Part2() {
	loader, err := loader.NewLoader("2022/day2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	/*loader.Lines = []string{
		"A Y",
		"B X",
		"C Z",
	}*/

	totalScore := 0
	for _, line := range loader.Lines {
		fields := strings.Fields(line)
		opponentHand := rpsLookup[fields[0]]
		playerHand := rpsLookup[fields[1]]
		totalScore += Part2HandScore(playerHand, opponentHand)
	}
	fmt.Printf("Day 2 Part 2: %d\n", totalScore)

}
