package Day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
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
	file, err := os.Open("Days/Day1/input.txt")

	if err != nil {
		panic(err)
	}

	return file
}

func sortedArray(file *os.File) ([]int, []int) {
	slice1 := []int{}
	slice2 := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		items := strings.Fields(line)

		if i, err := strconv.Atoi(items[0]); err == nil {
			slice1 = append(slice1, i)
		}

		if i, err := strconv.Atoi(items[1]); err == nil {
			slice2 = append(slice2, i)
		}
	}

	slices.Sort(slice1)
	slices.Sort(slice2)

	return slice1, slice2
}



// ----------------
// Part 1
// ----------------
func part1() {
	fmt.Println("Day 1 Part 1")

	// inital setup
	file := input()
	slice1, slice2 := sortedArray(file)


	sumDistance := 0

	for index, element := range slice1 {
		sumDistance += int(math.Abs(float64(element - slice2[index])))
	}

	fmt.Println(sumDistance)
}

// ----------------
// Part 2
// ----------------
func part2() {
	file := input()
	slice1, slice2 := sortedArray(file)
	score := 0

	// Create a hashmap for first array
	// [element] : [occurence in 1st array] [occurence in 2nd array]
	hashmap := make(map[int][2]int)

	// init the hashmap with occurence in 1st array
	for _, element := range slice1 {
		
		value, exists := hashmap[element]
		if(exists) {
			value[0]++;
			hashmap[element] = value
		} else {
			hashmap[element] = [2]int{1,0}
		}

	}

	// Go through second array
	// if value exist, up the occurence count
	for _, element := range slice2 {

		value, exists := hashmap[element]
		if(exists) {
			value[1]++;
			hashmap[element] = value
		}
	}

	// Go through hashmap
	// If the element has an occurence in the 2nd array, calc the score
	for key, value := range hashmap {
		if value[1] != 0 {
			score += key * value[1]
		}
	}

	fmt.Println(score)

}