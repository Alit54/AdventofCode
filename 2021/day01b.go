package main

import "fmt"

func main() {
	var n, c int
	var val []int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		val = append(val, n)
	}
	for i := 0; i < len(val)-3; i++ {
		if val[i]+val[i+1]+val[i+2] < val[i+1]+val[i+2]+val[i+3] {
			c++
		}
	}
	fmt.Println(c)
}
