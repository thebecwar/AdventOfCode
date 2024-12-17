package aoc2024

import (
	"advent/loader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type ComputerState struct {
	RegA       int
	RegB       int
	RegC       int
	Ip         int
	Program    []int
	CycleCount int
	Target     []int
}

func NewComputerState(setup []string) *ComputerState {
	c := &ComputerState{
		RegA:    0,
		RegB:    0,
		RegC:    0,
		Ip:      0,
		Program: []int{},
	}
	/*
	 */
	fmt.Sscanf(setup[0], "Register A: %d", &c.RegA)
	fmt.Sscanf(setup[1], "Register B: %d", &c.RegB)
	fmt.Sscanf(setup[2], "Register C: %d", &c.RegC)
	code := strings.Fields(setup[4])[1]
	programArray := strings.Split(code, ",")
	for _, raw := range programArray {
		n, _ := strconv.Atoi(raw)
		c.Program = append(c.Program, n)
	}
	return c
}

func (c *ComputerState) ComboValue(v int) int {
	switch v {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return c.RegA
	case 5:
		return c.RegB
	case 6:
		return c.RegC
	default:
		panic("Invalid combo value")
	}
}
func (c *ComputerState) Run() []int {
	output := []int{}
	for c.Ip < len(c.Program) {
		instruction := c.Program[c.Ip]
		operand := c.Program[c.Ip+1]

		switch instruction {
		case 0: // ADV
			c.RegA = c.RegA >> c.ComboValue(operand)
		case 1: // BXL
			c.RegB = c.RegB ^ operand
		case 2: // BST
			c.RegB = c.ComboValue(operand) & 0x07
		case 3: // JNZ
			if c.RegA != 0 {
				c.Ip = operand
			}
		case 4: // BXC
			c.RegB = c.RegB ^ c.RegC
		case 5: // OUT
			output = append(output, c.ComboValue(operand)%8)
			if len(c.Target) > 0 {
				if output[len(output)-1] != c.Target[len(output)-1] {
					return output
				}
			}
		case 6: // BDV
			c.RegB = c.RegA >> c.ComboValue(operand)
		case 7: // CDV
			c.RegC = c.RegA >> c.ComboValue(operand)
		default:
			panic("Invalid instruction")
		}

		if instruction != 3 || (instruction == 3 && c.RegA == 0) {
			c.Ip += 2
		}
		c.CycleCount++
	}
	return output
}

func Day17Part1() {
	loader, err := loader.NewLoader("2024/day17.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"Register A: 729",
		"Register B: 0",
		"Register C: 0",
		"",
		"Program: 0,1,5,4,3,0",
	}*/
	computer := NewComputerState(loader.Lines)
	output := computer.Run()
	result := []string{}
	for _, o := range output {
		result = append(result, strconv.Itoa(o))
	}

	fmt.Printf("Day 17 Part 1: %s\n", strings.Join(result, ","))
}

func Day17Part2() {
	loader, err := loader.NewLoader("2024/day17.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"Register A: 2024",
		"Register B: 0",
		"Register C: 0",
		"",
		"Program: 0,3,5,4,3,0",
	}*/

	target := strings.Fields(loader.Lines[4])[1]
	programTarget := []int{}
	for _, o := range strings.Split(target, ",") {
		n, _ := strconv.Atoi(o)
		programTarget = append(programTarget, n)
	}
	initialState := NewComputerState(loader.Lines)
	computer := NewComputerState(loader.Lines)
	//computer.Target = programTarget

	// regA == 8, output length == 2 (log2(8) == 3)
	// regA == 64, output length == 3 (log2(64) == 6)
	// regA == 512, output length == 4 (log2(512) == 9)
	// regA == 4096, output length == 5 (log2(4096) == 12)
	// regA == 32768, output length == 6 (log2(32768) == 15)
	// for output of length x, regA == 2^(3(x-1))

	// Outputs vary at intervals
	// output[0] varies every time
	// output[1] varies every 8 (2^3)
	// output[2] varies every 64 (2^6)
	// output[3] varies every 512 (2^9)
	// output[4] varies every 4096 (2^12)
	// output[5] varies every 32768 (2^15)
	// So if the last delta is in output[5], we can skip the next 2^15 values
	// Therefore for a diff in output[n] we can skip 2^(3(n)) states

	targetLength := float64(len(programTarget))
	lowerBound := int(math.Pow(2, 3*(targetLength-1)))
	upperBound := int(math.Pow(2, 3*targetLength))
	fmt.Printf("Bounds: %d - %d\n", lowerBound, upperBound)

	computer.RegA = 4096

	for i := lowerBound; i < upperBound+1; i++ {
		if i == initialState.RegA {
			continue
		}
		computer.RegA = i
		computer.RegB = initialState.RegB
		computer.RegC = initialState.RegC
		computer.Ip = 0
		computer.CycleCount = 0

		output := computer.Run()

		// Find the last place that the output differs from the target
		last := -1
		for j := len(output) - 1; j >= 0; j-- {
			if output[j] != programTarget[j] {
				last = j
				break
			}
		}
		if last < 0 {
			// We have no differences. Answer found
			fmt.Printf("Day 17 Part 2: %d\n", i)
			return
		}

		// Skip to the next value that will differ
		skip := int(math.Pow(2, 3*float64(last)))
		i += skip - 1
	}
	fmt.Println("Day 17 Part 2: Not found")
}
