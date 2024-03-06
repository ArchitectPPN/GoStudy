package main

import "fmt"

func main() {

	var level = ""
	var score = 11

	switch score {
	case 90:
		level = "A"
	case 80:
		level = "B"
	case 50:
		level = "C"
	default:
		level = "D"
	}

	fmt.Println("等级", level)

	var j = 2
	switch j {
	case 0:
		fallthrough
	case 1:
		fmt.Println("1")
	case 2, 4:
		fmt.Println(j)
	default:
		fmt.Println("def")
	}
}
