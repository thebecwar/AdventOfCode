package aoc2023

import (
	"advent/loader"
	"fmt"
	"strconv"
	"strings"
)

type SeedMapping struct {
	Destination int
	SourceStart int
	SourceRange int
}

func (sm *SeedMapping) Contains(seed int) bool {
	if seed >= sm.SourceStart && seed < sm.SourceStart+sm.SourceRange {
		return true
	}
	return false
}

type SeedMap struct {
	Mappings []SeedMapping
}

func (sm *SeedMap) GetDestination(seed int) int {
	for _, mapping := range sm.Mappings {
		if mapping.Contains(seed) {
			return mapping.Destination + seed - mapping.SourceStart
		}
	}
	return seed
}

type SeedMappings struct {
	SeedToSoil            SeedMap
	SoilToFertilizer      SeedMap
	FertilizerToWater     SeedMap
	WaterToLight          SeedMap
	LightToTemperature    SeedMap
	TemperatureToHumidity SeedMap
	HumidityToLocation    SeedMap
}

func (sms *SeedMappings) GetDestination(seed int, mapName string) int {
	result := seed
	result = sms.SeedToSoil.GetDestination(result)
	result = sms.SoilToFertilizer.GetDestination(result)
	result = sms.FertilizerToWater.GetDestination(result)
	result = sms.WaterToLight.GetDestination(result)
	result = sms.LightToTemperature.GetDestination(result)
	result = sms.TemperatureToHumidity.GetDestination(result)
	result = sms.HumidityToLocation.GetDestination(result)
	return result
}

func parseInput(lines []string) ([]int, SeedMappings) {
	seedsRaw := strings.Split(lines[0], " ")[1:]
	seeds := []int{}
	for _, raw := range seedsRaw {
		seed, _ := strconv.Atoi(raw)
		seeds = append(seeds, seed)
	}

	mappings := SeedMappings{}

	var currentMap *SeedMap = nil
	for i := 2; i < len(lines); i++ {
		trimmed := strings.TrimSpace(lines[i])
		if trimmed == "" {
			continue
		}
		if strings.Contains(trimmed, "map") {
			if strings.Contains(trimmed, "seed-to-soil") {
				currentMap = &mappings.SeedToSoil
			} else if strings.Contains(trimmed, "soil-to-fertilizer") {
				currentMap = &mappings.SoilToFertilizer
			} else if strings.Contains(trimmed, "fertilizer-to-water") {
				currentMap = &mappings.FertilizerToWater
			} else if strings.Contains(trimmed, "water-to-light") {
				currentMap = &mappings.WaterToLight
			} else if strings.Contains(trimmed, "light-to-temperature") {
				currentMap = &mappings.LightToTemperature
			} else if strings.Contains(trimmed, "temperature-to-humidity") {
				currentMap = &mappings.TemperatureToHumidity
			} else if strings.Contains(trimmed, "humidity-to-location") {
				currentMap = &mappings.HumidityToLocation
			}
			continue
		}
		items := strings.Split(trimmed, " ")
		destination, _ := strconv.Atoi(items[0])
		sourceStart, _ := strconv.Atoi(items[1])
		sourceRange, _ := strconv.Atoi(items[2])
		mapping := SeedMapping{
			Destination: destination,
			SourceStart: sourceStart,
			SourceRange: sourceRange,
		}
		currentMap.Mappings = append(currentMap.Mappings, mapping)
	}

	return seeds, mappings

}

func Day5Part1() {
	loader, err := loader.NewLoader("2023-day5-part1.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"seeds: 79 14 55 13",
			"",
			"seed-to-soil map:",
			"50 98 2",
			"52 50 48",
			"",
			"soil-to-fertilizer map:",
			"0 15 37",
			"37 52 2",
			"39 0 15",
			"",
			"fertilizer-to-water map:",
			"49 53 8",
			"0 11 42",
			"42 0 7",
			"57 7 4",
			"",
			"water-to-light map:",
			"88 18 7",
			"18 25 70",
			"",
			"light-to-temperature map:",
			"45 77 23",
			"81 45 19",
			"68 64 13",
			"",
			"temperature-to-humidity map:",
			"0 69 1",
			"1 0 69",
			"",
			"humidity-to-location map:",
			"60 56 37",
			"56 93 4",
			"",
		}
	*/

	seeds, maps := parseInput(loader.Lines)

	minLocation := -1
	for _, seed := range seeds {
		result := maps.GetDestination(seed, "all")
		if minLocation == -1 || result < minLocation {
			minLocation = result
		}
	}

	fmt.Printf("Day 5 Part 1: %d\n", minLocation)
}

func Day5Part2() {
	loader, err := loader.NewLoader("2023-day5-part1.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"seeds: 79 14 55 13",
			"",
			"seed-to-soil map:",
			"50 98 2",
			"52 50 48",
			"",
			"soil-to-fertilizer map:",
			"0 15 37",
			"37 52 2",
			"39 0 15",
			"",
			"fertilizer-to-water map:",
			"49 53 8",
			"0 11 42",
			"42 0 7",
			"57 7 4",
			"",
			"water-to-light map:",
			"88 18 7",
			"18 25 70",
			"",
			"light-to-temperature map:",
			"45 77 23",
			"81 45 19",
			"68 64 13",
			"",
			"temperature-to-humidity map:",
			"0 69 1",
			"1 0 69",
			"",
			"humidity-to-location map:",
			"60 56 37",
			"56 93 4",
			"",
		}
	*/

	seeds, maps := parseInput(loader.Lines)

	minLocation := -1
	for i := 0; i < len(seeds); i += 2 {
		seed := seeds[i]
		seedMax := seeds[i+1]

		for s := 0; s < seedMax; s++ {
			result := maps.GetDestination(seed+s, "all")
			if minLocation == -1 || result < minLocation {
				minLocation = result
			}
		}
	}

	fmt.Printf("Day 5 Part 2: %d\n", minLocation)
}
