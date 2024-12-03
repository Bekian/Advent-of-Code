// 2024 day 2
// my current plan is to get an array of subtractants between every item
// first check if there are both positive and negative ones in the array
// then check the absolute value of the subtractants is greater than 3
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// splits an input line into a slice of numbers
func splitInput(line string) []int {
	ints := []int{}
	splitLine := strings.Split(line, " ")
	for _, v := range splitLine {
		// convert each value into an integer
		currItem, err := strconv.Atoi(v)
		check(err)
		ints = append(ints, currItem)
	}
	return ints
}

// get the subtractants between each item
func getSubtractants(input []int) []int {
	lastItem := 0
	subtractants := []int{}
	for index, v := range input {
		if index != 0 {
			subtractants = append(subtractants, v-lastItem)
		}
		lastItem = v
	}
	return subtractants
}

// calculate the amount of safe reports in a line
func getSafeReports(subtractants []int) int {
	safeCount := 0
	positiveFlag := false
	negativeFlag := false
	zeroFlag := false
	for i, v := range subtractants {
		if math.Abs(float64(v)) > 3 { // omit subtractants greater than 3
			break
		}
		// define initial sentiment
		if v > 0 {
			positiveFlag = true
		} else if v < 0 {
			negativeFlag = true
		} else {
			zeroFlag = true
		}

		if positiveFlag && negativeFlag || zeroFlag {
			break
		} else if positiveFlag != negativeFlag {
			if i == len(subtractants)-1 { // if the list is still valid than catch it
				// fmt.Println(subtractants)
				safeCount += 1
			}
		} else {
			break
		}
	}
	return safeCount
}

// modified safeReports function
// this function removes an item from the subtractants
// then calls the getSafeReports function in succession until a safe report is found
// unless theres no possible combination, then returns 0
func problemDampener(subtractants []int) (safeReports int) {
	safeReports = 0
	for i := range subtractants {
		// Create a new slice excluding the element at index i
		modifiedSubtractants := append(subtractants[:i], subtractants[i+1:]...)

		// Check if the modified subtractants result in a safe report
		if getSafeReports(modifiedSubtractants) > 0 {
			// fmt.Println(modifiedSubtractants) // check if the subtractants are actually valid
			safeReports = 1
			return
		}
	}
	return
}

func main() {
	file, err := os.Open("input1.txt")
	check(err)
	defer file.Close()
	safeCount := 0
	scanner := bufio.NewScanner(file)
	// using the file content
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := splitInput(line)
		subtractants := getSubtractants(splitLine)
		safeCount += problemDampener(subtractants)
		// safeCount += getSafeReports(subtractants)
	}
	fmt.Println()
	fmt.Println(safeCount)
}
