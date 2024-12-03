package main

import (
	"bufio"
	"fmt"
	"os"
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
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		switch line {
		case "A Y":
			sum += 2
			continue
		case "B Z":
			sum += 3
			continue
		case "C X":
			sum += 1
			continue
		case "A X":
			sum += 4
			continue
		case "B Y":
			sum += 5
			continue
		case "C Z":
			sum += 6
			continue
		case "A Z":
			sum += 9
			continue
		case "B X":
			sum += 7
			continue
		case "C Y":
			sum += 8
			continue
		default:
			fmt.Println("parsing error: ", line)
		}

	}
	fmt.Println(sum)
	// this should(tm) work but it doesnt, it results in a score of 16004, but it should be 14859
	// im guessing the scores assignments are incorrect
}
