package aoc2025

import (
	"advent/loader"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/duckdb/duckdb-go/v2"
)

func Day9Part1() {
	loader, err := loader.NewLoader("2025/day9.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"7,1",
			"11,1",
			"11,7",
			"9,7",
			"9,5",
			"2,5",
			"2,3",
			"7,3",
		}
	*/
	type Point struct {
		X, Y int
	}
	parseLines := func(lines []string) []Point {
		result := []Point{}
		for _, l := range lines {
			parts := strings.Split(l, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			result = append(result, Point{X: x, Y: y})
		}
		return result
	}
	abs := func(i int) int {
		if i >= 0 {
			return i
		}
		return i * -1
	}
	rectArea := func(a, b Point) int {
		return (abs(a.X-b.X) + 1) * (abs(a.Y-b.Y) + 1)
	}

	maxArea := 0
	points := parseLines(loader.Lines)
	for i, a := range points {
		for j := i + 1; j < len(points); j++ {
			area := rectArea(a, points[j])
			maxArea = max(area, maxArea)
		}
	}

	fmt.Printf("Day 9 Part 1: %d\n", maxArea)
}

func Day9Part2() {
	loader, err := loader.NewLoader("2025/day9.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		loader.Lines = []string{
			"7,1",
			"11,1",
			"11,7",
			"9,7",
			"9,5",
			"2,5",
			"2,3",
			"7,3",
		}
	*/
	type Point struct {
		X, Y int
	}
	parseLines := func(lines []string) []Point {
		result := []Point{}
		for _, l := range lines {
			parts := strings.Split(l, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			result = append(result, Point{X: x, Y: y})
		}
		return result
	}
	abs := func(i int) int {
		if i >= 0 {
			return i
		}
		return i * -1
	}
	rectArea := func(a, b Point) int {
		return (abs(a.X-b.X) + 1) * (abs(a.Y-b.Y) + 1)
	}
	makePolygon := func(points []Point) string {
		pointStrings := []string{}
		for _, pt := range points {
			pointStrings = append(pointStrings, fmt.Sprintf("%d %d", pt.X, pt.Y))
		}
		pointStrings = append(pointStrings, fmt.Sprintf("%d %d", points[0].X, points[0].Y))
		return fmt.Sprintf("POLYGON (( %s ))", strings.Join(pointStrings, ", "))
	}
	rectPolygon := func(a, b Point) string {
		minX := min(a.X, b.X)
		maxX := max(a.X, b.X)
		minY := min(a.Y, b.Y)
		maxY := max(a.Y, b.Y)
		points := []Point{
			{X: minX, Y: minY},
			{X: maxX, Y: minY},
			{X: maxX, Y: maxY},
			{X: minX, Y: maxY},
		}
		return makePolygon(points)
	}
	database, err := sql.Open("duckdb", "")
	if err != nil {
		panic(err)
	}
	_, err = database.Exec("INSTALL spatial;")
	if err != nil {
		panic(err)
	}
	_, err = database.Exec("LOAD spatial;")
	if err != nil {
		panic(err)
	}
	isContained := func(poly, rect string) bool {
		var result bool
		res := database.QueryRow(fmt.Sprintf("SELECT ST_CONTAINS(ST_GeomFromText('%s'), ST_GeomFromText('%s'));", poly, rect))
		err := res.Scan(&result)
		if err != nil {
			panic(err)
		}
		return result
	}
	maxArea := 0
	points := parseLines(loader.Lines)
	bounds := makePolygon(points)

	for i, a := range points {
		for j := i + 1; j < len(points); j++ {
			rect := rectPolygon(a, points[j])
			if isContained(bounds, rect) {
				area := rectArea(a, points[j])
				maxArea = max(area, maxArea)
			}
		}
	}

	fmt.Printf("Day 9 Part 2: %d\n", maxArea)
}
