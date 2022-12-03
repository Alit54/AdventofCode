package main

import (
	"bufio"
	"fmt"
	"os"

	Alit "github.com/Alit54/AdventofCode/util"
)

func main() {
	fmt.Println(Day3(os.Args[1]))
}

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
		common := Alit.CommonRunes(row[0], row[1])
		for _, v := range common {
			if Alit.IsLower(string(v)) {
				sum += int(v-'a') + 1
			} else {
				sum += int(v-'A') + 27
			}
		}
		fmt.Println(sum, "Common Element:", row)
	}
	return sum
}

func Part3B(input string) int {
	return 0
}
