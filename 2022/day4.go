package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	Alit "github.com/Alit54/AdventofCode/util"
)

func main() {
	fmt.Println(Day4(os.Args[1]))
}

type elf struct {
	start  int
	finish int
}

func Day4(input string) (int, int) {
	return Part4A(input), Part4B(input)
}

func Part4A(input string) int {
	file := Alit.File(input)
	defer file.Close()
	counter := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Row will be composed by "num-num,num-num"
		pairs := strings.Split(scanner.Text(), ",")
		elf1 := strings.Split(pairs[0], "-")
		elf2 := strings.Split(pairs[1], "-")
		if checkOverlap(makeElf(elf1), makeElf(elf2)) {
			counter++
		}
	}
	return counter
}

func Part4B(input string) int {
	file := Alit.File(input)
	defer file.Close()
	counter := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Row will be composed by "num-num,num-num"
		pairs := strings.Split(scanner.Text(), ",")
		elf1 := strings.Split(pairs[0], "-")
		elf2 := strings.Split(pairs[1], "-")
		if checkAggressiveOverlap(makeElf(elf1), makeElf(elf2)) {
			counter++
		}
	}
	return counter
}

func makeElf(data []string) elf {
	var e elf
	e.start, _ = strconv.Atoi(data[0])
	e.finish, _ = strconv.Atoi(data[1])
	return e
}

func checkOverlap(elf1 elf, elf2 elf) bool {
	if elf1.start <= elf2.start && elf1.finish >= elf2.finish {
		return true
	}
	if elf2.start <= elf1.start && elf2.finish >= elf1.finish {
		return true
	}
	return false
}

func checkAggressiveOverlap(elf1 elf, elf2 elf) bool {
	if elf1.start <= elf2.start && elf1.finish >= elf2.start {
		return true
	}
	if elf2.start <= elf1.start && elf2.finish >= elf1.start {
		return true
	}
	return false
}
