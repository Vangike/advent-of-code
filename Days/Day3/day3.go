package Day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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

// INPUT FUNCTIONS
func input() (string) {
	content, err := os.ReadFile("Days/Day3/input.txt")

	if err != nil {
		panic(err)
	}

	fileString := string(content)
	return fileString
}


// PART 1
func part1() {
	score := 0

	// Find all mul(x, x) in string
	regex := `mul\((\d+),(\d+)\)`
	intputString := input()

	re := regexp.MustCompile(regex)

	allMatches := re.FindAllString(intputString, -1)

	// Parse mul(x, x)
	for _, match := range allMatches {
		values := re.FindStringSubmatch(match)

		value1, err1 := strconv.Atoi(values[1])
		value2, err2 := strconv.Atoi(values[2])

		if(err1 != nil || err2 != nil) {
			return
		}

		score += value1 * value2
	}

	fmt.Print(score)
	
}

func part2() {
	
}
