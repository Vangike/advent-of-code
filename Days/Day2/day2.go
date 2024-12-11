package Day2

import (
	"bufio"
	"fmt"
	"math"
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

	score := 0

	for _, arr := range arrayValue {
		if(increase(arr) || decrease(arr)) {
			if(absolute(arr)) {
				score++;
			}
		}
	}

	fmt.Print(score)
}

func increase(array []int) bool {
	for index := 0; index < len(array)-1; index++ {
		if array[index] > array[index+1] {
			return false
		}
	}

	return true
}

func decrease(array []int) bool {
	for index := 0; index < len(array)-1; index++ {
		if array[index] < array[index+1] {
			return false
		}
	}

	return true
}

func absolute(array []int) bool {

	for index := 0; index < len(array)-1; index++ {
		value := math.Abs(float64(array[index] - array[index+1]))

		if (value > 3 || value == 0) {
			return false
		}
	}

	return true
}

// ----------------
// Part 2
// ----------------
func part2() {
	
}