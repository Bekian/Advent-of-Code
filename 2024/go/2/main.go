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
	"slices"
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
func getSafeReports(report []int) (safeCount int) {
	safeCount = 0
	lastItem := 0
	cpy := slices.Clone(report)
	slices.Sort(cpy)
	// check if the difference between the current item and the last item is valid
	for index, v := range report {
		diffTooHigh := math.Abs(float64(v-lastItem)) > 3
		noDiff := math.Abs(float64(v-lastItem)) == 0
		if index != 0 && (diffTooHigh ||  noDiff) {
			return 0
		}
		lastItem = v
	}
	// check if the report is increasing or decreasing
	if slices.Compare(cpy, report) == 0 {
		return 1
	}
	slices.Reverse(cpy)
	if slices.Compare(cpy, report) == 0 {
		return 1
	}

	return
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

		safeReports = getSafeReports(modifiedSubtractants)
		if safeReports > 0 {
			break
		}
	}
	return
}

func main() {
	file, err := os.Open("input2.txt")
	check(err)
	defer file.Close()
	safeCount := 0
	// newSubtractants := []int{}
	scanner := bufio.NewScanner(file)
	// totalSubtractants := []int{}
	// using the file content
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := splitInput(line)
		safeCount += problemDampener(splitLine)
	}
	fmt.Println()
	fmt.Println(safeCount)
}

// 381, 380, 370, 329 are incorrect