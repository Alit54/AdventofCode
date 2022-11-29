package main

import (
	"fmt"
	"os"
)

func main() {
	var partA, partB int
	switch os.Args[1] {
	case "day1.txt":
		{
			partA, partB = day1(os.Args[1])
		}
	case "day2.txt":
		{
			partA, partB = day2(os.Args[1])
		}
	case "day3.txt":
		{
			partA, partB = day3(os.Args[1])
		}
	case "day4.txt":
		{
			partA, partB = day4(os.Args[1])
		}
	case "day5.txt":
		{
			partA, partB = day5(os.Args[1])
		}
	case "day6.txt":
		{
			partA, partB = day6(os.Args[1])
		}
	case "day7.txt":
		{
			partA, partB = day7(os.Args[1])
		}
	case "day8.txt":
		{
			partA, partB = day8(os.Args[1])
		}
	case "day9.txt":
		{
			partA, partB = day9(os.Args[1])
		}
	case "day10.txt":
		{
			partA, partB = day10(os.Args[1])
		}
	case "day11.txt":
		{
			partA, partB = day11(os.Args[1])
		}
	case "day12.txt":
		{
			partA, partB = day12(os.Args[1])
		}
	case "day13.txt":
		{
			partA, partB = day13(os.Args[1])
		}
	case "day14.txt":
		{
			partA, partB = day14(os.Args[1])
		}
	case "day15.txt":
		{
			partA, partB = day15(os.Args[1])
		}
	case "day16.txt":
		{
			partA, partB = day16(os.Args[1])
		}
	case "day17.txt":
		{
			partA, partB = day17(os.Args[1])
		}
	case "day18.txt":
		{
			partA, partB = day18(os.Args[1])
		}
	case "day19.txt":
		{
			partA, partB = day19(os.Args[1])
		}
	case "day20.txt":
		{
			partA, partB = day20(os.Args[1])
		}
	case "day21.txt":
		{
			partA, partB = day21(os.Args[1])
		}
	case "day22.txt":
		{
			partA, partB = day22(os.Args[1])
		}
	case "day23.txt":
		{
			partA, partB = day23(os.Args[1])
		}
	case "day24.txt":
		{
			partA, partB = day24(os.Args[1])
		}
	case "day25.txt":
		{
			partA, partB = day25(os.Args[1])
		}
	}
	fmt.Println("Part A:", partA)
	fmt.Println("Part B:", partB)
}
