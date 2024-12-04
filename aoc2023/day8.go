package aoc2023

import (
	"advent/loader"
	"fmt"
	"regexp"
	"strings"
)

type Node struct {
	Name      string
	LeftName  string
	RightName string
	Left      *Node
	Right     *Node
	Period    int
}

func (n *Node) IsTrapNode() bool {
	return n.LeftName == n.RightName && n.Name == n.LeftName && n.Name != "ZZZ"
}
func (n *Node) IsANode() bool {
	return n.Name[2] == 'A'
}
func (n *Node) IsZNode() bool {
	return n.Name[2] == 'Z'
}

type Tree struct {
	nodes map[string]*Node
}

var lineRegex = regexp.MustCompile(`^(\w+) = \((\w+), (\w+)\)$`)

func BuildTree(lines []string) *Tree {
	tree := &Tree{
		nodes: make(map[string]*Node),
	}
	for _, line := range lines {
		if line == "" {
			continue
		}
		node := &Node{}
		parsedParts := lineRegex.FindStringSubmatch(line)
		node.Name = parsedParts[1]
		node.LeftName = parsedParts[2]
		node.RightName = parsedParts[3]
		tree.nodes[node.Name] = node
	}
	for _, node := range tree.nodes {
		if node.LeftName != "" {
			node.Left = tree.nodes[node.LeftName]
		}
		if node.RightName != "" {
			node.Right = tree.nodes[node.RightName]
		}
	}
	return tree
}

func Day8Part1() {
	loader, err := loader.NewLoader("2023/day8.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"LLR",
			"",
			"AAA = (BBB, BBB)",
			"BBB = (AAA, ZZZ)",
			"ZZZ = (ZZZ, ZZZ)",
		}
	*/

	directions := strings.Split(loader.Lines[0], "")
	tree := BuildTree(loader.Lines[2:])

	steps := 0
	currentNode := tree.nodes["AAA"]
	for currentNode.Name != "ZZZ" {
		instruction := directions[steps%len(directions)]
		if instruction == "R" {
			currentNode = currentNode.Right
		} else {
			currentNode = currentNode.Left
		}
		if currentNode.IsTrapNode() {
			fmt.Println("Trap node found", currentNode.Name)
		}
		steps++
	}
	fmt.Printf("Day 8 part 1, steps: %d\n", steps)

}
func Day8Part2() {
	loader, err := loader.NewLoader("2023/day8.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"LR",
			"",
			"11A = (11B, XXX)",
			"11B = (XXX, 11Z)",
			"11Z = (11B, XXX)",
			"22A = (22B, XXX)",
			"22B = (22C, 22C)",
			"22C = (22Z, 22Z)",
			"22Z = (22B, 22B)",
			"XXX = (XXX, XXX)",
		}
	*/
	directions := strings.Split(loader.Lines[0], "")
	tree := BuildTree(loader.Lines[2:])

	nodes := []*Node{}
	for _, node := range tree.nodes {
		if node.IsANode() {
			nodes = append(nodes, node)
		}
	}

	for i := 0; i < len(nodes); i++ {
		steps := 0
		currentNode := nodes[i]
		for !currentNode.IsZNode() {
			instruction := directions[steps%len(directions)]
			if instruction == "R" {
				currentNode = currentNode.Right
			} else {
				currentNode = currentNode.Left
			}
			steps++
		}
		nodes[i].Period = steps
	}

	// Find the least common multiple of all periods
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	lcm := nodes[0].Period
	for i := 1; i < len(nodes); i++ {
		lcm = lcm * nodes[i].Period / gcd(lcm, nodes[i].Period)
	}

	fmt.Printf("Day 8 part 2, steps: %d\n", lcm)
}
