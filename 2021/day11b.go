package main

import (
	"fmt"
	"os"
	"strconv"
)

type matrix struct {
	number int
	flash  bool
}

func main() {
	var table [10][10]matrix
	var count int
	var check, answer bool
	var s rune
	// read the input
	file, _ := os.Open("day 11 input.txt")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Fscanf(file, "%c", &s)
			num, err := strconv.Atoi(string(s))
			if err == nil {
				table[i][j].number = num
			}
		}
	}
	// repeat until request's day: in this case 100
	for day := 1; ; day++ {
		check = true
		// check if some octopus blows and repeat it until there are no more blows (check = false)
		for check {
			check, count, table = checkBlows(table, count)
		}
		// +1 for every number of the matrix
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				table[i][j].number++
			}
		}
		// Put to 0 every number that is more than 8
		for i := 0; i < len(table); i++ {
			for j := 0; j < len(table[i]); j++ {
				if table[i][j].number > 9 {
					table[i][j].number = 0
				}
				table[i][j].flash = false
			}
		}
		answer = false
		answer = checkAnswer(table)
		if answer {
			fmt.Println(day)
			break
		}
	}
}

func checkAnswer(table [10][10]matrix) bool {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if table[i][j].number != 0 {
				return false
			}
		}
	}
	return true
}

func checkBlows(table [10][10]matrix, count int) (bool, int, [10][10]matrix) {
	var flash bool
	c := count
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if table[i][j].number > 8 && !table[i][j].flash {
				if i == 0 {
					if j == 0 {
						table[i][j].flash = true
						flash = true
						table[i][j+1].number++
						table[i+1][j].number++
						table[i+1][j+1].number++
						c++
					}
					if j == len(table[i])-1 {
						table[i][j].flash = true
						flash = true
						table[i][j-1].number++
						table[i+1][j].number++
						table[i+1][j-1].number++
						c++
					}
					if j != 0 && j != len(table[i])-1 {
						table[i][j].flash = true
						flash = true
						table[i][j-1].number++
						table[i][j+1].number++
						table[i+1][j-1].number++
						table[i+1][j].number++
						table[i+1][j+1].number++
						c++
					}
				}
				if i == len(table)-1 {
					if j == 0 {
						table[i][j].flash = true
						flash = true
						table[i][j+1].number++
						table[i-1][j].number++
						table[i-1][j+1].number++
						c++
					}
					if j == len(table[i])-1 {
						table[i][j].flash = true
						flash = true
						table[i][j-1].number++
						table[i-1][j].number++
						table[i-1][j-1].number++
						c++
					}
					if j != 0 && j != len(table[i])-1 {
						table[i][j].flash = true
						flash = true
						table[i][j-1].number++
						table[i][j+1].number++
						table[i-1][j-1].number++
						table[i-1][j].number++
						table[i-1][j+1].number++
						c++
					}
				}
				if i != 0 && i != len(table)-1 {
					if j == 0 {
						table[i][j].flash = true
						flash = true
						table[i][j+1].number++
						table[i-1][j].number++
						table[i-1][j+1].number++
						table[i+1][j].number++
						table[i+1][j+1].number++
						c++
					}
					if j == len(table[i])-1 {
						table[i][j].flash = true
						flash = true
						table[i][j-1].number++
						table[i-1][j].number++
						table[i-1][j-1].number++
						table[i+1][j].number++
						table[i+1][j-1].number++
						c++
					}
					if j != 0 && j != len(table[i])-1 {
						table[i][j].flash = true
						flash = true
						table[i][j-1].number++
						table[i][j+1].number++
						table[i-1][j-1].number++
						table[i-1][j].number++
						table[i-1][j+1].number++
						table[i+1][j-1].number++
						table[i+1][j].number++
						table[i+1][j+1].number++
						c++
					}
				}
			}

		}
	}
	return flash, c, table
}
