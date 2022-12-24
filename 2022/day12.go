// Special Credits to Luca Airoldi (https://github.com/zAiro12), who wrote the program with me and helped me to reach the solution.

package main

import (
	"os"

	Alit "github.com/Alit54/General/gofunc"
)

type Node struct {
	high byte
	step int
}

/*type Position struct {
	row    int
	column int
}

func main() {
	fmt.Println(Day12())
}*/

func Day12() (int, int) {
	return Part12A(), Part12B()
}

func Part12A() int {
	file := Alit.ReadMatrixByte(os.Args[1])
	// Create Matrix
	matrix := createMatrix(file)
	// Search
	start := find(matrix, 'S')
	finish := find(matrix, 'E')
	return UCSexl(start[0], matrix, finish)
}

func Part12B() int {
	file := Alit.ReadMatrixByte(os.Args[1])
	matrix := make(map[Position]Node)
	min := Part12A()
	var tmp int
	// Create matrix
	matrix = createMatrix(file)
	// Search
	start := find(matrix, 'S')
	start = append(start, find(matrix, 'a')...)
	finish := find(matrix, 'E')
	for _, v := range start {
		matrix = createMatrix(file)
		tmp = UCSexl(v, matrix, finish)
		if tmp != Alit.Min() {
		}
		if tmp < min {
			min = tmp
		}
	}
	return min
}

// Returns the positions of the Nodes called like char.
func find(matrice map[Position]Node, char byte) []Position {
	slice := make([]Position, 0)
	for k, v := range matrice {
		if v.high == char {
			slice = append(slice, k)
		}
	}
	return slice
}

// Returns the minimun number of steps made from start to finish, using a UCS and a list of expanded nodes. (UCS + EXL)
func UCSexl(start Position, matrix map[Position]Node, goal []Position) int {
	var border []Position
	EXL := make(map[Position]bool)
	var pos, next Position
	border = append(border, start)
	for len(border) != 0 {
		pos = border[0]
		border = Alit.RemoveFirstElement(border)
		// ASCII problem: 'S' is << than 'a'
		if matrix[pos].high == 'S' {
			matrix[pos] = Node{'a', 0}
		}
		if !EXL[pos] {
			// Found a goal
			if matrix[pos].high == 'E' {
				return matrix[pos].step
			}
			EXL[pos] = true
			// look UP
			next = Position{row: pos.row, column: pos.column - 1}
			if ((matrix[next].high == 'E' && matrix[pos].high == 'z') || (matrix[next].high <= matrix[pos].high+1 && matrix[next].high != 0 && matrix[next].high != 'E')) && !EXL[next] {
				border = append(border, next)
				matrix[next] = Node{matrix[next].high, matrix[pos].step + 1}
			}
			// look DOWN
			next = Position{row: pos.row, column: pos.column + 1}
			if ((matrix[next].high == 'E' && matrix[pos].high == 'z') || (matrix[next].high <= matrix[pos].high+1 && matrix[next].high != 0 && matrix[next].high != 'E')) && !EXL[next] {
				border = append(border, next)
				matrix[next] = Node{matrix[next].high, matrix[pos].step + 1}
			}
			// look LEFT
			next = Position{row: pos.row - 1, column: pos.column}
			if ((matrix[next].high == 'E' && matrix[pos].high == 'z') || (matrix[next].high <= matrix[pos].high+1 && matrix[next].high != 0 && matrix[next].high != 'E')) && !EXL[next] {
				border = append(border, next)
				matrix[next] = Node{matrix[next].high, matrix[pos].step + 1}
			}
			// look RIGHT
			next = Position{row: pos.row + 1, column: pos.column}
			if ((matrix[next].high == 'E' && matrix[pos].high == 'z') || (matrix[next].high <= matrix[pos].high+1 && matrix[next].high != 0 && matrix[next].high != 'E')) && !EXL[next] {
				border = append(border, next)
				matrix[next] = Node{matrix[next].high, matrix[pos].step + 1}
			}
		}
	}
	return Alit.Min()
}

// Creates the map required from a matrix of bytes.
func createMatrix(input [][]byte) map[Position]Node {
	matrix := make(map[Position]Node)
	var pos Position
	var node Node
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			pos.row = i
			pos.column = j
			node.high = input[i][j]
			node.step = 0
			matrix[pos] = node
		}
	}
	return matrix
}
