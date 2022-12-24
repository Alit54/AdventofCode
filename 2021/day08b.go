package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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
	// try to decode every line
	for i := 0; i < len(input); i++ {
		num := decode(input[i])
		sum += num
	}
	// sum the values
	fmt.Println(sum)
}

func decode(line []string) int {
	var number int
	var code string
	display := make(map[int]string, 10)
	// Reorder all the strings
	for i := 0; i < len(line); i++ {
		line[i] = SortString(line[i])
	}
	// put the number you're sure of
	for i := 0; i < 10; i++ {
		switch len(line[i]) {
		case 2:
			display[1] = line[i]
		case 3:
			display[7] = line[i]
		case 4:
			display[4] = line[i]
		case 7:
			display[8] = line[i]
		}
	}
	// the string of len 6 which doesn't have both 1's chars is 6
	for i := 0; i < 10; i++ {
		if len(line[i]) != 6 {
			continue
		}
		if !strings.Contains(line[i], display[1][:1]) || !strings.Contains(line[i], display[1][1:]) {
			display[6] = line[i]
			break
		}
	}
	// the string of len 5 which is contained by the 6 is 5
	for i := 0; i < 10; i++ {
		if len(line[i]) != 5 {
			continue
		}
		c := 0
		for j := 0; j < 5; j++ {
			if strings.Contains(display[6], string(line[i][j])) {
				c++
			}
		}
		if c == 5 {
			display[5] = line[i]
			break
		}
	}
	// the string of len 6 (=/= from 6) which contains 5 is 9
	for i := 0; i < 10; i++ {
		if len(line[i]) != 6 || line[i] == display[6] {
			continue
		}
		c := 0
		for j := 0; j < 5; j++ {
			if strings.Contains(line[i], string(display[5][j])) {
				c++
			}
		}
		if c == 5 {
			display[9] = line[i]
			break
		}
	}
	// the string of len 6 which is not 9 or 6, is 0
	for i := 0; i < 10; i++ {
		if len(line[i]) == 6 && line[i] != display[9] && line[i] != display[6] {
			display[0] = line[i]
			break
		}
	}
	// the string of len 5 that contains 7, is 3. The other one is 2
	for i := 0; i < 10; i++ {
		if len(line[i]) != 5 {
			continue
		}
		c := 0
		for j := 0; j < 3; j++ {
			if strings.Contains(line[i], string(display[7][j])) {
				c++
			}
		}
		if c == 3 {
			display[3] = line[i]
			break
		}
	}
	for i := 0; i < 10; i++ {
		if len(line[i]) == 5 && line[i] != display[3] && line[i] != display[5] {
			display[2] = line[i]
			break
		}
	}
	// read and convert
	for i := 11; i < len(line); i++ {
		switch line[i] {
		case display[0]:
			code += "0"
		case display[1]:
			code += "1"
		case display[2]:
			code += "2"
		case display[3]:
			code += "3"
		case display[4]:
			code += "4"
		case display[5]:
			code += "5"
		case display[6]:
			code += "6"
		case display[7]:
			code += "7"
		case display[8]:
			code += "8"
		case display[9]:
			code += "9"
		}
	}
	number, _ = strconv.Atoi(code)
	return number
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
