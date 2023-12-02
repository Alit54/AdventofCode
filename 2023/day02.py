# Read the question here: https://adventofcode.com/2023/day/2

file_path = 'input/day02.txt'
file_prova = 'input/prova.txt'

def part1():
    with open(file_path) as f:
        s = 0
        for line in f:
            line = line.rstrip()
            numbers = {
                'blue': 0,
                'green': 0,
                'red': 0
            }
            line = line.split(": ") # Index: line[0]; Game: line[1]
            sets = line[1].replace(";", ",")
            cubes = sets.split(", ")
            for cube in cubes:
                values = cube.split(" ")
                numbers[values[1]] = int(values[0])
                if (numbers['blue'] > 14) | (numbers['green'] > 13) | (numbers['red'] > 12):
                    break
            else:
                s += int(line[0].split(" ")[1])
    return s

def part2():
    with open(file_path) as f:
        s = 0
        for line in f:
            line = line.rstrip()
            numbers = {
                'blue': 0,
                'green': 0,
                'red': 0
            }
            line = line.split(": ") # Index: line[0]; Game: line[1]
            sets = line[1].replace(";", ",")
            cubes = sets.split(", ")
            for cube in cubes:
                values = cube.split(" ")
                if int(values[0]) > numbers[values[1]]:
                    numbers[values[1]] = int(values[0])
            s += numbers['blue'] * numbers['green'] * numbers['red']
    return s

print(f'Part 1: {part1()}')
print(f'Part 2: {part2()}')