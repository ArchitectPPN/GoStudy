package main

import "fmt"

func main() {
	var sliceOne []int
	sliceOne[0] = 1

	defer deferTest("1")

	defer func(msg string) {
		fmt.Println("defer 我被执行了", msg)
	}("2")

	defer func() {
		if e := recover(); e != nil {
			println("--- defer ---")
		}
	}()
}

func deferTest(msg string) {
	fmt.Println("defer 我被执行了", msg)
}
