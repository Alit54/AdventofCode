package main

import (
	"bufio"

	Alit "github.com/Alit54/AdventofCode/util"
	"golang.org/x/text/search"
)

func Day3(input string) (int, int) {
	return Part3A(input), Part3B(input)
}

func Part3A(input string) int {
	// What is the sum of all the priorities?
	sum := 0
	file := Alit.File(input)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := Alit.SplitStrings(scanner.Text(), len(scanner.Text())/2)
		common := 
	}
	return sum
}
