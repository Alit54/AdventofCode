package main

import (
	"fmt"
	"math"
)

func main() {
	var input []int
	var fuel, min int
	// save the inputs
	for {
		_, err := fmt.Scan(&fuel)
		if err != nil {
			break
		}
		input = append(input, fuel)
	}
	min = 92000000000000000
	// calculate for every column the fuel necessary
	for i := 0; i < 2000; i++ {
		fuel = 0
		for j := 0; j < len(input); j++ {
			fuel += sum(int(math.Abs(float64(input[j] - i))))
			if fuel > min {
				break
			}
		}
		if fuel < min {
			min = fuel
		}
	}
	// take the minimum: that's the answer
	fmt.Println(min)
}

func sum(n int) int {
	var sum int
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}
