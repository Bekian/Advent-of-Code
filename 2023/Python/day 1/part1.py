# iterate over each line, 
# find all the numbers in each line,
# get the first and last numbers from the line
# then turn those into an integer
# then caculate the sum of all the integers
# example:
"""
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
""" 
# would result in: 
""" 
12
38f
15
77
"""
# then the result is: 
# 142
lines = []
with open('input.txt', 'r') as file:
    for line in file:
        lines.append(line)


# biggest = ""
# for i in lines:
#     if len(i) >= len(biggest):
#         biggest = i


allTheInts = []
#print(len(biggest))
for i in lines:
    # each line
    lineInts = []
    for j in i:
        # each item
        try: 
            # get each int in a line
            int(j) == type(int)
            lineInts.append(j)
        except:
            continue
    # then put the line into another list
    allTheInts.append(lineInts)

print(allTheInts)

lineNumbers = []
for i in allTheInts:
    computedLineNumber = ''
    # each line
    computedLineNumber = i[0] + i[-1]
    lineNumbers.append(int(computedLineNumber))

# print(lineNumbers[-1])
# print(sum(lineNumbers))

# my current idea is to get all the ints in a line
# then compute what the pair of ints is supposed to be from that