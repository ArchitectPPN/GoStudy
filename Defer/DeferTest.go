package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获panic", err)
		}
	}()

	var sliceOne []int
	sliceOne[0] = 1

	return
}
