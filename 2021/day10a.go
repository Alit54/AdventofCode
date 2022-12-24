package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	var row, exp []string
	var sum, i int
	var flag bool
	// Read file
	file, err := os.Open("day 10 input.txt")
	if err != nil {
		panic("Error while opening the file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		flag = true
		exp = nil
		s = scanner.Text()
		row = strings.Split(s, "")
		for i = 0; i < len(row); i++ {
			switch row[i] {
			// If you find an opening bracket, then you can replace the expected character with the one you found
			case "(":
				exp = append(exp, ")")
			case "[":
				exp = append(exp, "]")
			case "{":
				exp = append(exp, "}")
			case "<":
				exp = append(exp, ">")
			// If you find a closing bracket, you need to check if it is a valid bracket. If not, that one is the first problem
			case ")", "]", "}", ">":
				if exp[len(exp)-1] == row[i] {
					exp = removeLastElement(exp)
					continue
				}
				if exp[len(exp)-1] != row[i] && i <= len(row) {
					flag = false
				}
			}
			if !flag {
				break
			}
		}
		if !flag {
			/* ): 3 points.
			]: 57 points.
			}: 1197 points.
			>: 25137 points. */
			switch row[i] {
			case ")":
				sum += 3
			case "]":
				sum += 57
			case "}":
				sum += 1197
			case ">":
				sum += 25137
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic("Error while reading")
	}
	// Print the sum: that's your answer
	fmt.Println(sum)
}

func removeLastElement(slice []string) []string {
	return append(slice[:len(slice)-1])
}
