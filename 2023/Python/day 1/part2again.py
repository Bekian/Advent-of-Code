# third attempt using tdd 
# correct sum is 706
lines = []
with open('input2.txt', 'r') as file:
    for line in file:
        lines.append(line)

correctOuput1 = []
with open('correct.txt', 'r') as file:
    for line in file:
        correctOuput1.append(line)

correctOuput2 = []
with open('correct2.txt', 'r') as file:
    for line in file:
        correctOuput2.append(line)

# replace number words with number numerals
# create dictionary with names and their values to switch to
numbers = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
    "twone": "21",
    "oneight": "18",
    "fiveight": "58"
}

def compareStrs(str1, str2):
    if str1 != str2:
        return False
    return True

def compareLists(list1, list2):
    if len(list1) != len(list2):
        print("ERROR: LIST LENGTH MISMATCH")
        return
    errorTally = 0
    for i in range(len(list1)):
        if list1[i] != list2[i]:
            errorTally += 1
        else:
            continue
    if errorTally == 0:
        return "Lists match"
    return errorTally


def convert(numberString):
    return numbers[numberString]


# convert each word number to numeral number (i.e. one -> 1)
newLines = []
for i in range(len(lines) - 1):
    newLine = []
    # find the numbers[i] to attempt to replace it in the str
    for item in lines[i]:
        

# for i in newLines:
#     print(''.join(i))

# print(newLines)
# print(len(newLines), len(correctOuput1))
# print(compareLists(correctOuput1, newLines))
# print(lines[0], ''.join(newLines[0]))
