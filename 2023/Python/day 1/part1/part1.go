package main

// adding regexp2 to have lookbehind expression on line 55
import (
	"fmt"
	"os"
	"regexp2"
	"strings"
)

func print(input any) {
	fmt.Print(input)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var mergedNumberWords = map[string]string{
	"twone":   "21",
	"oneight": "18",
	"eightwo": "82",
}

var numberWords = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func removeNonDigits(input string) string {
	// Define a regular expression pattern to match non-digit characters
	reg := regexp2.MustCompile("[^0-9]")

	// Use the regular expression to replace non-digit characters with an empty string
	return reg.ReplaceAllString(input, "")

}

func replaceText(strArray []string, old, new string) {
	replacer := strings.NewReplacer(old, new)
	for i, str := range strArray {
		strArray[i] = replacer.Replace(str)
	}
}

func findLast(search string) (chr string) {
	reg := regexp2.MustCompile("(?:\\d)(?!.*\\d)")
	chr = reg.ReplaceAllString(search, "")
	return
}

func main() {
	// get test data
	testData, err := os.ReadFile("correct.txt")
	check(err)
	testStr := string(testData)
	testLines := strings.Fields(testStr)

	// input, convert to string slice
	dat, err := os.ReadFile("input2.txt")
	check(err)
	strDat := string(dat)
	lines := strings.Fields(strDat)
	print(findLast(lines[0]))

	print(lines)
	print("\n")
	print(testLines)
}
