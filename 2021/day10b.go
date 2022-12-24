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
	var i int
	var sum []int
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
		if flag {
			fmt.Println(exp)
			sum = append(sum, 0)
			for {
				if len(exp) == 0 {
					break
				}
				sum[(len(sum) - 1)] *= 5
				switch exp[len(exp)-1] {
				case ")":
					sum[(len(sum) - 1)] += 1
				case "]":
					sum[(len(sum) - 1)] += 2
				case "}":
					sum[(len(sum) - 1)] += 3
				case ">":
					sum[(len(sum) - 1)] += 4
				}
				exp = removeLastElement(exp)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic("Error while reading")
	}
	// Order sum
	for i := 0; i < len(sum); i++ {
		for j := 0; j < len(sum); j++ {
			if sum[i] < sum[j] {
				sum[i], sum[j] = sum[j], sum[i]
			}
		}
	}
	fmt.Println(sum)
	// print middle value
	fmt.Println(sum[len(sum)/2])
}

func removeLastElement(slice []string) []string {
	return append(slice[:len(slice)-1])
}
