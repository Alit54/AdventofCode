package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	Alit "github.com/Alit54/AdventofCode/util"
)

func main() {
	partA, partB := Day2(os.Args[1])
	fmt.Println("Part A:", partA)
	fmt.Println("Part B:", partB)
}

func Day2(input string) (int, int) {
	return Part2A(input), part2B(input)
}

// How many points will we do following the strategy guide?
func Part2A(input string) int {
	sum := 0
	file := Alit.File(input)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		sum += answer(row[1]) + score(row[0], row[1])
	}
	return sum
}

func part2B(input string) int {
	sum := 0
	file := Alit.File(input)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		sum += answerCorrect(row[0], row[1]) + scoreCorrect(row[1])
	}
	return sum
}

// Returns the points scored for choosing the answer
func answer(input string) int {
	switch input {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	default:
		return 0
	}
}

// Returns the score based on 2 answers
func score(elf string, ans string) int {
	/* Notes:
	A, X, 1: Rock
	B, Y, 2: Paper
	C, Z, 3: Scissors
	0: Lost
	3: Draw
	6: Win */
	switch elf {
	case "A":
		switch ans {
		case "X":
			return 3
		case "Y":
			return 6
		case "Z":
			return 0
		default:
			return 0
		}
	case "B":
		switch ans {
		case "X":
			return 0
		case "Y":
			return 3
		case "Z":
			return 6
		default:
			return 0
		}
	case "C":
		switch ans {
		case "X":
			return 6
		case "Y":
			return 0
		case "Z":
			return 3
		default:
			return 0
		}
	default:
		return 0
	}
}

// Now that elf show me the interpretation, I need 2 params
func answerCorrect(elf string, ans string) int {
	/* Notes:
	X: Lose
	Y: Draw
	Z: Win
	*/
	switch elf {
	case "A":
		switch ans {
		case "X":
			return 3
		case "Y":
			return 1
		case "Z":
			return 2
		default:
			return 0
		}
	case "B":
		switch ans {
		case "X":
			return 1
		case "Y":
			return 2
		case "Z":
			return 3
		default:
			return 0
		}
	case "C":
		switch ans {
		case "X":
			return 2
		case "Y":
			return 3
		case "Z":
			return 1
		default:
			return 0
		}
	default:
		return 0
	}
}

func scoreCorrect(input string) int {
	/* Notes:
	X: Lose
	Y: Draw
	Z: Win
	*/
	switch input {
	case "X":
		return 0
	case "Y":
		return 3
	case "Z":
		return 6
	default:
		return 0
	}
}
