package main

import "fmt"

type p struct {
	a int8
	b int8
	c int8
}

func main()  {
	// 结构体占用一块连续的内存。
	n := p{1,2,3}

	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
}
