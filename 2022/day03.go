package main

import (
	"bufio"

	Alit "github.com/Alit54/General/gofunc"
)

/*func main() {
	fmt.Println(Day3(os.Args[1]))
}*/

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
		common := Alit.CommonRunes2(row[0], row[1])
		sum += priorities(common)
	}
	return sum
}

func Part3B(input string) int {
	// What is the sum of all the priorities?
	sum, index := 0, 0
	var group [3]string
	file := Alit.File(input)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		group[index] = scanner.Text()
		index++
		if index == 3 {
			index = 0
			common := Alit.CommonRunes3(group[0], group[1], group[2])
			sum += priorities(common)
		}
	}
	return sum
}

// Returns the sum of priorities of a single string.
func priorities(common []rune) int {
	sum := 0
	for _, v := range common {
		if Alit.IsLower(string(v)) {
			sum += int(v-'a') + 1
		} else {
			sum += int(v-'A') + 27
		}
	}
	return sum
}
