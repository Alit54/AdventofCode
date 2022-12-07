package main

import (
	"fmt"

	Alit "github.com/Alit54/General/gofunc"
)

/*func main() {
	fmt.Println(Day6(os.Args[1]))
}*/

func Day6(input string) (int, int) {
	return Part6A(input), Part6B(input)
}

func Part6A(input string) int {
	word := Alit.ReadRows(input)[0]
	for i := 3; i < len(word); i++ {
		if checkWord(word[i-3], word[i-2], word[i-1], word[i]) {
			return i + 1
		}
	}
	return -1
}

func Part6B(input string) int {
	word := Alit.ReadRows(input)[0]
	presents := make(map[byte]int)
	for i := 13; i < len(word); i++ {
		for j := i - 13; j <= i; j++ {
			presents[word[j]]++
		}
		if checkMap(presents) {
			fmt.Println(presents)
			return i + 1
		}
		presents = make(map[byte]int)
	}
	return -1
}

func checkWord(a, b, c, d byte) bool {
	return a != b && a != c && a != d && b != c && b != d && c != d
}

func checkMap(presents map[byte]int) bool {
	for _, v := range presents {
		if v > 1 {
			return false
		}
	}
	return true
}

/* NOTE: NEW func checkword
Mappa da byte a int, che rappresenta quante volte è presente un carattere negli ultimi 4/14 bit.
Se un byte ha più di 1 come valore, return false.
Else return true*/
