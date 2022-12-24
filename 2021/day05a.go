package main

import (
	"fmt"
	"math"
)

type position struct {
	x, y int
}

func main() {
	var table [1000][1000]int
	var n, dots int
	var p position
	var p1, p2 []position
	var bin string
	// read the inputs and save them in 2 different slices
	for {
		// Save Position 1
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		p.x = n
		_, err = fmt.Scan(&n)
		if err != nil {
			break
		}
		p.y = n
		p1 = append(p1, p)
		// There is an arrow in the input, so read it and bin it
		fmt.Scan(&bin)
		// Save Position 2
		_, err = fmt.Scan(&n)
		if err != nil {
			break
		}
		p.x = n
		_, err = fmt.Scan(&n)
		if err != nil {
			break
		}
		p.y = n
		p2 = append(p2, p)
	}
	// check the horizontal and vertical lines
	p1, p2 = checkLines(p1, p2)
	// increase the table's dots in which there is a lines
	table = increase(p1, p2, table)
	// check in how many positions there is a >=2 number
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if table[i][j] >= 2 {
				dots++
			}
		}
	}
	fmt.Println(dots)
}

func increase(p1 []position, p2 []position, table [1000][1000]int) [1000][1000]int {
	var min, max int
	for i := 0; i < len(p1); i++ {
		// horizontal movement
		if p1[i].x == p2[i].x {
			min = int(math.Min(float64(p1[i].y), float64(p2[i].y)))
			max = int(math.Max(float64(p1[i].y), float64(p2[i].y)))
			for j := min; j <= max; j++ {
				table[p1[i].x][j]++
			}
		}
		// vertical movement
		if p1[i].y == p2[i].y {
			min = int(math.Min(float64(p1[i].x), float64(p2[i].x)))
			max = int(math.Max(float64(p1[i].x), float64(p2[i].x)))
			for j := min; j <= max; j++ {
				table[j][p1[i].y]++
			}
		}
	}
	return table
}

func checkLines(p1 []position, p2 []position) ([]position, []position) {
	// This function takes the slices and returns only the lines in which there is an horizontal or vertical movement
	var np1, np2 []position
	for i := 0; i < len(p1); i++ {
		if p1[i].x == p2[i].x || p1[i].y == p2[i].y {
			np1 = append(np1, p1[i])
			np2 = append(np2, p2[i])
		}
	}
	return np1, np2
}
