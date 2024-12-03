// 2015 day 3
// my current plan is to make some 2D object, ideally a 2D map or an encoded map
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
	file, err := os.Open("input1.txt")
	check(err)
	defer file.Close()
	map2D := make(map[int]int)
	xPos := 0
	yPos := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for v := range line {
			item := string(v)
			if item == ">" {
				xPos += 1
			} else if item == "<" {
				xPos -= 1
			} else if item == "^" {
				yPos += 1
			} else if item == "v" {
				yPos -= 1
			}
		}
	}
	fmt.Println(xPos, yPos, map2D)
}
