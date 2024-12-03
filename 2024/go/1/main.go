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

func main() {
	file, err := os.Open("input2.txt")
	check(err)
	defer file.Close()
	column1 := []int{}
	column2 := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, "   ")
		value1, err1 := strconv.Atoi(splitLine[0])
		value2, err2 := strconv.Atoi(splitLine[1])
		check(err1)
		check(err2)
		column1 = append(column1, value1)
		column2 = append(column2, value2)
	}
	slices.Sort(column1)
	slices.Sort(column2)
	// part 1
	var sum float64
	for i := 0; i < len(column1); i++ {
		v1 := column1[i]
		v2 := column2[i]
		sum += math.Abs(float64(v1 - v2))
	}
	// fmt.Println(int(sum)) // correct answer for part 1: 2430334
	// part 2
	similarityScore := 0
	similarityCount := make(map[int]int)
	for index, item := range column1 {
		// check how many times item from col1 appears in col 2
		for _, col2 := range column2 {
			if col2 == item {
				similarityCount[index] += 1
			}
		}
		similarityScore += similarityCount[index] * item
	}
	fmt.Println(similarityScore)
}
