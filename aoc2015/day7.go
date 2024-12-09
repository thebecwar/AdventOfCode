package aoc2015

import (
	"advent/loader"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

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

func parseWireInfo(line string) (*WireNode, error) {
	parts := strings.Fields(line)
	if len(parts) == 3 {
		wireNode := &WireNode{
			Name:      parts[2],
			Operator:  "IDENT",
			rightName: parts[0],
		}
		if unicode.IsNumber(rune(parts[0][0])) {
			uintValue, err := strconv.ParseUint(parts[0], 10, 32)
			if err != nil {
				return nil, err
			}
			wireNode.Value = uint16(uintValue)
			wireNode.HasValue = true
		}
		return wireNode, nil
	}
	if len(parts) == 4 {
		wireNode := &WireNode{
			Name:      parts[3],
			Operator:  "NOT",
			rightName: parts[1],
		}
		if unicode.IsNumber(rune(parts[1][0])) {
			uintValue, err := strconv.ParseUint(parts[1], 10, 32)
			if err != nil {
				return nil, err
			}
			wireNode.Value = ^uint16(uintValue)
			wireNode.HasValue = true
		}
		return &WireNode{
			Name:      parts[3],
			Operator:  "NOT",
			rightName: parts[1],
		}, nil
	}
	if len(parts) == 5 {
		return &WireNode{
			Name:      parts[4],
			Operator:  parts[1],
			leftName:  parts[0],
			rightName: parts[2],
		}, nil
	}

	return nil, fmt.Errorf("unknown wire info %s", line)
}
func parseWires(lines []string) map[string]*WireNode {
	wires := make(map[string]*WireNode)
	for _, line := range lines {
		wire, err := parseWireInfo(line)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		wires[wire.Name] = wire
	}

	var literalNode = func(s string) (uint16, bool) {
		if unicode.IsNumber(rune(s[0])) {
			uintValue, err := strconv.ParseUint(s, 10, 32)
			if err != nil {
				return 0, false
			}
			return uint16(uintValue), true
		}
		return 0, false
	}

	for _, wire := range wires {
		if wire.leftName != "" {
			if value, ok := literalNode(wire.leftName); ok {
				wire.Left = &WireNode{
					Name:     wire.leftName,
					Operator: "IDENT",
					Value:    value,
					HasValue: true,
				}
				wires[wire.leftName] = wire.Left
			} else {
				wire.Left = wires[wire.leftName]
			}
		}
		if wire.rightName != "" {
			if value, ok := literalNode(wire.rightName); ok {
				wire.Right = &WireNode{
					Name:     wire.rightName,
					Operator: "IDENT",
					Value:    value,
					HasValue: true,
				}
				wires[wire.rightName] = wire.Right
			} else {
				wire.Right = wires[wire.rightName]
			}
		}
	}
	return wires
}

func Day7Part1() {
	loader, err := loader.NewLoader("2015/day7.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	wires := parseWires(loader.Lines)
	if _, ok := wires["a"]; !ok {
		fmt.Println("missing wire a")
	}

	result, err := wires["a"].Evaluate()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Day 7 Part 1: %d\n", result)
}

func Day7Part2() {
	loader, err := loader.NewLoader("2015/day7.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	wires := parseWires(loader.Lines)
	if _, ok := wires["a"]; !ok {
		fmt.Println("missing wire a")
	}

	result, err := wires["a"].Evaluate()
	if err != nil {
		fmt.Println(err)
		return
	}

	wires = parseWires(loader.Lines)
	wires["b"].Value = result
	wires["b"].HasValue = true

	result, err = wires["a"].Evaluate()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Day 7 Part 2: %d\n", result)
}
