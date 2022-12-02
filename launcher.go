/*package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	Alit "github.com/Alit54/AdventofCode/util"
)

func main() {
	// partA, partB := day1()
	// partA, partB := day2()
	partA, partB := day3()
	Alit.Max()
	fmt.Println("Part A:", partA)
	fmt.Println("Part B:", partB)
}

func day1() (int, int) {
	// Read the question here: https://adventofcode.com/2015/day/1
	input := readString(os.Args[1])
	return findFloor(input), findBasement(input)
}

func day2() (int, int) {
	// Read the question here: https://adventofcode.com/2015/day/2
	input := readRows(os.Args[1])
	return shopping(input)
}

func day3() (int, int) {
	// Read the question here: https://adventofcode.com/2015/day/3
	return 0, 0
}

// returns a string with the entire file read
func readString(name string) string {
	var input string
	file, err := os.Open(name)
	if err != nil {
		panic("File not found")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input += scanner.Text()
	}
	return input
}

// returns a slice of string with all the rows read from the file
func readRows(name string) []string {
	input := make([]string, 0)
	file, err := os.Open(name)
	if err != nil {
		panic("file not found")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		input = append(input, row)
	}
	return input
}

// Day1: returns the floor in which Santa arrives
func findFloor(input string) int {
	floor := 0
	for _, v := range input {
		switch string(v) {
		case "(":
			floor++
		case ")":
			floor--
		}
	}
	return floor
}

// Day1: returns the position of the character that sends Santa into the basement
func findBasement(input string) int {
	floor := 0
	for k, v := range input {
		switch string(v) {
		case "(":
			floor++
		case ")":
			floor--
		}
		if floor == -1 {
			return k + 1
		}
	}
	// in the case Santa dces not go into the basement
	return -1
}

// Day2: returns the shopping need
func shopping(input []string) (int, int) {
	paper := 0
	carbon := 0
	var value []string
	for _, v := range input {
		value = strings.Split(v, "x")
		l, _ := strconv.Atoi(value[0])
		w, _ := strconv.Atoi(value[1])
		h, _ := strconv.Atoi(value[2])
		paper += paperNeed(l, w, h)
		carbon += carboonNeed(l, w, h)
	}
	return paper, carbon
}

// Day2: returns the paper need
func paperNeed(l int, w int, h int) int {
	return 2*l*w + 2*w*h + 2*l*h + findMin(l*w, w*h, l*h)
}

// Day2: returns the carboon need
func carboonNeed(l int, w int, h int) int {
	return l*w*h + 2*findMin(l+h, h+w, l+w)
}

// returns min between 3 numbers
func findMin(a, b, c int) int {
	min := math.MaxInt
	if a < min {
		min = a
	}
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}
*/

package main

func main() {

}
