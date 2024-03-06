package main

import "fmt"

func main()  {
	// 1. & 取地址符
	// 2. * 根据地址取值
	n := 19
	fmt.Println(&n)

	p := &n
	fmt.Println(p)
	fmt.Printf("p Type Of %T", p)

	// 根据地址取值
	o := *p
	fmt.Println(o)
}