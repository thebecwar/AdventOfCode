package aoc2023

import (
	"fmt"
	"math"
)

type RaceRecord struct {
	Time     int
	Distance int
}

// Hold relates to time and distance as distance = time * hold - hold**2
// (0.5 * (time - sqrt(time**2 - 4 * record)) <= hold <= (0.5 * (time + sqrt(time**2 - 4*record)))
func (r *RaceRecord) HoldMin() int {
	time := float64(r.Time)
	distance := float64(r.Distance)
	bound := 0.5 * (time - math.Sqrt(math.Pow(time, 2)-4.0*distance))

	hold := int(math.Floor(bound))
	if r.Time*hold-hold*hold <= r.Distance {
		return hold + 1
	}
	return hold
}
func (r *RaceRecord) HoldMax() int {
	time := float64(r.Time)
	distance := float64(r.Distance)
	bound := 0.5 * (time + math.Sqrt(math.Pow(time, 2)-4.0*distance))

	hold := int(math.Ceil(bound))
	if r.Time*hold-hold*hold <= r.Distance {
		return hold - 1
	}
	return hold
}
func (r *RaceRecord) DistanceForHold(hold int) int {
	return r.Time*hold - hold*hold
}

var ExampleRaces []RaceRecord = []RaceRecord{
	RaceRecord{Time: 7, Distance: 9},
	RaceRecord{Time: 15, Distance: 40},
	RaceRecord{Time: 30, Distance: 200},
}
var Races []RaceRecord = []RaceRecord{
	RaceRecord{Time: 47, Distance: 282},
	RaceRecord{Time: 70, Distance: 1079},
	RaceRecord{Time: 75, Distance: 1147},
	RaceRecord{Time: 66, Distance: 1062},
}
var Part2ExampleRaces []RaceRecord = []RaceRecord{
	RaceRecord{Time: 71530, Distance: 940200},
}
var Part2Races []RaceRecord = []RaceRecord{
	RaceRecord{Time: 47707566, Distance: 282107911471062},
}

func Day6Part1() {
	// no file
	races := Races

	count := 0
	margin := 1
	for _, race := range races {
		// Solving the inequality dRecord <
		min := race.HoldMin()
		max := race.HoldMax()
		count += max - min + 1
		margin *= max - min + 1
	}

	fmt.Printf("Day 6 Part 1: Ways: %d Margin: %d\n", count, margin)

}
func Day6Part2() {
	// no file
	races := Part2Races

	count := 0
	margin := 1
	for _, race := range races {
		// Solving the inequality dRecord <
		min := race.HoldMin()
		max := race.HoldMax()
		count += max - min + 1
		margin *= max - min + 1
	}

	fmt.Printf("Day 6 Part 2: Ways: %d Margin: %d\n", count, margin)
}
