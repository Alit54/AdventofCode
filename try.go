package main

import (
	"fmt"
	"os"

	Alit "github.com/Alit54/AdventofCode/util"
)

func main() {
	a := []any{"ciao", 5, 6.4, "Caora"}
	fmt.Println(Alit.Max())
	fmt.Println(Alit.Min())
	fmt.Println(Alit.RemoveLastElement(a))
	f := Alit.File(os.Args[1])
	defer f.Close()
}
