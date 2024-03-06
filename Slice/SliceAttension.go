package main

import "fmt"

func main() {
	sliceOne := []string{"2", "3"}

	fmt.Println("before add", sliceOne)
	fmt.Printf("before add addr: %p \n", sliceOne)

	addSlice(&sliceOne)

	fmt.Println("after add: ", sliceOne)
}

func addSlice(sliceOne *[]string) {
	sliceOne = append(sliceOne, "123")
	fmt.Printf("%p \n", sliceOne)
}
