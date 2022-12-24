package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input [][]string
	var sum, num int
	var s string
	_ = sum
	// save the inputs
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		s = scanner.Text()
		if s == "end" {
			break
		}
		input = append(input, strings.Split(s, ""))
	}
	// check the lowest points, paying attention to the borders
	for i := 0; i < len(input); i++ { // rows
		for j := 0; j < len(input[i]); j++ { // columns
			if i == 0 { // first row
				if j == 0 { // front left corner
					if input[i][j] < input[i+1][j] && input[i][j] < input[i][j+1] {
						num, _ = strconv.Atoi(input[i][j])
						sum += num + 1
					}
					continue
				}
				if j == len(input[i])-1 { // front right corner
					if input[i][j] < input[i+1][j] && input[i][j] < input[i][j-1] {
						num, _ = strconv.Atoi(input[i][j])
						sum += num + 1
					}
					continue
				}
				if input[i][j] < input[i+1][j] && input[i][j] < input[i][j+1] && input[i][j] < input[i][j-1] {
					num, _ = strconv.Atoi(input[i][j])
					sum += num + 1
					continue
				}
			}
			if i == len(input)-1 { // last row
				if j == 0 { // rear left corner
					if input[i][j] < input[i-1][j] && input[i][j] < input[i][j+1] {
						num, _ = strconv.Atoi(input[i][j])
						sum += num + 1
					}
					continue
				}
				if j == len(input[i])-1 { // rear right corner
					if input[i][j] < input[i-1][j] && input[i][j] < input[i][j-1] {
						num, _ = strconv.Atoi(input[i][j])
						sum += num + 1
					}
					continue
				}
				if input[i][j] < input[i-1][j] && input[i][j] < input[i][j+1] && input[i][j] < input[i][j-1] {
					num, _ = strconv.Atoi(input[i][j])
					sum += num + 1
					continue
				}
			}
			if i != 0 && i != len(input)-1 {
				if j == 0 { // first column
					if input[i][j] < input[i-1][j] && input[i][j] < input[i+1][j] && input[i][j] < input[i][j+1] {
						num, _ = strconv.Atoi(input[i][j])
						sum += num + 1
					}
					continue
				}
				if j == len(input[i])-1 { // last column
					if input[i][j] < input[i-1][j] && input[i][j] < input[i][j-1] && input[i][j] < input[i+1][j] {
						num, _ = strconv.Atoi(input[i][j])
						sum += num + 1
					}
					continue
				}
				if j != 0 && j != len(input[i])-1 {
					if input[i][j] < input[i-1][j] && input[i][j] < input[i+1][j] && input[i][j] < input[i][j-1] && input[i][j] < input[i][j+1] {
						num, _ = strconv.Atoi(input[i][j])
						sum += num + 1
					}
				}

			}
		}
	}
	// Print the sum
	fmt.Println(sum)
}
