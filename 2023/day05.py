# Read the question here: https://adventofcode.com/2023/day/5

# NOTA: Per ora funziona con solo una categoria

file_path = 'input/day05.txt'
file_prova = 'input/prova.txt'

def part1():
    seed_to_soil = {i: i for i in range(100)}
    soil_to_fertilizer = {}
    fertilizer_to_water = {}
    water_to_light = {}
    light_to_temperature = {}
    temperature_to_humidity = {}
    humidity_to_location = {}
    with open(file_prova) as f:
        s = 0
        seeds = f.readline().split(": ")[1].split()
        print(seeds)     
        for line in f:
            line = line.rstrip()
            print(line)
            if line == '':
                # Change of Category [!!! Arrivato qui !!!]
                f.readline() # Skip Category Line
            else:
                dst, src, rng = [int(elem) for elem in line.split()]
                for i in range(rng):
                    print(src+i, dst+i, 'Print i')
                    seed_to_soil[src+i] = dst+i
        for k in seed_to_soil:
            print(k, seed_to_soil[k])            
    return s

def part2():
    with open(file_prova) as f:
        rows = f.readlines()
    s = 0
    return s

print(f'Part 1: {part1()}')
print(f'Part 2: {part2()}')