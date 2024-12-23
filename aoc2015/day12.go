package aoc2015

import (
	"advent/loader"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strings"
)

func recursiveSum(value interface{}) float64 {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Map:
		sum := 0.0
		for _, key := range v.MapKeys() {
			sum += recursiveSum(v.MapIndex(key).Interface())
		}
		return sum
	case reflect.Int:
		return float64(v.Int())
	case reflect.Float64:
		return v.Float()
	case reflect.Slice:
		fallthrough
	case reflect.Array:
		// if the slice type is int, sum the values, if interface{}, recurse
		sum := 0.0
		for i := 0; i < v.Len(); i++ {
			sum += recursiveSum(v.Index(i).Interface())
		}
		return sum
	case reflect.String:
		return 0
	default:
		fmt.Printf("not implemented type: %v\n", v.Kind())
	}
	return 0
}

func recursiveSum2(value interface{}) float64 {
	switch value := value.(type) {
	case map[string]interface{}:
		sum := 0.0
		for _, v := range value {
			next := recursiveSum2(v)
			if math.IsNaN(next) {
				// The recursive call had a string "red"
				return 0
			}
			sum += next
		}
		return sum
	case []interface{}:
		sum := 0.0
		for _, v := range value {
			next := recursiveSum2(v)
			if math.IsNaN(next) {
				// The recursive call had a string "red"
				continue
			}
			sum += next
		}
		return sum
	case float64:
		return value
	case string:
		if value == "red" {
			return math.NaN()
		}
		return 0
	default:
		fmt.Printf("not implemented type: %v\n", reflect.TypeOf(value))
	}
	return 0
}

func Day12Part1() {
	loader, err := loader.NewLoader("2015/day12.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	data := []byte(strings.Join(loader.Lines, "\n"))

	var value interface{}
	err = json.Unmarshal(data, &value)
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := int(recursiveSum(value))

	fmt.Printf("Day 12 Part 1: %d\n", sum)
}

func Day12Part2() {
	loader, err := loader.NewLoader("2015/day12.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	data := []byte(strings.Join(loader.Lines, "\n"))

	var value interface{}
	err = json.Unmarshal(data, &value)
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := int(recursiveSum2(value))

	fmt.Printf("Day 12 Part 2: %d\n", sum)
}
