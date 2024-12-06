package main

import (
	"aoc/Days/Day1"
	"aoc/Days/Day2"
	"fmt"
	"os"
	"reflect"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: go run main.go <day> <part>")
		fmt.Printf("Example: go run main.go Day1 2")
		return
	}

	dayName := os.Args[1]
	part := os.Args[2]

	funcMap := map[string]interface{} {
		"Day1": Day1.Run,
		"Day2": Day2.Run,
	}

	if fn, exists := funcMap[dayName]; exists {
		reflectArgs := []reflect.Value{reflect.ValueOf(part)}
		reflect.ValueOf(fn).Call(reflectArgs)
	} else {
		fmt.Printf("Function %s not found\n", dayName)
	}
}