package aoc2023

import (
	"advent/loader"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Round struct {
	Red   int
	Green int
	Blue  int
}
type Game struct {
	GameNumber int
	Rounds     []Round
}

func parseLine(line string) Game {
	game := Game{}
	gameRegex := regexp.MustCompile(`Game (\d+):`)
	pickRegex := regexp.MustCompile(`(\d+) (red|green|blue)`)

	gameMatch := gameRegex.FindStringSubmatch(line)
	game.GameNumber, _ = strconv.Atoi(gameMatch[1])

	line = strings.Replace(line, gameMatch[0], "", 1)
	rounds := strings.Split(line, ";")
	for _, round := range rounds {
		roundMatch := pickRegex.FindAllStringSubmatch(round, -1)
		if len(roundMatch) == 0 {
			continue
		}

		r := Round{}
		for _, match := range roundMatch {
			number, _ := strconv.Atoi(match[1])
			switch match[2] {
			case "red":
				r.Red = number
			case "green":
				r.Green = number
			case "blue":
				r.Blue = number
			}
		}
		game.Rounds = append(game.Rounds, r)
	}

	return game
}

func Day2Part1() {
	loader, err := loader.NewLoader("2023-day2-part1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	maximums := Round{Red: 12, Green: 13, Blue: 14}

	idSum := 0
	possible := 0

	for _, line := range loader.Lines {
		game := parseLine(line)
		impossible := false

		for _, round := range game.Rounds {
			if round.Red > maximums.Red || round.Green > maximums.Green || round.Blue > maximums.Blue {
				impossible = true
				break
			}
		}
		if !impossible {
			possible++
			idSum += game.GameNumber
		} else {
			fmt.Printf("Game %d is impossible (%s)\n", game.GameNumber, line)
		}
	}
	fmt.Printf("Day 2 Part 1: ID Sum: %d, Possible games: %d\n", idSum, possible)

}

func Day2Part2() {
	loader, err := loader.NewLoader("2023-day2-part1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	power := 0

	for _, line := range loader.Lines {
		game := parseLine(line)
		reds := 0
		greens := 0
		blues := 0
		for _, round := range game.Rounds {
			if round.Red > reds {
				reds = round.Red
			}
			if round.Green > greens {
				greens = round.Green
			}
			if round.Blue > blues {
				blues = round.Blue
			}
		}
		fmt.Printf("Game %d: power: %d\n", game.GameNumber, reds*greens*blues)
		power += reds * greens * blues
	}

	fmt.Printf("Day 2 Part 2: Total power: %d\n", power)
}
