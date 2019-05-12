package main

import (
	"fmt"
)

func main(){
	res := Counter()
	fmt.Print("%T \n", res())
	fmt.Print("res: \t", res)
}

func Counter() func() int {
	i := 0
	res := func() int {
		i++
		return i
	}
	fmt.Println("Counter内部: ", res)
	return res
}
