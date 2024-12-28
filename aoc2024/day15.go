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

func (state *Day15State) IsBoxTile(x, y int) bool {
	value := *state.Grid.Get(x, y)
	return value == boxTile || value == bigBoxLeft || value == bigBoxRight
}
func (state *Day15State) CountBoxes() int {
	total := 0
	for y := 0; y < state.Grid.Height(); y++ {
		for x := 0; x < state.Grid.Width(); x++ {
			cell := *state.Grid.Get(x, y)
			if cell == boxTile || cell == bigBoxLeft {
				total++
			}
		}
	}
	return total
}
func (state *Day15State) CountWallTiles() int {
	total := 0
	for y := 0; y < state.Grid.Height(); y++ {
		for x := 0; x < state.Grid.Width(); x++ {
			if *state.Grid.Get(x, y) == wallTile {
				total++
			}
		}
	}
	return total
}
func (state *Day15State) Validate() bool {
	for y := 0; y < state.Grid.Height(); y++ {
		for x := 0; x < state.Grid.Width(); x++ {
			cell := *state.Grid.Get(x, y)
			if cell == boxTile || cell == bigBoxLeft {
				if *state.Grid.Get(x+1, y) != bigBoxRight {
					return false
				}
			}
		}
	}
	return true
}

func (state *Day15State) CalculateGPS() int {
	total := 0
	for y := 0; y < state.Grid.Height(); y++ {
		for x := 0; x < state.Grid.Width(); x++ {
			cell := *state.Grid.Get(x, y)
			if cell == boxTile || cell == bigBoxLeft {
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

func (state *Day15State) MoveHorizontal(direction string, start containers.Point) bool {
	if *state.Grid.Get(start.X, start.Y) == wallTile {
		return false
	}
	if *state.Grid.Get(start.X, start.Y) == emptyTile {
		return true
	}

	if state.MoveHorizontal(direction, containers.Point{X: start.X + directionMap[direction].X, Y: start.Y}) {
		state.Grid.Swap(start.X, start.Y, start.X+directionMap[direction].X, start.Y)
		return true
	}

	return false
}

func (state *Day15State) MoveVertical(direction string, startX, endX, Y int) bool {
	if *state.Grid.Get(startX, Y) == bigBoxRight {
		startX--
	}
	if *state.Grid.Get(endX, Y) == bigBoxLeft {
		endX++
	}
	for *state.Grid.Get(startX, Y) == emptyTile {
		startX++
	}
	for *state.Grid.Get(endX, Y) == emptyTile {
		endX--
	}

	// Look at next row to see if there's a wall in the way
	canMove := true
	checkNextRow := false
	for x := startX; x <= endX; x++ {
		nextIsBox := state.IsBoxTile(x, Y+directionMap[direction].Y)
		currentIsBox := state.IsBoxTile(x, Y) || *state.Grid.Get(x, Y) == robotTile
		nextIsFree := *state.Grid.Get(x, Y+directionMap[direction].Y) == emptyTile
		nextIsWall := *state.Grid.Get(x, Y+directionMap[direction].Y) == wallTile
		if currentIsBox && nextIsWall {
			return false
		}
		if nextIsBox && currentIsBox {
			checkNextRow = true
			canMove = false
		}
		if nextIsFree && currentIsBox {
			canMove = canMove && true
		}
	}

	if checkNextRow {
		canMove = state.MoveVertical(direction, startX, endX, Y+directionMap[direction].Y)
	}

	if canMove {
		for x := startX; x <= endX; x++ {
			previousY := Y + directionMap[direction].Y
			if state.IsBoxTile(x, Y) || *state.Grid.Get(x, Y) == robotTile {
				state.Grid.Swap(x, previousY, x, Y)
			}
		}
		return true
	}

	return false
}

func (state *Day15State) Move(direction string) {
	moved := false
	if direction == "<" || direction == ">" {
		moved = state.MoveHorizontal(direction, state.Position)
	} else if direction == "^" || direction == "v" {
		moved = state.MoveVertical(direction, state.Position.X, state.Position.X, state.Position.Y)
	}
	if moved {
		state.Position = containers.Point{X: state.Position.X + directionMap[direction].X, Y: state.Position.Y + directionMap[direction].Y}
	}
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
		//state.Grid.PrintGrid()
		//fmt.Println(string(state.Directions[i]))
		state.Move(string(state.Directions[i]))
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
		"#######",
		"#...#.#",
		"#.....#",
		"#..OO@#",
		"#..O..#",
		"#.....#",
		"#######",
		"",
		"<vv<<^^<<^^",
	}*/
	loader.Lines = []string{
		"##################",
		"####....##....####",
		"##..[][]..[][]..##",
		"##...[]....[]...##",
		"##....[]..[][]..##",
		"##..[][]..[][]..##",
		"##....[]...[]...##",
		"##....[]..[]....##",
		"##.....[][].....##",
		"##......[]......##",
		"##.......@......##",
		"##################",
		"",
		"^",
	}
	state := parseDay15State(loader.Lines)
	if state.Grid == nil {
		return
	}
	//state.Grid.PrintGrid()
	//state.ExpandMap()
	//state.Grid.PrintGrid()

	startCount := state.CountBoxes()
	startWalls := state.CountWallTiles()

	sx, sy := state.Grid.Find(func(value string) bool { return value == robotTile })
	state.Position = containers.Point{X: sx, Y: sy}
	for i := 0; i < len(state.Directions); i++ {
		//state.Grid.PrintGrid()
		//fmt.Println(string(state.Directions[i]))
		state.Move(string(state.Directions[i]))
		count := state.CountBoxes()
		if count != startCount {
			panic(fmt.Sprintf("Box count changed from %d to %d at %d", startCount, count, i))
		}
		count = state.CountWallTiles()
		if count != startWalls {
			panic(fmt.Sprintf("Wall count changed from %d to %d at %d", startWalls, count, i))
		}
		if !state.Validate() {
			panic(fmt.Sprintf("Invalid state at %d", i))
		}
	}
	state.Grid.PrintGrid()
	endCount := state.CountBoxes()

	fmt.Printf("Day 15 Part 2: %d, (start: %d, end %d)\n", state.CalculateGPS(), startCount, endCount)
}
