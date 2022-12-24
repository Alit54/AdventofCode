package main

import "fmt"

func main() {
	var n, prev, c int
	fmt.Scan(&prev)
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		if n > prev {
			c++
		}
		prev = n
	}
	fmt.Println(c)
}
