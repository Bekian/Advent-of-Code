package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getNums(line string) (value int) {
	fmt.Println(line)
	// values are an array as an (index, value) pair
	firstValue := []int{len(line) - 1, 0}
	lastValue := []int{0, 0}
	for index, item := range line {
		if unicode.IsDigit(item) {
			if index < firstValue[0] {
				firstValue[0] = index
				firstValue[1] = int(item)
				//fmt.Printf("%c\n", item)
			}
			if index > lastValue[0] {
				lastValue[0] = index
				lastValue[1] = int(item)
				//fmt.Printf("%c\n", item)
			}
		}
	}
	// something abt this is bonked
	value, err := strconv.Atoi(fmt.Sprintf("%c%c", firstValue[1], lastValue[1]))
	check(err)
	return
}

func main() {
	// Open the file
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Create a new scanner
	scanner := bufio.NewScanner(file)

	// Scan line by line
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		val := getNums(line)
		fmt.Println(val)
		total += val
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Close the file
	defer file.Close()
}
