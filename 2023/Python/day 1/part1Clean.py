lines = []
with open('input.txt', 'r') as file:
    for line in file:
        lines.append(line)
allTheInts = []
for i in lines:
    lineInts = []
    for j in i:
        try: 
            int(j) == type(int)
            lineInts.append(j)
        except:
            continue
    allTheInts.append(lineInts)
lineNumbers = []
for i in allTheInts:
    computedLineNumber = ''
    computedLineNumber = i[0] + i[-1]
    lineNumbers.append(int(computedLineNumber))
print(sum(lineNumbers))