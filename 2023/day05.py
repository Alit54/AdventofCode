# Read the question here: https://adventofcode.com/2023/day/5

# NOTA: Per ora funziona con solo una categoria

file_path = 'input/day05.txt'
file_prova = 'input/prova.txt'

def part1():

    max = 10000000000

    translate = {i: i for i in range(max)}
    tmp = {i: i for i in range(max)}
    with open(file_path) as f:
        s = 0
        seeds = f.readline().split(": ")[1].split()
        for line in f:
            line = line.rstrip()
            if line == '':
                translate = {i: tmp[translate[i]] for i in translate}
                tmp = {i: i for i in range(max)}
                print('Finita Mappa')
                f.readline() # Skip Category Line
            else:
                dst, src, rng = [int(elem) for elem in line.split()]
                for i in range(rng):
                    tmp[src+i] = dst+i
    min = translate[int(seeds[0])]
    for elem in seeds:
        if translate[int(elem)] < min:
            min = translate[int(elem)]
    return min

def part2():
    with open(file_prova) as f:
        rows = f.readlines()
    s = 0
    return s

print(f'Part 1: {part1()}')
print(f'Part 2: {part2()}')