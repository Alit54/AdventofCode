package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	Alit "github.com/Alit54/General/gofunc"
)

func main() {
	fmt.Println(Day5(os.Args[1]))
}

func Day5(input string) (string, string) {
	return Part5A(input), Part5B(input)
}

func Part5A(input string) string {
	// Let's rearrange this
	stacks := createMap()
	file := Alit.File(input)
	defer file.Close()
	var quantity, start, finish int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &quantity, &start, &finish)
		// take "quantity" elements from start
		move := stacks[start][len(stacks[start])-quantity:]
		move = Alit.ReverseSlice(move)
		// append move to finish
		stacks[finish] = append(stacks[finish], move...)
		// remove quantity elements from start
		for i := 0; i < quantity; i++ {
			stacks[start] = Alit.RemoveLastElement(stacks[start])
		}
	}
	var answer string
	for i := 1; i < len(stacks); i++ {
		answer += stacks[i][len(stacks[i])-1]
	}
	return answer
}

func Part5B(input string) string {
	// Let's rearrange this
	stacks := createMap()
	file := Alit.File(input)
	defer file.Close()
	var quantity, start, finish int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &quantity, &start, &finish)
		// take "quantity" elements from start
		move := stacks[start][len(stacks[start])-quantity:]
		// append move to finish
		stacks[finish] = append(stacks[finish], move...)
		// remove quantity elements from start
		for i := 0; i < quantity; i++ {
			stacks[start] = Alit.RemoveLastElement(stacks[start])
		}
	}
	var answer string
	for i := 1; i < len(stacks); i++ {
		answer += stacks[i][len(stacks[i])-1]
	}
	return answer
}

func createMap() map[int][]string {
	// Func presa direttamente dall'AoC di Zairo. Genera la mappa con le stringhe al contrario, quindi si necessita un reverse
	file := Alit.File(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	stack := make(map[int][]string)
	for scanner.Scan() {
		nPile := (len(scanner.Text()) + 1) / 4
		if scanner.Text() == "" {
			break
		}
		index := 1
		for j := 1; j <= nPile; j++ {
			val := scanner.Text()[index]
			if unicode.IsLetter(rune(val)) {
				stack[j] = append(stack[j], string(val))
			}
			index += 4
		}
	}
	// Reorder
	for i := 1; i <= len(stack); i++ {
		stack[i] = Alit.ReverseSlice(stack[i])
	}
	return stack
}
