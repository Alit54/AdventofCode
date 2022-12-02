package main

import (
	"bufio"
	"strconv"

	Alit "github.com/Alit54/AdventofCode/util"
)

/*func main() {
	partA, partB := Day1(os.Args[1])
	fmt.Println("Part A:", partA)
	fmt.Println("Part B:", partB)
}*/

func Day1(input string) (int, int) {
	return part1A(input), part1B(input)
}

// Which elf needs more energy?
func part1A(input string) int {
	max, count, support := 0, 0, 0
	f := Alit.File(input)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		support, _ = strconv.Atoi(scanner.Text())
		if scanner.Text() == "" {
			if count > max {
				max = count
			}
			count = 0
		}
		count += support
	}
	return max
}

// Which elves need more energy (Top 3)? Returns the sum of their needs
func part1B(input string) int {
	count, support := 0, 0
	top := [3]int{0, 0, 0}
	f := Alit.File(input)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		support, _ = strconv.Atoi(scanner.Text())
		if scanner.Text() == "" {
			if count > top[0] {
				top[2] = top[1]
				top[1] = top[0]
				top[0] = count
				count = 0
				continue
			}
			if count > top[1] {
				top[2] = top[1]
				top[1] = count
				count = 0
				continue
			}
			if count > top[2] {
				top[2] = count
				count = 0
				continue
			}
			count = 0
		}
		count += support
	}
	return top[0] + top[1] + top[2]
}
