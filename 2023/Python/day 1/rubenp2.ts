let result = 0

const table = {
    one: 1,
    two: 2,
    three: 3,
    four: 4,
    five: 5,
    six: 6,
    seven: 7,
    eight: 8,
    nine: 9
}

const keys = Object.keys(table)

function convert(s: string): string {
    return s in table ? Reflect.get(table, s).toString() : s
}

function findResults(line: string): string[] {
    const maxLen = 5 
    const results = new Array<string>()

    for (let i = 0, len = line.length;i < len;i++) {
        const char = line[i]
        if (/\d/.test(char)) {
            results.push(char)
            continue
        }

        for (let x = 0;x <= maxLen;x++) {
            const match = line.slice(i, i + x)
            if (match in table) results.push(match)
        }
    }

    return results
}

for (const line of adv1) {
    const numbers = findResults(line)
    const n = Number(convert(numbers[0]) + convert(numbers.at(-1)!))
    result += n
}

console.log(result)