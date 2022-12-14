package main

import (
	"bufio"
	"strconv"
	"strings"

	Alit "github.com/Alit54/General/gofunc"
)

type monkey struct {
	items    []int
	sign     string
	num      int
	test     int
	tmonkey  int
	fmonkey  int
	activity int
}

/*func main() {
	fmt.Println(Day11(os.Args[1]))
}*/

func Day11(input string) (int, int) {
	return Part11A(input), Part11B(input)
}

func Part11A(input string) int {
	var n int
	monkeys := createMonkeys(input)
	for round := 0; round < 20; round++ {
		for i := 0; i < len(monkeys); i++ {
			for _, item := range monkeys[i].items {
				n = operation(monkeys[i].sign, item, monkeys[i].num) / 3
				if test(n, monkeys[i].test) {
					monkeys[monkeys[i].tmonkey].items = append(monkeys[monkeys[i].tmonkey].items, n)
				} else {
					monkeys[monkeys[i].fmonkey].items = append(monkeys[monkeys[i].fmonkey].items, n)
				}
				monkeys[i].activity++
			}
			monkeys[i].items = make([]int, 0)
		}
	}
	return findActivity(monkeys)
}

func Part11B(input string) int {
	/*var n int
	monkeys := createMonkeys(input)
	for round := 0; round < 10000; round++ {
		for i := 0; i < len(monkeys); i++ {
			for _, item := range monkeys[i].items {
				n = operation(monkeys[i].sign, item, monkeys[i].num)
				if test(n, monkeys[i].test) {
					monkeys[monkeys[i].tmonkey].items = append(monkeys[monkeys[i].tmonkey].items, n)
				} else {
					monkeys[monkeys[i].fmonkey].items = append(monkeys[monkeys[i].fmonkey].items, n)
				}
				monkeys[i].activity++
			}
			monkeys[i].items = make([]int, 0)
		}
	}
	return findActivity(monkeys)*/
	return 0
}

func createMonkeys(input string) []monkey {
	// Variables
	monkeys := make([]monkey, 8)
	var index, item int
	// Input
	file := Alit.File(input)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			// If so, a monkey has finished and we prepare the next monkey
			continue
		}
		row := strings.Split(scanner.Text(), " ")
		// Setting "Monkey &num" row
		if row[0] == "Monkey" {
			index, _ = strconv.Atoi(row[1][:len(row[1])-1])
			continue
		}
		// Setting "Starting items" row
		if row[2] == "Starting" {
			for i := 4; i < len(row); i++ {
				if i == len(row)-1 {
					item, _ = strconv.Atoi(row[i])
					monkeys[index].items = append(monkeys[index].items, item)
					break
				}
				item, _ = strconv.Atoi(row[i][:len(row[i])-1])
				monkeys[index].items = append(monkeys[index].items, item)
			}
			continue
		}
		// Setting "Operation" row
		if row[2] == "Operation:" {
			monkeys[index].num, _ = strconv.Atoi(row[7])
			monkeys[index].sign = row[6]
			continue
		}
		// Setting "Test" row
		if row[2] == "Test:" {
			monkeys[index].test, _ = strconv.Atoi(row[5])
			continue
		}
		// Setting TRUE monkey
		if row[5] == "true:" {
			monkeys[index].tmonkey, _ = strconv.Atoi(row[9])
			continue
		}
		// Setting FALSE monkey
		if row[5] == "false:" {
			monkeys[index].fmonkey, _ = strconv.Atoi(row[9])
			continue
		}
	}
	return monkeys
}

func test(old, num int) bool {
	return old%num == 0
}

func operation(sign string, old, num int) int {
	if sign == "+" {
		return old + num
	}
	if sign == "*" {
		if num == 0 {
			return old * old
		} else {

			return old * num
		}
	}
	return 0
}

// Returns the moltiplication between the 2 maximun values
func findActivity(slice []monkey) int {
	max, mmax := 0, 0
	for _, monkey := range slice {
		if monkey.activity > mmax {
			max = mmax
			mmax = monkey.activity
			continue
		}
		if monkey.activity > max {
			max = monkey.activity
		}
	}
	return max * mmax
}
