package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	Alit "github.com/Alit54/General/gofunc"
)

/*func main() {
	fmt.Println(Day10(os.Args[1]))
	fmt.Println()
	Part10B(os.Args[1])
}*/

func Day10(input string) int {
	return Part10A(input)
}

func Part10A(input string) int {
	clock, register, signal := 0, 1, 0
	file := Alit.File(input)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		clock++
		if checkClock(clock) {
			signal += clock * register
		}
		if row[0] == "addx" {
			clock++
			if checkClock(clock) {
				signal += clock * register
			}
			amount, _ := strconv.Atoi(row[1])
			register += amount
		}
	}
	return signal
}

func Part10B(input string) {
	clock, register := 0, 1
	var matrix [6][40]bool
	file := Alit.File(input)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		clock++
		matrix[(clock-1)/40][(clock-1)%40] = draw((clock-1)%40, register)
		if row[0] == "addx" {
			clock++
			matrix[(clock-1)/40][(clock-1)%40] = draw((clock-1)%40, register)
			amount, _ := strconv.Atoi(row[1])
			register += amount
		}
	}
	drawMatrix(matrix)
}

func checkClock(clock int) bool {
	return Alit.OR(clock == 20, clock == 60, clock == 100, clock == 140, clock == 180, clock == 220)
}

func draw(clock, register int) bool {
	return Alit.OR(clock == register-1, clock == register, clock == register+1)
}

func drawMatrix(matrix [6][40]bool) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
