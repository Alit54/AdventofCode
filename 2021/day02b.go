package main

import "fmt"

func main() {
	var dep, hor, ans, aim int
	var input string
	for {
		_, err :=fmt.Scan(&input)
		fmt.Scan(&ans)
		if err != nil {
			break
		}
		fmt.Println(input, ans)
		switch input {
		case "forward":
			{
				hor += ans
				dep += ans * aim
			}
		case "down":
			{
				aim += ans
			}
		case "up":
			{
				aim -= ans
			}
		default:
			{
				panic("error")
			}
		}
	}
	fmt.Println(hor * dep)
}
