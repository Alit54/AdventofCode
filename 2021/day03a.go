package main

import (
	"fmt"
	"math"
)

func main() {
	var n, gamma, epsilon string
	input := make([]string, 0)
	map0 := make(map[int]int)
	map1 := make(map[int]int)
	var gam, eps float64
	for {
		fmt.Scan(&n)
		if n == "end" {
			break
		}
		input = append(input, n)
	}
	for i := 0; i < 12; i++ { //12 digits
		for j := 0; j < len(input); j++ {
			if input[j][i] == '0' {
				map0[i]++
			} else {
				map1[i]++
			}
		}
	}
	for i := 0; i < 12; i++ { // again 12 digits
		if map0[i] > map1[i] {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}
	for i := 0; i < 12; i++ {
		gam += float64(gamma[i]-48) * math.Pow(2, float64(11-i))
		eps += float64(epsilon[i]-48) * math.Pow(2, float64(11-i))
	}
	fmt.Println(gam * eps)
}
