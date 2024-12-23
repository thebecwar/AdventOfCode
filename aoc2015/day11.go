package aoc2015

import (
	"advent/loader"
	"fmt"
)

func NextPassword(password string) string {
	chars := []byte(password)
	for i := len(chars) - 1; i >= 0; i-- {
		if chars[i] < 'z' {
			chars[i]++
			break
		} else {
			chars[i] = 'a'
		}
	}

	return string(chars)
}
func IsValid(password string) bool {
	hasStraight := false
	noIol := true
	hasTwoPairs := false
	for i := 0; i < len(password); i++ {
		// Has straight
		if i < len(password)-2 {
			if password[i+1] == password[i]+1 && password[i+2] == password[i]+2 {
				hasStraight = true
			}
		}
		if password[i] == 'i' || password[i] == 'o' || password[i] == 'l' {
			noIol = false
		}
		if i < len(password)-1 {
			if !hasTwoPairs && password[i] == password[i+1] {
				for j := i + 2; j < len(password)-1; j++ {
					if password[j] == password[j+1] && password[j] != password[i] {
						hasTwoPairs = true
						break
					}
				}
			}
		}
	}
	return hasStraight && noIol && hasTwoPairs
}

func Day11Part1() {
	loader, err := loader.NewLoader("2015/day11.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	count := 0
	password := loader.Lines[0]
	for !IsValid(password) {
		password = NextPassword(password)
		count++
	}
	fmt.Printf("Day 11 Part 1: %s (Tried: %d)\n", password, count)

	count = 0
	password = NextPassword(password)
	for !IsValid(password) {
		password = NextPassword(password)
		count++
	}
	fmt.Printf("Day 11 Part 2: %s (Tried: %d)\n", password, count)
}

func Day11Part2() {
}
