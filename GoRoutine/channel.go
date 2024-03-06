package main

import "fmt"

func main() {

	// 通道: 引用类型， 需要初始化之后才能使用

	/**
	1. 使用var声明
	*/
	var ch1 chan int
	ch1 = make(chan int, 1)
	ch1 <- 10
	x := <-ch1

	fmt.Println(x)
	close(ch1)

	// 像
	ch4 := make(chan int)
	go func() {
		ch4 <- 10
		close(ch4)
	}()
	p := <-ch4
	fmt.Println(p)

}
