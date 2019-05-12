package main

import "fmt"

func main() {
	res := Count()
	//fmt.Printf("%T: ", res, "\n")
	fmt.Println("res:", res())
	fmt.Println("res:", res())
	fmt.Println("res:", res())

	res2 := Count()
	fmt.Printf("%T", res2)
	fmt.Println("res:", res())
}

// 闭包函数实现计数器功能
func Count() func() int {
	i := 0
	res := func() int {
		i++
		return i
	}
	return res
}
