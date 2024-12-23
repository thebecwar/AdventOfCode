package aoc2024

import (
	"advent/containers"
	"advent/loader"
	"fmt"
	"strings"
)

var directionMap = map[string]*containers.Point{
	"v": &containers.Point{X: 0, Y: 1},
	"^": &containers.Point{X: 0, Y: -1},
	">": &containers.Point{X: 1, Y: 0},
	"<": &containers.Point{X: -1, Y: 0},
}

var emptyTile = "."
var wallTile = "#"
var boxTile = "O"
var robotTile = "@"
var bigBoxLeft = "["
var bigBoxRight = "]"

var mapExpansions = map[string][2]string{
	".": {".", "."},
	"#": {"#", "#"},
	"O": {"[", "]"},
	"@": {"@", "."},
}

type Day15State struct {
	Grid       *containers.Grid[string]
	Position   containers.Point
	Directions string
}

func (state Day15State) HasUgh() bool {
	for y := 0; y < state.Grid.Height(); y++ {
		for x := 0; x < state.Grid.Width()-1; x++ {
			if state.Grid.Cells[y][x] == bigBoxLeft && state.Grid.Cells[y][x+1] != bigBoxRight {
				return true
			}
		}
	}
	return false
}

func (state Day15State) MoveBigBoxes(direction string, lX, rX, leftY int) bool {
	// Swaps don't work (as easily) in the big box case, since we "create" empty spaces by moving the box

	// If our search region is half a box on either side, expand to cover all boxes.
	leftX := lX
	if *state.Grid.Get(leftX, leftY) == bigBoxRight {
		leftX--
	}
	rightX := rX
	if *state.Grid.Get(rightX, leftY) == bigBoxLeft {
		rightX++
	}

	canMove := true
	robotRow := false
	for i := leftX; i <= rightX; i++ {
		if *state.Grid.Get(i, leftY) == wallTile {
			return false
		}
		// If the previous row has a box at this position, we can't move if this row has a box, but we could move if the next row is free
		previous := *state.Grid.Get(i, leftY-directionMap[direction].Y)
		current := *state.Grid.Get(i, leftY)
		if previous == bigBoxLeft || previous == bigBoxRight {
			if current != emptyTile {
				canMove = false
			}
		}
		if previous == robotTile {
			robotRow = true
		}
	}

	if canMove && !robotRow {
		// Need to support the case here:
		// ###############
		// #..[][][].....#
		// #...[][]......#
		// #....[].......# // current
		// #..[]..[].....# // previous
		// #.....^.......#
		// ###############
		// We have a compatible row to make this move.
		for i := leftX; i <= rightX; i++ {
			//current := *state.Grid.Get(i, leftY)
			previous := *state.Grid.Get(i, leftY-directionMap[direction].Y)
			if previous == bigBoxLeft || previous == bigBoxRight {
				state.Grid.Set(i, leftY, &previous)
				state.Grid.Set(i, leftY-directionMap[direction].Y, &emptyTile) // Where the box used to be
			}
		}
		return true
	}

	next := state.MoveBigBoxes(direction, leftX, rightX, leftY+directionMap[direction].Y)
	if next && !robotRow {
		// We moved the next row, so we can move this row
		for i := leftX; i <= rX; i++ {
			previous := *state.Grid.Get(i, leftY-directionMap[direction].Y)
			if previous == bigBoxLeft || previous == bigBoxRight {
				state.Grid.Set(i, leftY, &previous)
				state.Grid.Set(i, leftY-directionMap[direction].Y, &emptyTile) // Where the box used to be
			}
		}
		return true
	} else if next {
		// implies robot row
		return true
	}

	return false

}
func (state *Day15State) MoveRecurse(direction string, x, y int, bigBoxMode bool) bool {
	if bigBoxMode &&
		(direction == "^" || direction == "v") &&
		(*state.Grid.Get(x+directionMap[direction].X, y+directionMap[direction].Y) == bigBoxLeft ||
			*state.Grid.Get(x+directionMap[direction].X, y+directionMap[direction].Y) == bigBoxRight) {
		leftX, leftY := x, y
		if *state.Grid.Get(x, y+directionMap[direction].Y) == bigBoxRight {
			leftX--
		}
		moved := state.MoveBigBoxes(direction, leftX, leftX+1, leftY+directionMap[direction].Y)
		if moved {
			// We need to move the robot if we moved the box stack
			state.Grid.Set(state.Position.X, state.Position.Y, &emptyTile)

			// We only do the special movement logic if we're moving up or down
			state.Position.Y += directionMap[direction].Y
			state.Grid.Set(state.Position.X, state.Position.Y, &robotTile)
		}
		return false
	}

	// If we are at a wall return false
	if *state.Grid.Get(x, y) == wallTile {
		return false
	}

	// If we're at an empty space, swap and return true
	if *state.Grid.Get(x, y) == emptyTile {
		current := *state.Grid.Get(x, y)
		previous := *state.Grid.Get(x-directionMap[direction].X, y-directionMap[direction].Y)
		state.Grid.Set(x, y, &previous)
		state.Grid.Set(x-directionMap[direction].X, y-directionMap[direction].Y, &current)
		if previous == robotTile {
			state.Position.X += directionMap[direction].X
			state.Position.Y += directionMap[direction].Y
		}
		return previous != robotTile
	}
	if (*state.Grid.Get(x, y) == bigBoxLeft || *state.Grid.Get(x, y) == bigBoxRight) && direction != ">" && direction != "<" {
		// Special case for moving big boxes up and down
		// If we're at a big box, we need to recursively check both halves of the box, but we can only move if _both_ halves can move
		leftX, leftY := x, y
		if *state.Grid.Get(x, y) == bigBoxRight {
			leftX--
		}
		moved := state.MoveBigBoxes(direction, leftX, leftX+1, leftY+directionMap[direction].Y)
		if moved {
			// We need to move the robot if we moved the box stack
			state.Grid.Set(state.Position.X, state.Position.Y, &emptyTile)

			// We only do the special movement logic if we're moving up or down
			state.Position.Y += directionMap[direction].Y
			state.Grid.Set(state.Position.X, state.Position.Y, &robotTile)
		}
		return false
	}
	next := state.MoveRecurse(direction, x+directionMap[direction].X, y+directionMap[direction].Y, false)
	if next {
		current := *state.Grid.Get(x, y)
		previous := *state.Grid.Get(x-directionMap[direction].X, y-directionMap[direction].Y)
		state.Grid.Set(x, y, &previous)
		state.Grid.Set(x-directionMap[direction].X, y-directionMap[direction].Y, &current)
		if previous == robotTile {
			state.Position.X += directionMap[direction].X
			state.Position.Y += directionMap[direction].Y
		}
		return previous != robotTile
	}
	return false
}

func (state *Day15State) Move(direction string, bigBoxMode bool) {
	// from the robot's current position look in the target direction for an empty space. If we find it, shuffle as we go back.
	state.MoveRecurse(direction, state.Position.X, state.Position.Y, bigBoxMode)
}

func (state *Day15State) CalculateGPS() int {
	total := 0
	for y := 0; y < state.Grid.Height(); y++ {
		for x := 0; x < state.Grid.Width(); x++ {
			if *state.Grid.Get(x, y) == boxTile || *state.Grid.Get(x, y) == bigBoxLeft {
				total += 100*y + x
			}
		}
	}
	return total
}
func (state *Day15State) ExpandMap() {
	newGrid := containers.NewGrid[string](state.Grid.Width()*2, state.Grid.Height())
	for y := 0; y < state.Grid.Height(); y++ {
		for x := 0; x < state.Grid.Width(); x++ {
			expansion := mapExpansions[*state.Grid.Get(x, y)]
			newGrid.Set(x*2, y, &expansion[0])
			newGrid.Set(x*2+1, y, &expansion[1])
		}
	}
	state.Grid = newGrid
}

func parseDay15State(lines []string) Day15State {
	state := Day15State{}

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			state.Grid = containers.NewStringGrid(lines[0:i])
		} else if lines[i][0] == '#' {
			continue
		}
		state.Directions += strings.TrimSpace(lines[i])
	}

	return state
}

func Day15Part1() {
	loader, err := loader.NewLoader("2024/day15.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*loader.Lines = []string{
		"##########",
		"#..O..O.O#",
		"#......O.#",
		"#.OO..O.O#",
		"#..O@..O.#",
		"#O#..O...#",
		"#O..O..O.#",
		"#.OO.O.OO#",
		"#....O...#",
		"##########",
		"",
		"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^",
		"vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v",
		"><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<",
		"<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^",
		"^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><",
		"^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^",
		">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^",
		"<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>",
		"^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>",
		"v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
	}*/
	/*loader.Lines = []string{
		"########",
		"#..O.O.#",
		"##@.O..#",
		"#...O..#",
		"#.#.O..#",
		"#...O..#",
		"#......#",
		"########",
		"",
		"<^^>>>vv<v>>v<<",
	}*/
	state := parseDay15State(loader.Lines)
	if state.Grid == nil {
		return
	}

	sx, sy := state.Grid.Find(func(value string) bool { return value == robotTile })
	state.Position = containers.Point{X: sx, Y: sy}
	for i := 0; i < len(state.Directions); i++ {
		state.Move(string(state.Directions[i]), false)
	}
	//state.Grid.PrintGrid()

	fmt.Printf("Day 15 Part 1: %d\n", state.CalculateGPS())
}

func Day15Part2() {
	loader, err := loader.NewLoader("2024/day15.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{
		"##########",
		"#..O..O.O#",
		"#......O.#",
		"#.OO..O.O#",
		"#..O@..O.#",
		"#O#..O...#",
		"#O..O..O.#",
		"#.OO.O.OO#",
		"#....O...#",
		"##########",
		"",
		"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^",
		"vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v",
		"><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<",
		"<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^",
		"^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><",
		"^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^",
		">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^",
		"<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>",
		"^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>",
		"v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
	}
	/*loader.Lines = []string{
		"########",
		"#..O.O.#",
		"##@.O..#",
		"#...O..#",
		"#.#.O..#",
		"#...O..#",
		"#......#",
		"########",
		"",
		"<^^>>>vv<v>>v<<",
	}*/
	state := parseDay15State(loader.Lines)
	if state.Grid == nil {
		return
	}
	//state.Grid.PrintGrid()
	state.ExpandMap()
	//state.Grid.PrintGrid()

	sx, sy := state.Grid.Find(func(value string) bool { return value == robotTile })
	state.Position = containers.Point{X: sx, Y: sy}
	for i := 0; i < len(state.Directions); i++ {
		fmt.Println(string(state.Directions[i]))
		state.Move(string(state.Directions[i]), true)
		if state.HasUgh() {
			fmt.Println("UGH")
		}
		state.Grid.PrintGrid()
	}

	fmt.Printf("Day 15 Part 1: %d\n", state.CalculateGPS())
}
