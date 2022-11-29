package main

import (
	Alit "github.com/Alit54/AdventofCode/util"
)

func main() {
	a := []int{1, 2, 3}
	t := Alit.ArrayToTree(a, 0)
	Alit.PrintTree(t)
}
