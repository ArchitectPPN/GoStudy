package main

import "fmt"

func main() {
	switchTest()
}

func switchTest() {
	num1, num2, res := 12, 13, 0

	var Operation = "+"
	switch Operation {
	case "+":
		res = num1 + num2
		//fallthrough
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		res = num1 / num2
	}

	fmt.Printf("res : %d", res)
}
