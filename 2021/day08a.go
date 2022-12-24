package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	var input [][]string
	var sum int
	// read and save inputs
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		s = scanner.Text()
		if s == "end" {
			break
		}
		input = append(input, strings.Split(s, " "))
	}
	// Consider only digits: 1, 4, 7, 8
	// lenght 1: 2
	// lenght 4: 4
	// lenght 7: 3
	// lenght 8: 7
	for i := 0; i < len(input); i++ {
		for j := 11; j < len(input[i]); j++ {
			if len(input[i][j]) == 2 || len(input[i][j]) == 3 || len(input[i][j]) == 4 || len(input[i][j]) == 7 {
				sum++
			}
		}
	}
	fmt.Println(sum)
}
