package main

import "fmt"

func main(){
	for i := 0; i<10; i++{
		fmt.Println(adder(i))
	}
}

func adder(num int) int {
	sum := 0
	return func(num int) int {
		sum += num
		return sum
	}(num)
}