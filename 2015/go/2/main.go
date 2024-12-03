package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Function to find the two smallest values in a slice (allowing duplicates)
func findTwoSmallest(values []int) (int, int) {
	if len(values) < 2 {
		panic("Slice must contain at least two elements")
	}

	smallest := int(^uint(0) >> 1)       // Set to max int
	secondSmallest := int(^uint(0) >> 1) // Set to max int

	for _, value := range values {
		if value < smallest {
			secondSmallest = smallest
			smallest = value
		} else if value == smallest && secondSmallest == int(^uint(0)>>1) {
			// If the value is equal to the smallest and secondSmallest hasn't been set yet
			secondSmallest = smallest
		} else if value < secondSmallest {
			secondSmallest = value
		}
	}

	// If secondSmallest is still max int, it means we didn't find a valid second smallest
	if secondSmallest == int(^uint(0)>>1) {
		secondSmallest = smallest // Allow it to be the same as smallest
	}

	return smallest, secondSmallest
}

func main() {
	file, err := os.Open("input2.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// we need the total area of all the presents
	areaSum := 0
	// for part 2 we need the length of the ribbon
	// given by the cubic volume (L * W * H)
	// and shortest distance around its sides or
	// the perimeter around its shortest face
	ribbonLength := 0
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, "x")
		length, err1 := strconv.Atoi(splitLine[0])
		width, err2 := strconv.Atoi(splitLine[1])
		height, err3 := strconv.Atoi(splitLine[2])
		check(err1)
		check(err2)
		check(err3)
		// side1 := length * width
		// side2 := width * height
		// side3 := height * length
		sides := []int{length, width, height}
		smallest1, smallest2 := findTwoSmallest(sides)
		perimeter := smallest1*2 + smallest2*2
		volume := length * width * height
		ribbonLength += perimeter + volume
		areaSum += (2 * length * width) + (2 * width * height) + (2 * height * length) + smallest1
	}
	fmt.Println(ribbonLength) // part 2 correct answer 3812909
	fmt.Println(areaSum)      // part 1 correct answer 1598415
}
