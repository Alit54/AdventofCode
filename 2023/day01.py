### First AOC Day in Python!
# Read the question here: https://adventofcode.com/2023/day/1

file_path = 'input/day01.txt'
file_prova = 'input/prova.txt'

def part1():
    with open(file_path) as f:
        s = 0
        for line in f:
            first = 0
            for letter in line:
                if letter.isdigit():
                    first = letter
                    break
            for letter in line[::-1]:
                if letter.isdigit():
                    last = letter
                    break
            s += int(first + last)
        return s   

def part2():
    return 0


print(f'Part 1: {part1()}')
print(f'Part 2: {part2()}')