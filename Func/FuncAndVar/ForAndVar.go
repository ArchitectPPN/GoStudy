package main

import "fmt"

func main() {
	// for 循环不会影响 i的值
	i := 100
	for i := 0; i <= 100; i++ {
		//fmt.Println(i)
	}
	
	// i的值还是100
	fmt.Println(i)
}
