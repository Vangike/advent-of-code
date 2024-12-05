package main

import (
	"aoc/Days/Day1"
	"aoc/Days/Day2"
	"fmt"
	"os"
	"reflect"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Specify a day to run - i.e. main.go Day1")
		return
	}

	dayName := os.Args[1]

	funcMap := map[string]interface{} {
		"Day1": Day1.Run,
		"Day2": Day2.Run,
	}

	if fn, exists := funcMap[dayName]; exists {
		reflect.ValueOf(fn).Call(nil)
	} else {
		fmt.Printf("Function %s not found\n", dayName)
	}
}