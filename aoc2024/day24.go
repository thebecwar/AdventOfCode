package aoc2024

import (
	"advent/loader"
	"fmt"
	"math"
	"sort"
	"strings"
)

// 2015 Day 7 - Little bobby tables' wiring kits
type WireNode struct {
	Name     string
	Operator string
	Value    uint16
	HasValue bool
	Left     *WireNode
	Right    *WireNode

	rightName string
	leftName  string
}

func (w *WireNode) Matches(x, y, op string) bool {
	if w.Operator != op {
		return false
	}
	if w.leftName == x && w.rightName == y {
		return true
	}
	if w.leftName == y && w.rightName == x {
		return true
	}
	return false
}

func (w *WireNode) Evaluate() (uint16, error) {
	if w.HasValue || w.Operator == "" {
		return w.Value, nil
	}

	var err error
	var left uint16
	var right uint16

	if w.Left != nil {
		left, err = w.Left.Evaluate()
		if err != nil {
			return 0, err
		}
	}
	if w.Right != nil {
		right, err = w.Right.Evaluate()
		if err != nil {
			return 0, err
		}
	}

	switch w.Operator {
	case "AND":
		w.Value = left & right
	case "OR":
		w.Value = left | right
	case "XOR": // Added XOR for 2024 problem
		w.Value = left ^ right
	case "LSHIFT":
		w.Value = left << right
	case "RSHIFT":
		w.Value = left >> right
	case "NOT":
		w.Value = ^right
	case "IDENT":
		w.Value = right
	default:
		return 0, fmt.Errorf("unknown operator %s", w.Operator)
	}
	w.HasValue = true
	return w.Value, nil
}

func parseWire(wire string) *WireNode {
	if strings.Contains(wire, ":") {
		// Literal value
		fields := strings.Split(wire, ": ")
		name := fields[0]
		value := 0
		if fields[1] == "1" {
			value = 1
		}
		return &WireNode{
			Name:     name,
			Value:    uint16(value),
			HasValue: true,
		}
	} else {
		// Operator
		fields := strings.Fields(wire)
		return &WireNode{
			Name:      fields[4],
			Operator:  fields[1],
			leftName:  fields[0],
			rightName: fields[2],
		}
	}
}

func parseWires(data []string) map[string]*WireNode {
	wires := map[string]*WireNode{}
	for _, line := range data {
		if line == "" {
			continue
		}
		wire := parseWire(line)
		// Do a little assumption validation
		if wire.leftName != "" {
			if wire.leftName[0] == 'x' && wire.rightName[0] != 'y' {
				fmt.Println("Invalid wire", wire)
			}
			if wire.leftName[0] == 'y' && wire.rightName[0] != 'x' {
				fmt.Println("Invalid wire", wire)
			}

			if wire.leftName[0] == 'x' || wire.leftName[0] == 'y' {
				if wire.leftName[1] != wire.rightName[1] || wire.leftName[2] != wire.rightName[2] {
					fmt.Println("Invalid wire", wire)
					return nil
				}
			}
		}
		wires[wire.Name] = wire
	}

	// Link the nodes in the graph
	for _, wire := range wires {
		if wire.leftName != "" {
			wire.Left = wires[wire.leftName]
		}
		if wire.rightName != "" {
			wire.Right = wires[wire.rightName]
		}
	}

	return wires
}

type WireSet map[string]*WireNode

func (w *WireSet) GetAddendBitLength() int {
	for i := 0; i < 64; i++ {
		if _, ok := (*w)[fmt.Sprintf("z%02d", i)]; !ok {
			return i
		}
	}
	return 0
}
func (w *WireSet) ResetOutput() {
	for _, wire := range *w {
		if wire.Operator != "" {
			wire.HasValue = false
		}
	}
}
func (w *WireSet) SetValues(x, y uint16) {
	bits := w.GetAddendBitLength()
	for i := 0; i < bits; i++ {
		(*w)[fmt.Sprintf("x%02d", i)].Value = (x >> i) & 1
		(*w)[fmt.Sprintf("y%02d", i)].Value = (y >> i) & 1
	}
	w.ResetOutput()
}
func (w *WireSet) Evaluate(x, y uint16) int {
	w.SetValues(x, y)

	total := 0
	for i := 0; i < w.GetAddendBitLength(); i++ {
		v, _ := (*w)[fmt.Sprintf("z%02d", i)].Evaluate()
		total += int(v) * int(math.Pow(2, float64(i)))
	}
	return total
}
func (w *WireSet) Swap(a, b string) {
	(*w)[a].Name = b
	(*w)[b].Name = a
	(*w)[a], (*w)[b] = (*w)[b], (*w)[a]
}

func Day24Part1() {
	loader, err := loader.NewLoader("2024/day24.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"x00: 1",
		"x01: 0",
		"x02: 1",
		"x03: 1",
		"x04: 0",
		"y00: 1",
		"y01: 1",
		"y02: 1",
		"y03: 1",
		"y04: 1",

		"ntg XOR fgs -> mjb",
		"y02 OR x01 -> tnw",
		"kwq OR kpj -> z05",
		"x00 OR x03 -> fst",
		"tgd XOR rvg -> z01",
		"vdt OR tnw -> bfw",
		"bfw AND frj -> z10",
		"ffh OR nrd -> bqk",
		"y00 AND y03 -> djm",
		"y03 OR y00 -> psh",
		"bqk OR frj -> z08",
		"tnw OR fst -> frj",
		"gnj AND tgd -> z11",
		"bfw XOR mjb -> z00",
		"x03 OR x00 -> vdt",
		"gnj AND wpb -> z02",
		"x04 AND y00 -> kjc",
		"djm OR pbm -> qhw",
		"nrd AND vdt -> hwm",
		"kjc AND fst -> rvg",
		"y04 OR y02 -> fgs",
		"y01 AND x02 -> pbm",
		"ntg OR kjc -> kwq",
		"psh XOR fgs -> tgd",
		"qhw XOR tgd -> z09",
		"pbm OR djm -> kpj",
		"x03 XOR y03 -> ffh",
		"x00 XOR y04 -> ntg",
		"bfw OR bqk -> z06",
		"nrd XOR fgs -> wpb",
		"frj XOR qhw -> z04",
		"bqk OR frj -> z07",
		"y03 OR x01 -> nrd",
		"hwm AND bqk -> z03",
		"tgd XOR rvg -> z12",
		"tnw OR pbm -> gnj",
	}*/

	wires := parseWires(loader.Lines)

	total := 0
	i := 0
	for {
		name := fmt.Sprintf("z%02d", i)
		w, ok := wires[name]
		if !ok {
			break
		}
		value, err := w.Evaluate()
		if err != nil {
			fmt.Println(err)
			return
		}
		total += int(value) * int(math.Pow(2, float64(i)))
		i++
	}

	fmt.Printf("Day 24 Part 1: %d\n", total)
}

func Day24Part2() {
	loader, err := loader.NewLoader("2024/day24.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	wires := parseWires(loader.Lines)

	// Try 1: Check the gates match a full adder
	// Note: No gate takes invalid xNN or yNN combinations, and every half sum and half carry is covered.

	// Find all the half sum bits and carry bits
	halfSums := map[string]string{}
	carryBits := map[string]string{}
	otherXors := map[string]*WireNode{} // Map from any wire name to xors that use it
	otherAnds := map[string]*WireNode{} // Map from any wire name to ands that use it
	orGates := map[string]*WireNode{}   // Map from any wire name to ors that use it
	for _, w := range wires {
		if w.Operator == "XOR" && (w.leftName[0] == 'x' || w.leftName[0] == 'y') {
			key := "h" + w.leftName[1:]
			halfSums[key] = w.Name
		}
		if w.Operator == "AND" && (w.leftName[0] == 'x' || w.leftName[0] == 'y') {
			key := "c" + w.leftName[1:]
			carryBits[key] = w.Name
		}
		if w.leftName != "" {
			if w.leftName[0] != 'x' && w.leftName[0] != 'y' {
				if w.Operator == "XOR" {
					otherXors[w.leftName] = w
					otherXors[w.rightName] = w
				}
				if w.Operator == "AND" {
					otherAnds[w.leftName] = w
					otherAnds[w.rightName] = w
				}
				if w.Operator == "OR" {
					orGates[w.leftName] = w
					orGates[w.rightName] = w
				}
			}
		}
	}

	// Verify z00 as a special case
	if halfSums["h00"] != "z00" {
		fmt.Println("z00 is not a half sum")
		return
	}

	// Sum zN = xN xor yN xor cN-1
	// Carry cN = (xN and yN) xor (cN-1 AND (xN xor yN))
	// xN and yN == halfSum

	// Find the components of z2 - h02 xor c01
	carryOuts := map[string]string{
		"c00": carryBits["c00"],
	}
	wrongNodes := map[string]bool{}
	for i := 1; i < 44; i++ {
		// sum == a xor b xor cIn == halfSum xor cIn
		halfSum := halfSums[fmt.Sprintf("h%02d", i)]
		cIn := carryOuts[fmt.Sprintf("c%02d", i-1)]

		// Find the sum wire
		sum, ok := otherXors[halfSum]
		if !ok {
			// missing the half side of the full sum
			wrongNodes[halfSum] = true
			sum, ok = otherXors[cIn]
		}
		if !ok {
			// Missing the carry side of the full sum
			wrongNodes[cIn] = true
		} else {
			if sum.leftName != cIn && sum.rightName != cIn {
				// Missing the carry in
				wrongNodes[cIn] = true
			}
			// Check that we're writing the correct bit
			if sum.Name != fmt.Sprintf("z%02d", i) {
				wrongNodes[sum.Name] = true
				wrongNodes[fmt.Sprintf("z%02d", i)] = true
			}
		}

		// Figure out the next carry
		// carry == a and b or (cIn and (a xor b))
		cInCarry, ok := otherAnds[cIn]
		if !ok {
			// Missing half sum carry
			wrongNodes[cIn] = true
			cInCarry, ok = otherAnds[halfSum]
		}
		if !ok {
			// Missing the half sum side of the carry
			wrongNodes[halfSum] = true
		} else {
			if cInCarry.leftName != halfSum && cInCarry.rightName != halfSum {
				// Missing the half sum
				wrongNodes[halfSum] = true
			}
		}

		halfCarry := carryBits[fmt.Sprintf("c%02d", i)]
		cOut, ok := orGates[halfCarry]
		if !ok {
			// Missing an or for the half sum carry
			wrongNodes[halfCarry] = true
			cOut, ok = orGates[cInCarry.Name]
		}
		if !ok {
			wrongNodes[cInCarry.Name] = true
		} else {
			carryOuts[fmt.Sprintf("c%02d", i)] = cOut.Name
			//we have the or, check that it's correct
			if cInCarry != nil {
				if cInCarry.leftName != cIn && cInCarry.rightName != cIn {
					// Missing the carry in
					wrongNodes[cIn] = true
				}
				if cInCarry.leftName != halfSum && cInCarry.rightName != halfSum {
					// Missing the half sum
					wrongNodes[halfSum] = true
				}
			}
		}

	}
	// Final bit is the previous carry out

	if len(wrongNodes) == 8 {
		allWrongNodes := []string{}
		for k := range wrongNodes {
			allWrongNodes = append(allWrongNodes, k)
		}
		sort.Strings(allWrongNodes)
		fmt.Println("Day 24 Part 2: %s\n", strings.Join(allWrongNodes, ","))
		return
	}

	fmt.Println("Day 24 Part 2: Not found")
}
