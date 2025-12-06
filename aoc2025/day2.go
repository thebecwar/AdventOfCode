package aoc2025

import (
	"advent/loader"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
)

func Day2Part1() {
	loader, err := loader.NewLoader("2025/day2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//loader.Lines = []string{"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"}
	ranges := strings.Split(loader.Lines[0], ",")

	findNext := func(num int) int {
		digits := int(math.Floor(math.Log10(float64(num))) + 1)
		if digits%2 == 1 {
			digits++
			k := digits / 2
			// Odd digits
			next := int(math.Pow10(k-1) * (math.Pow10(k) + 1))
			return next
		}
		k := digits / 2
		firstHalf := num / int(math.Pow10(k))
		secondHalf := num % int(math.Pow10(k))
		if firstHalf > secondHalf {
			return firstHalf * int(math.Pow10(k)+1)
		} else if secondHalf > firstHalf {
			return (firstHalf + 1) * int(math.Pow10(k)+1)
		} else {
			return num
		}
	}
	findPrev := func(num int) int {
		digits := int(math.Floor(math.Log10(float64(num))) + 1)
		if digits%2 == 1 {
			digits--
			k := digits / 2
			// Odd digits
			next := int((math.Pow10(k) - 1) * (math.Pow10(k) + 1))
			return next
		}
		k := digits / 2
		firstHalf := num / int(math.Pow10(k))
		secondHalf := num % int(math.Pow10(k))
		if firstHalf < secondHalf {
			return firstHalf * int(math.Pow10(k)+1)
		} else if secondHalf < firstHalf {
			return (firstHalf - 1) * int(math.Pow10(k)+1)
		} else {
			return num
		}
	}

	sum := uint(0)

	for _, r := range ranges {
		ends := strings.Split(r, "-")
		if len(ends[0]) == len(ends[1]) && len(ends[0])%2 == 1 {
			// Impossible to have repetition in this range
			continue
		}
		start, _ := strconv.Atoi(ends[0])
		end, _ := strconv.Atoi(ends[1])
		_ = end

		first := findNext(start)
		last := findPrev(end)

		digits := int(math.Floor(math.Log10(float64(first))) + 1)
		k := digits / 2
		numerator := last - first
		denominator := int(math.Pow10(k) + 1)
		count := numerator / denominator
		if first < last {
			count++
		} else if first == last {
			count = 1
		}

		if math.Floor(math.Log10(float64(first))) != math.Floor(math.Log10(float64(last))) {
			fmt.Println("breakpoint")
		}
		sum += uint((float64(count) / 2) * (float64(first) + float64(last)))

	}

	fmt.Printf("Day 2 Part 1: %d\n", sum)
}

func Day2Part2() {
	loader, err := loader.NewLoader("2025/day2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//loader.Lines = []string{"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"}
	ranges := strings.Split(loader.Lines[0], ",")
	sum := uint64(0)

	exp := regexp2.MustCompile("^(.+)\\1{1,}$", regexp2.DefaultUnmarshalOptions)

	for _, r := range ranges {
		ends := strings.Split(r, "-")
		start, _ := strconv.Atoi(ends[0])
		end, _ := strconv.Atoi(ends[1])

		for i := start; i <= end; i++ {
			digits := strconv.Itoa(i)
			if matched, err := exp.MatchString(digits); err == nil && matched {
				//fmt.Println(digits)
				sum += uint64(i)
			} else if err != nil {
				panic(err)
			}
		}
	}

	fmt.Printf("Day 2 Part 2: %d\n", sum)
}
