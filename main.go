package main

import (
	"advent/aoc2023"
	"advent/aoc2024"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var year = flag.Int("year", 2023, "Year to run")
var day = flag.Int("day", 1, "Day to run")
var file = flag.String("file", "", "File to run")

type Dispatcher interface {
	Run(day int)
}

func main() {
	fmt.Println(os.Args)
	flag.Parse()

	var dispatcher Dispatcher
	if file != nil && *file != "" {
		fileparseregex := regexp.MustCompile(`aoc(\d+).+day(\d+).go`)
		fileparse := fileparseregex.FindAllStringSubmatch(*file, -1)
		*year, _ = strconv.Atoi(fileparse[0][1])
		*day, _ = strconv.Atoi(fileparse[0][2])
	}
	if *year == 2023 {
		dispatcher = &aoc2023.Aoc23Dispatcher{}
	} else if *year == 2024 {
		dispatcher = &aoc2024.Aoc24Dispatcher{}
	}

	dispatcher.Run(*day)
}
