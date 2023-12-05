# Read the question here: https://adventofcode.com/2023/day/4

file_path = 'input/day04.txt'
file_prova = 'input/prova.txt'

def part1():
    with open(file_path) as f:
        s=0
        for line in f:
            counter = 0
            line = line.rstrip()
            data = line.split(" | ")
            winners = data[0].split(": ")[1].split()
            for number in data[1].split():
                if number in winners:
                    counter+=1
            if counter:
                s+=2**(counter-1)
    return s

def part2():
    with open(file_path) as f:
        scratchcards = {i: 1 for i in range(1, 200)}
        for line in f:
            counter = 0
            line = line.rstrip()
            data = line.split(" | ")
            winners = data[0].split(": ")[1].split()
            for number in data[1].split():
                if number in winners:
                    counter+=1
            if counter:
                index = int(data[0].split(": ")[0][-3:])
                for i in range(index+1, index+counter+1):
                    scratchcards[i] += scratchcards[index]
    s=0
    for value in scratchcards:
        s += scratchcards[value]
    return s

print(f'Part 1: {part1()}')
print(f'Part 2: {part2()}')