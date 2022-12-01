package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	Alit "github.com/Alit54/AdventofCode/util"
)

func main() {
	fmt.Println(Day1(os.Args[1]))
}

func Day1(input string) (int, int) {
	return partA(input), partB(input)
}

// Which elf needs more energy?
func partA(input string) int {
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
func partB(input string) int {
	count, support := 0, 0
	top := [3]int{0, 0, 0}
	f := Alit.File(input)
	defer f.Close()
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
			count=0
		}
		count += support
	}
	return top[0] + top[1] + top[2]
}
