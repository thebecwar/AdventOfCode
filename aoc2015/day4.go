package aoc2015

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func Day4Part1() {
	input := "yzbqklnj"

	for i := 0; i < 1000000; i++ {
		hash := md5.Sum([]byte(input + strconv.Itoa(i)))
		hashStr := fmt.Sprintf("%x", hash)
		if hashStr[:5] == "00000" {
			fmt.Printf("Day 4 Part 1: %d\n", i)
			return
		}
	}

	fmt.Printf("Day 4 Part 1: NOT FOUND\n")
}

func Day4Part2() {
	input := "yzbqklnj"

	for i := 0; i < 100000000; i++ {
		hash := md5.Sum([]byte(input + strconv.Itoa(i)))
		hashStr := fmt.Sprintf("%x", hash)
		if hashStr[:6] == "000000" {
			fmt.Printf("Day 4 Part 2: %d\n", i)
			return
		}
	}

	fmt.Printf("Day 4 Part 2: NOT FOUND\n")
}
