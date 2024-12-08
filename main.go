package main

import (
	"advent/aoc2015"
	"advent/aoc2016"
	"advent/aoc2017"
	"advent/aoc2018"
	"advent/aoc2019"
	"advent/aoc2020"
	"advent/aoc2021"
	"advent/aoc2022"
	"advent/aoc2023"
	"advent/aoc2024"
	"bytes"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"text/template"
)

var year = flag.Int("year", 2023, "Year to run")
var day = flag.Int("day", 1, "Day to run")
var file = flag.String("file", "", "File to run")
var generate = flag.Bool("generate", false, "Generate templates for year")

type Dispatcher interface {
	Run(day int)
}

var codeTemplate = `
package aoc{{.Year}}

import (
	"advent/loader"
	"fmt"
)

func Day{{.Day}}Part1() {
	loader, err := loader.NewLoader("{{.Year}}/day{{.Day}}")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day {{.Day}} Part 1: %d\n", 0)
}

func Day{{.Day}}Part2() {
	loader, err := loader.NewLoader("{{.Year}}/day{{.Day}}")
	if err != nil {
		fmt.Println(err)
		return
	}
	loader.Lines = []string{}

	fmt.Printf("Day {{.Day}} Part 2: %d\n", 0)
}
`
var dispatcherTemplate = `
package aoc{{.Year}}

import (
	"fmt"
	"time"
)

type Aoc{{.ShortYear}}Dispatcher struct {
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func (d *Aoc{{.ShortYear}}Dispatcher) Run(day int) {
	defer timer(fmt.Sprintf("Day %d", day))()
	switch day {
	case 1:
		Day1Part1()
		Day1Part2()
	case 2:
		Day2Part1()
		Day2Part2()
	case 3:
		Day3Part1()
		Day3Part2()
	case 4:
		Day4Part1()
		Day4Part2()
	case 5:
		Day5Part1()
		Day5Part2()
	case 6:
		Day6Part1()
		Day6Part2()
	case 7:
		Day7Part1()
		Day7Part2()
	case 8:
		Day8Part1()
		Day8Part2()
	case 9:
		Day9Part1()
		Day9Part2()
	case 10:
		Day10Part1()
		Day10Part2()
	case 11:
		Day11Part1()
		Day11Part2()
	case 12:
		Day12Part1()
		Day12Part2()
	case 13:
		Day13Part1()
		Day13Part2()
	case 14:
		Day14Part1()
		Day14Part2()
	case 15:
		Day15Part1()
		Day15Part2()
	case 16:
		Day16Part1()
		Day16Part2()
	case 17:
		Day17Part1()
		Day17Part2()
	case 18:
		Day18Part1()
		Day18Part2()
	case 19:
		Day19Part1()
		Day19Part2()
	case 20:
		Day20Part1()
		Day20Part2()
	case 21:
		Day21Part1()
		Day21Part2()
	case 22:
		Day22Part1()
		Day22Part2()
	case 23:
		Day23Part1()
		Day23Part2()
	case 24:
		Day24Part1()
		Day24Part2()
	case 25:
		Day25Part1()
		Day25Part2()
	default:
		fmt.Println("Day not implemented")
		return
	}
}
`

func main() {
	fmt.Println(os.Args)
	flag.Parse()

	if *generate {
		os.Mkdir(fmt.Sprintf("aoc%d", *year), 0755)

		data := map[string]interface{}{
			"Year":      strconv.Itoa(*year),
			"ShortYear": strconv.Itoa(*year)[2:],
		}

		buf := bytes.Buffer{}
		dispatcherTmpl, err := template.New("dispatcher").Parse(dispatcherTemplate)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = dispatcherTmpl.Execute(&buf, data)
		if err != nil {
			fmt.Println(err)
			return
		}
		os.WriteFile(fmt.Sprintf("aoc%d/dispatcher.go", *year), buf.Bytes(), 0644)

		codeTmpl := template.New("code")
		codeTmpl.Parse(codeTemplate)

		for i := 1; i <= 25; i++ {
			fn := fmt.Sprintf("aoc%d/day%d.go", *year, i)
			if _, err := os.Stat(fn); err != nil {
				data["Day"] = strconv.Itoa(i)
				buf = bytes.Buffer{}
				err = codeTmpl.Execute(&buf, data)
				if err != nil {
					fmt.Println(err)
					return
				}
				os.WriteFile(fmt.Sprintf("aoc%d/day%d.go", *year, i), buf.Bytes(), 0644)
			}

			// create input file
			os.Mkdir(fmt.Sprintf("data/%d", *year), 0755)
			fn = fmt.Sprintf("data/%d/day%d.txt", *year, i)
			if _, err := os.Stat(fn); err != nil {
				os.WriteFile(fn, []byte{}, 0644)
			}
		}

		return
	}

	var dispatcher Dispatcher
	if file != nil && *file != "" {
		fileparseregex := regexp.MustCompile(`aoc(\d+).+day(\d+).go`)
		fileparse := fileparseregex.FindAllStringSubmatch(*file, -1)
		*year, _ = strconv.Atoi(fileparse[0][1])
		*day, _ = strconv.Atoi(fileparse[0][2])
	}
	if *year == 2015 {
		dispatcher = &aoc2015.Aoc15Dispatcher{}
	} else if *year == 2016 {
		dispatcher = &aoc2016.Aoc16Dispatcher{}
	} else if *year == 2017 {
		dispatcher = &aoc2017.Aoc17Dispatcher{}
	} else if *year == 2018 {
		dispatcher = &aoc2018.Aoc18Dispatcher{}
	} else if *year == 2019 {
		dispatcher = &aoc2019.Aoc19Dispatcher{}
	} else if *year == 2020 {
		dispatcher = &aoc2020.Aoc20Dispatcher{}
	} else if *year == 2021 {
		dispatcher = &aoc2021.Aoc21Dispatcher{}
	} else if *year == 2022 {
		dispatcher = &aoc2022.Aoc22Dispatcher{}
	} else if *year == 2023 {
		dispatcher = &aoc2023.Aoc23Dispatcher{}
	} else if *year == 2024 {
		dispatcher = &aoc2024.Aoc24Dispatcher{}
	}

	dispatcher.Run(*day)
}
