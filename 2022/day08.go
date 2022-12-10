package main

import (
	"bufio"
	"strconv"
	"strings"

	Alit "github.com/Alit54/General/gofunc"
)

/*func main() {
	fmt.Println(Day8(os.Args[1]))
}*/

func Day8(input string) (int, int) {
	return Part8A(input), Part8B(input)
}

func Part8A(input string) int {
	count := 0
	matrix := ReadMatrixInt(input)
	for i := 0; i < len(matrix); i++ { // Row
		for j := 0; j < len(matrix[i]); j++ { // Column
			if visible(matrix, i, j) {
				count++
			}
		}
	}
	return count
}

func Part8B(input string) int {
	max := 0
	matrix := ReadMatrixInt(input)
	for i := 0; i < len(matrix); i++ { // Row
		for j := 0; j < len(matrix[i]); j++ { // Column
			if visibility(matrix, i, j) > max {
				max = visibility(matrix, i, j)
			}
		}
	}
	return max
}

// This will need to be imported from Alit once Github updates.
func ReadMatrixInt(name string) [][]int {
	var input [][]int
	var sliceInt []int
	file := Alit.File(name)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		for _, v := range row {
			value, _ := strconv.Atoi(v)
			sliceInt = append(sliceInt, value)
		}
		input = append(input, sliceInt)
		sliceInt = make([]int, 0)
	}
	return input
}

// Part A purposes: returns if a tree is visible from outside.
func visible(matrix [][]int, row, column int) bool {
	if row == 0 || row == len(matrix)-1 || column == 0 || column == len(matrix[row])-1 {
		return true
	}
	flag := []bool{true, true, true, true}
	// Check Up
	for i := row - 1; i >= 0; i-- {
		if matrix[row][column] <= matrix[i][column] {
			flag[0] = false
		}
	}
	// Check Down
	for i := row + 1; i < len(matrix); i++ {
		if matrix[row][column] <= matrix[i][column] {
			flag[1] = false
		}
	}
	// Check Left
	for j := column - 1; j >= 0; j-- {
		if matrix[row][column] <= matrix[row][j] {
			flag[2] = false
		}
	}
	// Check Right
	for j := column + 1; j < len(matrix[row]); j++ {
		if matrix[row][column] <= matrix[row][j] {
			flag[3] = false
		}
	}
	return Alit.OR(flag...)
}

// Part B purposes: returns the visibility for a tree.
func visibility(matrix [][]int, row, column int) int {
	if row == 0 || row == len(matrix)-1 || column == 0 || column == len(matrix[row])-1 {
		return 0
	}
	view := make([]int, 4)
	// Check Up
	for i := row - 1; i >= 0; i-- {
		view[0]++
		if matrix[row][column] <= matrix[i][column] {
			break
		}
	}
	// Check Down
	for i := row + 1; i < len(matrix); i++ {
		view[1]++
		if matrix[row][column] <= matrix[i][column] {
			break
		}
	}
	// Check Left
	for j := column - 1; j >= 0; j-- {
		view[2]++
		if matrix[row][column] <= matrix[row][j] {
			break
		}
	}
	// Check Right
	for j := column + 1; j < len(matrix[row]); j++ {
		view[3]++
		if matrix[row][column] <= matrix[row][j] {
			break
		}
	}
	return Alit.Product(view...)
}
