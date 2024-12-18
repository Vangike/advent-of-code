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
// Part 2 UNRESOLVED
// ----------------
func part2() {
	// file := input()
	// arrayValue := twoDArray(file)

	// score := 0

	// failSafe := new(int)
	// *failSafe = 1

	// for _, arr := range arrayValue {

	// 	incArr, incResult := increaseF(arr, failSafe) 

	// 	fmt.Print("INCREASE \n")
	// 	if(incResult) {
	// 		fmt.Print(incArr)
	// 		fmt.Print(*failSafe)

	// 		if(absoluteF(incArr, failSafe, 0) || absoluteF(incArr, failSafe, 1) || absoluteF(incArr, failSafe, 2)) {
	// 			score++
	// 			*failSafe = 1
	// 			fmt.Print("SCORE \n")
	// 			fmt.Print(incArr)
	// 		}
	// 	}

	// 	*failSafe = 1

	// 	// decArr, decResult := decreaseF(arr, failSafe)

	// 	// fmt.Print("\n DECREASE \n")
	// 	// if(decResult) {
	// 	// 	fmt.Print(incArr)
	// 	// 	fmt.Print(*failSafe)

	// 	// 	if(absoluteF(decArr, failSafe)) {
	// 	// 		score++
	// 	// 	}
	// 	// }

	// 	// *failSafe = 1

	// 	break;
	// }


}

func increaseF(arr []int, failsafe *int) ([]int, bool) {

	arrCopy := make ([]int, len(arr))
	copy(arrCopy, arr)

	for in := 0; in < len(arrCopy)-1; {
		if arrCopy[in] > arrCopy[in+1] {
			if(*failsafe != 0) {
				arrCopy = append(arrCopy[:in], arrCopy[in+1:]...)
				*failsafe = 0
			} else {
				return arrCopy, false
			}
		} else {
			in++
		}
	}
	
	return arrCopy, true
}

func decreaseF(arr []int, failsafe *int) ([]int, bool) {

	arrCopy := make ([]int, len(arr))
	copy(arrCopy, arr)

	for in := 0; in < len(arrCopy)-1; {
		if arrCopy[in] < arrCopy[in+1] {
			if(*failsafe != 0) {
				arrCopy = append(arrCopy[:in], arrCopy[in+1:]...)
				*failsafe = 0
			} else {
				return arrCopy, false
			}
		} else {
			in++
		}
	}
	
	return arrCopy, true
}

func absoluteF(arr []int, failsafe int, offset int) bool {

	array := make ([]int, len(arr))
	copy(array, arr)

	fmt.Print(array)
	fmt.Print(failsafe)

	fmt.Print("\n")

	for in := 0; in < len(array)-1; {
		value := math.Abs(float64(array[in] - array[in+1]))

		fmt.Printf(" %d %f \n", in, value)

		if (value > 3 || value == 0) {
			if(failsafe != 0) {
				fmt.Printf("Removed: %d\n", array[in])
				array = append(array[:in+offset], array[in+1+offset:]...)
				in--
				failsafe = 0
			} else {
				fmt.Print("FAILED")
				fmt.Print(array)
				return false
			}
		} else {
			in++
		}
	}
	
	// fmt.Print(array)
	return true
}