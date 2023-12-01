### First AOC Day in Python!
# Read the question here: https://adventofcode.com/2023/day/1

file_path = 'input/day01.txt'
file_prova = 'input/prova.txt'

def part1():
    with open(file_path) as f:
        s = 0
        for line in f:
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
    numbers = {
        'one': '1',
        'two': '2',
        'three': '3',
        'four': '4',
        'five': '5',
        'six': '6',
        'seven': '7',
        'eight': '8',
        'nine': '9'
    }
    with open(file_path) as f:
        s = 0
        for line in f:
            for i in range(len(line)):
                if line[i].isdigit():
                    first = line[i]
                    break
                else:     
                    for j in range(i+2, len(line)):
                        if line[i:j] in numbers:
                            first = numbers[line[i:j]]
                            break
                    else:
                        continue
                    break
            for i in range(len(line))[::-1]:
                if line[i].isdigit():
                    last = line[i]
                    break
                else:     
                    for j in range(0, i)[::-1]:
                        if line[j:i+1] in numbers:
                            last = numbers[line[j:i+1]]
                            break
                    else:
                        continue
                    break
            s += int(first + last)
        return s  


print(f'Part 1: {part1()}')
print(f'Part 2: {part2()}')