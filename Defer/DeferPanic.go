package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover Panic")
		} else {
			fmt.Println("fatal")
		}
	}()

	testOne()
}

func testOne() {
	var sliceOne []int
	sliceOne[0] = 1
}
