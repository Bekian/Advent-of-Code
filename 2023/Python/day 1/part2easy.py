# get lines
lines = []
with open('input2.txt', 'r') as file:
    for line in file:
        lines.append(line)

import re

# replace number words with number numerals
# create dictionary with names and their values to switch to
numbers = {
    "twone": "21",
    "oneight": "18",
    "eightwo": "82",
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
}

# convert each word number to numeral number (i.e. one -> 1)
newLines = []
for line in lines:
    newLine = ''
    for number in numbers:
        while number in line:
            

newLines = []
for line in lines:
    newLine = ''
    for number in numbers:
        if number in line:
            tmp = line.replace(number, numbers[number])
            print(tmp)
    if len(tmp) != 0:
        newLine += tmp
    newLines.append(newLine)



# create new lines composed of only integers
newLines2 = []
for line in newLines:
    newLine = []
    for char in line:
        if char.isdigit():
            newLine.append(char)
    # print(newLine)
    newLines2.append(newLine)
#print(len(newLines2))

# lineNumbers = []
# for item in newLines2:
#     computedLineNumber = ''
#     computedLineNumber = item[0] + item[-1]
#     lineNumbers.append(int(computedLineNumber))
# print(sum(lineNumbers))