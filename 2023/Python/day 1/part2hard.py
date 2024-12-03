# this is the more complex solution of the second part of day 1 (wip)
# it manually searches for words, collecting their indexes
# and then manually searches for the numerals, again recording the indexes 
# then using that information builds new lines of numerals in the correct order
# then recomputes the algorithim in the first part

# second part discovered that there are also numbers
# in the form of words! so i need to add those too
lines = []
with open('input.txt', 'r') as file:
    for line in file:
        lines.append(line)

mappedNumber = { "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9 }
newLines = []

# search each line
for line in lines:
    # create a dict for the ints in the row
    # index: item
    lineDict = {}
    # search for word ints
    for number in mappedNumber:
        # for each number in the map
        # check if the current map number word is in the line
        if number in line:
            # add index and item
            lineDict.update({line.find(number): mappedNumber[number]})
    if len(lineDict) == 0:
        continue
    newLines.append(lineDict)

# search for int literals
newLines2 = []
for line in newLines:
    lineDict = {}
    for item in line:
        try: 
            # get each int in a line
            int(item) == type(int)
        except:
            continue
        else:
            lineDict.update({line.find(item): int(item)})
    newLines2.append(lineDict)

print(lines[0])
# print(newLines[0])
# print(newLines2[0])

# merge the two dicts
newLines[0].update(newLines2[0])
print(newLines[0])

for w in sorted(newLines[0], key=newLines[0].get):
    print(w, newLines[0][w])

# assemble the lines
assembledLines = []
for line in newLines:
    # create a new line
    newLine = []
    # then find the right item from the other line to add
    for item in line:
        pass
        #print(item)


# print(assembledLines)