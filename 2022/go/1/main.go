package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// this is my second attempt at day 1 of 2022
// its mostly the same except uses std methods to get the answers via sorting methods instead of manual maximum calculations
// and doesnt not include extraneous statistics like the amount of elves and total calories

func main() {
	file, err := os.Open("input2.txt")
	check(err)
	defer file.Close()                // Defer closing the file until the function exits
	scanner := bufio.NewScanner(file) // scanner to read the file as text
	currentCalCount := 0
	calCounts := []int{}
	for scanner.Scan() { // read each line
		line := scanner.Text() // get each line as text
		if line == "" {
			calCounts = append(calCounts, currentCalCount)
			currentCalCount = 0
			continue
		}
		currentCalories, err := strconv.Atoi(line)
		check(err)
		currentCalCount += currentCalories
	}
	slices.Sort(calCounts) // part 1 get the highest sum
	fmt.Println(calCounts[len(calCounts)-1])
	fmt.Println(calCounts[len(calCounts)-1] + calCounts[len(calCounts)-2] + calCounts[len(calCounts)-3]) // part 2 get the sume of the last 3 in the array
}
