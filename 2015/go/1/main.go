// 2015 day 1
// the first part of this problem involves counting pairs of parenthesis
// my plan for part 1 is to use a count value to add 1 or subtract 1 based on the left or right parenthesis
// the second part of this problem involves finding the first case where the current floor is -1
// so i just exited when the count was -1 and added one to get the answer
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
	floorCount := 0
	firstOut := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for i, v := range line {
			item := string(v)
			if item == "(" {
				floorCount += 1
			} else if item == ")" {
				floorCount -= 1
			}
			if floorCount == -1 {
				firstOut = i
				break
			}
		}
	}
	fmt.Println(floorCount, firstOut)
	// part 2, 1794 too low, 1795 is correct
}
