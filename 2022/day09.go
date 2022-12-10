package main

import (
	"bufio"
	"strconv"
	"strings"

	Alit "github.com/Alit54/General/gofunc"
)

type Point struct {
	row    int
	column int
}

/*func main() {
	fmt.Println(Day9(os.Args[1]))
}*/

func Day9(input string) (int, int) {
	return Part9A(input), Part9B(input)
}

func Part9A(input string) int {
	var matrix [1000][1000]bool
	var step int
	file := Alit.File(input)
	scanner := bufio.NewScanner(file)
	// Inizialise head and tail in the center of the matrix
	var head, tail Point = Point{len(matrix) / 2, len(matrix) / 2}, Point{len(matrix) / 2, len(matrix) / 2}
	matrix[tail.row][tail.column] = true
	for scanner.Scan() {
		step++
		row := strings.Split(scanner.Text(), " ")
		steps, _ := strconv.Atoi(row[1])
		switch row[0] {
		case "U":
			{
				// Decrease the row of head
				for i := 0; i < steps; i++ {
					head.row = head.row - 1
					tail = followHead(head, tail)
					matrix[tail.row][tail.column] = true
				}
			}
		case "D":
			{
				// Increase the row of head
				for i := 0; i < steps; i++ {
					head.row = head.row + 1
					tail = followHead(head, tail)
					matrix[tail.row][tail.column] = true
				}
			}
		case "R":
			{
				// Increase the column of head
				for i := 0; i < steps; i++ {
					head.column = head.column + 1
					tail = followHead(head, tail)
					matrix[tail.row][tail.column] = true
				}
			}
		case "L":
			{
				// Decrease the row of head
				for i := 0; i < steps; i++ {
					head.column = head.column - 1
					tail = followHead(head, tail)
					matrix[tail.row][tail.column] = true
				}
			}
		}
	}
	return checkStep(matrix)
}

func Part9B(input string) int {
	// Here is the idea: array of 10 ints. index 0 is head for index 1, index 1 is head for index 2, etc...
	var matrix [1000][1000]bool
	var step int
	file := Alit.File(input)
	scanner := bufio.NewScanner(file)
	// Inizialise head and tail in the center of the matrix
	snake := [10]Point{{len(matrix) / 2, len(matrix) / 2}, {len(matrix) / 2, len(matrix) / 2}, {len(matrix) / 2, len(matrix) / 2}, {len(matrix) / 2, len(matrix) / 2}, {len(matrix) / 2, len(matrix) / 2}, {len(matrix) / 2, len(matrix) / 2}, {len(matrix) / 2, len(matrix) / 2}, {len(matrix) / 2, len(matrix) / 2}, {len(matrix) / 2, len(matrix) / 2}, {len(matrix) / 2, len(matrix) / 2}}
	matrix[snake[9].row][snake[9].column] = true
	for scanner.Scan() {
		step++
		row := strings.Split(scanner.Text(), " ")
		steps, _ := strconv.Atoi(row[1])
		switch row[0] {
		case "U":
			{
				// Decrease the row of head
				for i := 0; i < steps; i++ {
					snake[0].row = snake[0].row - 1
					for j := 0; j < 9; j++ {
						snake[j+1] = followHead(snake[j], snake[j+1])
					}
					matrix[snake[9].row][snake[9].column] = true
				}
			}
		case "D":
			{
				// Increase the row of head
				for i := 0; i < steps; i++ {
					snake[0].row = snake[0].row + 1
					for j := 0; j < 9; j++ {
						snake[j+1] = followHead(snake[j], snake[j+1])
					}
					matrix[snake[9].row][snake[9].column] = true
				}
			}
		case "R":
			{
				// Increase the column of head
				for i := 0; i < steps; i++ {
					snake[0].column = snake[0].column + 1
					for j := 0; j < 9; j++ {
						snake[j+1] = followHead(snake[j], snake[j+1])
					}
					matrix[snake[9].row][snake[9].column] = true
				}
			}
		case "L":
			{
				// Decrease the row of head
				for i := 0; i < steps; i++ {
					snake[0].column = snake[0].column - 1
					for j := 0; j < 9; j++ {
						snake[j+1] = followHead(snake[j], snake[j+1])
					}
					matrix[snake[9].row][snake[9].column] = true
				}
			}
		}
	}
	return checkStep(matrix)
}

func followHead(head, tail Point) Point {
	// tail needs to move if head is 2 steps further than tail
	// Vertical Movement UP
	if Alit.OR(head.row == tail.row-2 && head.column == tail.column) {
		tail.row = tail.row - 1
	}
	// Vertical Movement DOWN
	if Alit.OR(head.row == tail.row+2 && head.column == tail.column) {
		tail.row = tail.row + 1
	}
	// Horizontal Movement RIGHT
	if Alit.OR(head.column == tail.column+2 && head.row == tail.row) {
		tail.column = tail.column + 1
	}
	// Horizontal Movement LEFT
	if Alit.OR(head.column == tail.column-2 && head.row == tail.row) {
		tail.column = tail.column - 1
	}
	// Diagonal Movement NW
	if Alit.OR(head.row == tail.row-2 && head.column == tail.column-1, head.row == tail.row-1 && head.column == tail.column-2, head.row == tail.row-2 && head.column == tail.column-2) {
		tail.row = tail.row - 1
		tail.column = tail.column - 1
	}
	// Diagonal Movement NE
	if Alit.OR(head.row == tail.row-2 && head.column == tail.column+1, head.row == tail.row-1 && head.column == tail.column+2, head.row == tail.row-2 && head.column == tail.column+2) {
		tail.row = tail.row - 1
		tail.column = tail.column + 1
	}
	// Diagonal Movement SW
	if Alit.OR(head.row == tail.row+2 && head.column == tail.column-1, head.row == tail.row+1 && head.column == tail.column-2, head.row == tail.row+2 && head.column == tail.column-2) {
		tail.row = tail.row + 1
		tail.column = tail.column - 1
	}
	// Diagonal Movement SE
	if Alit.OR(head.row == tail.row+2 && head.column == tail.column+1, head.row == tail.row+1 && head.column == tail.column+2, head.row == tail.row+2 && head.column == tail.column+2) {
		tail.row = tail.row + 1
		tail.column = tail.column + 1
	}
	return tail
}

func checkStep(matrix [1000][1000]bool) int {
	count := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] {
				count++
			}
		}
	}
	return count
}
