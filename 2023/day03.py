# Read the question here: https://adventofcode.com/2023/day/3

file_path = 'input/day03.txt'
file_prova = 'input/prova.txt'

def checkNumber(rows, start, finish):
    # I need to check from rows[i-1][j-1] to rows[i+1][j:k]
    for i in range(len(rows)):
        for j in range(start, finish):
            if (not rows[i][j].isdigit()) & (not rows[i][j] == '.'):
                return True
    return False

def part1():
    with open(file_path) as f:
        rows = f.readlines()
    s = 0
    for i in range(len(rows)):
        j = 0
        while j < len(rows[i]):
            if not rows[i][j].isdigit():
                j+=1
                continue
            for k in range(j+1, len(rows[i])):
                if not rows[i][j:k].isdecimal():
                    # Number Found: it is in position rows[i][j:k-1]
                    number = int(rows[i][j:k-1])
                    if checkNumber(rows[max(0,i-1):min(i+2, len(rows))], max(0,j-1), min(len(rows[i]),k)):
                        s += number
                    j = k
                    break 
                if k == len(rows[i])-1:
                    number = int(rows[i][j:k])
                    if checkNumber(rows[max(0,i-1):min(i+2, len(rows))], max(0,j-1), min(len(rows[i]),k)):
                        s += number
                    j = k
                    break
    return s

def part2():
    with open(file_prova) as f:
        rows = f.readlines()
    s = 0
    return s

print(f'Part 1: {part1()}')
print(f'Part 2: {part2()}')