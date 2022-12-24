package main

import "fmt"

func main() {
	var dep, hor, ans int
	var input string
	for {
		fmt.Scan(&input)
		_, err := fmt.Scan(&ans)
		if err != nil {
			break
		}
		fmt.Println(input, ans)
		switch input {
		case "forward":
			{
				hor += ans
			}
		case "down":
			{
				dep += ans
			}
		case "up":
			{
				dep -= ans
			}
		default:
			{
				panic("error")
			}
		}
	}
	fmt.Println(hor * dep)
}
