package Day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run(part string) {
	switch part {
	case "1":
		part1()
	case "2":
		part2()
	default:
		fmt.Printf("Part %s not implmented", part)
	}
}

// ----------------
// Supporting Functions
// ----------------

// Run input file
func input() (*os.File) {
	file, err := os.Open("Days/Day2/input.txt")

	if err != nil {
		panic(err)
	}

	return file
}

func twoDArray(file *os.File) ([][]int) {
	array := [][]int{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		items := strings.Fields(line)

		valueArr := []int{}

		for _, str := range items {
			if value, err := strconv.Atoi(str); err == nil {
				valueArr = append(valueArr, value)
			}
		}

		array = append(array, valueArr)
	}

	return array
}

// ----------------
// Part 1
// ----------------
func part1() {
	file := input()
	arrayValue := twoDArray(file)

	fmt.Print(arrayValue)
}

// ----------------
// Part 2
// ----------------
func part2() {
	
}