package main

import (
	"fmt"
	"os"
)

func main() {
	partA, partB := Day2(os.Args[1])
	fmt.Println("Part A:", partA)
	fmt.Println("Part B:", partB)
}

func Day2(input string) (int, int) {
	return Part2A(input), part2B(input)
}

func Part2A(input string) int {
	return 0
}

func part2B(input string) int {
	return 0
}
