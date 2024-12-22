package aoc2024

import (
	"advent/containers"
	"advent/loader"
	"fmt"
	"strings"
)

type Keypad map[byte]containers.Point
type KeypadAdjacency map[string]map[string][]string

func (k *Keypad) GetMove(from, to byte) []string {
	if from == to {
		return []string{"A"}
	}

	if fromPoint, ok := (*k)[from]; ok {
		if toPoint, ok := (*k)[to]; ok {
			horizontal := ""
			if fromPoint.X > toPoint.X {
				// left
				horizontal = strings.Repeat("<", fromPoint.X-toPoint.X)
			} else if fromPoint.X == toPoint.X {
				// No movement
			} else {
				// right or none
				horizontal = strings.Repeat(">", toPoint.X-fromPoint.X)
			}

			// If we're moving up, we want to move up then left/right. If down, left/right then down.
			vertical := ""
			if fromPoint.Y > toPoint.Y {
				// Up
				vertical = strings.Repeat("^", fromPoint.Y-toPoint.Y)
			} else if fromPoint.Y == toPoint.Y {
			} else {
				// down
				vertical = strings.Repeat("v", toPoint.Y-fromPoint.Y)
			}

			result := []string{}
			//Mids are (vertical first) f.x,t.y and (horizontal first) t.x,f.y.
			// Check to see if the midpoints are in the dead spot (0,3)
			if fromPoint.X != 0 || toPoint.Y != 3 {
				// vertical first
				result = append(result, vertical+horizontal+"A")
			}
			if toPoint.X != 0 || fromPoint.Y != 3 {
				// horizontal first
				result = append(result, horizontal+vertical+"A")
			}
			if len(result) > 1 && result[0] == result[1] {
				return []string{result[0]}
			}
			return result
		}
	}

	panic("No move possible")
}
func (k *Keypad) GenerateAdjacencyMap() KeypadAdjacency {
	options := []byte{'7', '8', '9', '4', '5', '6', '1', '2', '3', '-', '0', 'A'}
	adjacencyMap := map[string]map[string][]string{}
	for _, from := range options {
		adjacencyMap[string(from)] = map[string][]string{}
		for _, to := range options {
			adjacencyMap[string(from)][string(to)] = k.GetMove(from, to)
		}
	}
	return adjacencyMap
}

var keypadMap = Keypad{
	'7': {X: 0, Y: 0},
	'8': {X: 1, Y: 0},
	'9': {X: 2, Y: 0},
	'4': {X: 0, Y: 1},
	'5': {X: 1, Y: 1},
	'6': {X: 2, Y: 1},
	'1': {X: 0, Y: 2},
	'2': {X: 1, Y: 2},
	'3': {X: 2, Y: 2},
	'-': {X: 0, Y: 3},
	'0': {X: 1, Y: 3},
	'A': {X: 2, Y: 3},
}

type Dpad map[byte]containers.Point

var dPadMap = Dpad{
	'-': {X: 0, Y: 0},
	'^': {X: 1, Y: 0},
	'A': {X: 2, Y: 0},
	'<': {X: 0, Y: 1},
	'v': {X: 1, Y: 1},
	'>': {X: 2, Y: 1},
}

func (d *Dpad) GetMove(from, to byte) []string {
	if from == to {
		return []string{"A"}
	}

	// We can't have 0,0 be the midpoint because that's the dead spot
	if fromPoint, ok := (*d)[from]; ok {
		if toPoint, ok := (*d)[to]; ok {
			horizontal := ""
			if fromPoint.X > toPoint.X {
				// left
				horizontal = strings.Repeat("<", fromPoint.X-toPoint.X)
			} else if fromPoint.X == toPoint.X {
				// No movement
			} else {
				// right or none
				horizontal = strings.Repeat(">", toPoint.X-fromPoint.X)
			}

			// opposite of the keypad map because the dead spot is top left instead of bottom left
			// Moving up, we want to move l/r then up. Moving down, down then l/r.
			vertical := ""
			if fromPoint.Y > toPoint.Y {
				// Up
				vertical = strings.Repeat("^", fromPoint.Y-toPoint.Y)
			} else if fromPoint.Y == toPoint.Y {
				// nothing
			} else {
				// down
				vertical = strings.Repeat("v", toPoint.Y-fromPoint.Y)
			}

			result := []string{}
			//Mids are (vertical first) f.x,t.y and (horizontal first) t.x,f.y.
			// Check to see if the midpoints are in the dead spot
			if fromPoint.X != 0 || toPoint.Y != 0 {
				// vertical first
				result = append(result, vertical+horizontal+"A")
			}
			if toPoint.X != 0 || fromPoint.Y != 0 {
				// horizontal first
				result = append(result, horizontal+vertical+"A")
			}
			if len(result) > 1 && result[0] == result[1] {
				return []string{result[0]}
			}
			return result
		}
	}

	panic("No move possible")
}

func (k KeypadAdjacency) GetMove(from, to string) []string {
	return k[from][to]
}

type keypadMemoKey struct {
	Start    string
	Depth    int
	IsKeypad bool
}

func (d *Dpad) GenerateAdjacencyMap() KeypadAdjacency {
	options := []byte{'-', '^', 'A', '<', 'v', '>'}
	adjacencyMap := map[string]map[string][]string{}
	for _, from := range options {
		adjacencyMap[string(from)] = map[string][]string{}
		for _, to := range options {
			adjacencyMap[string(from)][string(to)] = d.GetMove(from, to)
		}
	}
	return adjacencyMap
}

var dPadAdjacency = dPadMap.GenerateAdjacencyMap()
var keypadAdjacency = keypadMap.GenerateAdjacencyMap()

func (k KeypadAdjacency) GetLength(start string, depth int, memo *map[keypadMemoKey]int, isKeypad bool) int {
	if depth == 0 {
		return len(start)
	}
	if value, ok := (*memo)[keypadMemoKey{Start: start, Depth: depth, IsKeypad: isKeypad}]; ok {
		return value
	}

	if len(start) == 1 {
		return 1
	}

	next := 0
	for i := 0; i < len(start); i++ {
		from := "A"
		if i > 0 {
			from = string(start[i-1])
		}
		to := string(start[i])
		var nextItem []string
		if isKeypad {
			nextItem = keypadAdjacency[from][to]
		} else {
			nextItem = dPadAdjacency[from][to]
		}
		minNext := 100000000000
		for _, nextOption := range nextItem {
			if len(nextOption) > 0 && nextOption[len(nextOption)-1] != 'A' {
				fmt.Println("breakpoint")
			}
			option := k.GetLength(nextOption, depth-1, memo, false)
			if option < minNext {
				minNext = option
			}
		}
		next += minNext
	}

	(*memo)[keypadMemoKey{Start: start, Depth: depth, IsKeypad: isKeypad}] = next
	return next
}

func Day21Part1() {
	loader, err := loader.NewLoader("2024/day21.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"029A",
		"980A",
		"179A",
		"456A",
		"379A",
	}*/

	keypadAdjacency := keypadMap.GenerateAdjacencyMap()
	dPadAdjacency := dPadMap.GenerateAdjacencyMap()
	_ = keypadAdjacency
	_ = dPadAdjacency

	complexity := 0
	memo := map[keypadMemoKey]int{}
	for _, line := range loader.Lines {
		moves := dPadAdjacency.GetLength(line, 3, &memo, true)
		value := 0
		fmt.Sscanf(line, "%dA", &value)
		//fmt.Printf("Line: %s, Moves: %d, Value: %d, Complexity: %d\n", line, moves, value, moves*value)
		complexity += moves * value
	}

	fmt.Printf("Day 21 Part 1: %d\n", complexity)
}

func Day21Part2() {
	loader, err := loader.NewLoader("2024/day21.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{}*/

	keypadAdjacency := keypadMap.GenerateAdjacencyMap()
	dPadAdjacency := dPadMap.GenerateAdjacencyMap()
	_ = keypadAdjacency
	_ = dPadAdjacency

	complexity := 0
	memo := map[keypadMemoKey]int{}
	for _, line := range loader.Lines {
		moves := dPadAdjacency.GetLength(line, 26, &memo, true)
		value := 0
		fmt.Sscanf(line, "%dA", &value)
		complexity += moves * value
	}

	fmt.Printf("Day 21 Part 2: %d\n", complexity)
}
