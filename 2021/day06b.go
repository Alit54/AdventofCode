package main

import (
	"fmt"
)

func main() {
	var n int
	var fish [9]uint64
	// Memorise the lanternfishes
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		fish[n]++
	}
	for day := 1; day <= 256; day++ {
		fish[0], fish[1], fish[2], fish[3], fish[4], fish[5], fish[6], fish[7], fish[8] = fish[1], fish[2], fish[3], fish[4], fish[5], fish[6], fish[0]+fish[7], fish[8], fish[0]
	}
	fmt.Println(fish[0] + fish[1] + fish[2] + fish[3] + fish[4] + fish[5] + fish[6] + fish[7] + fish[8])
}

// Thanks to L. Laboni for this piece of shit
