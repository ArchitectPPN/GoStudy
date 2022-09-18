package main

import "fmt"

func main() {

	/**
	1. 使用var声明
	*/
	var ch1 chan int // 引用类型， 需要初始化之后才能使用
	ch1 = make(chan int, 1)
	ch1 <- 10
	x := <-ch1

	fmt.Println(x)
	close(ch1)

	/**
	2. 类型推导
	*/
	//ch2 := make(chan int) // 无缓冲区通道, 直接使用会出错 fatal error: all goroutines are asleep - deadlock!
	//ch2 <- 10
	//p := <-ch2
	//fmt.Println(p)
	//close(ch2)

	ch3 := make(chan int, 1) // 有缓冲区通道
	ch3 <- 20
	i := <-ch3
	fmt.Println(i)
	// len(ch3) // 取通道中元素的数量
	// cap(ch3) // 拿到通道的容量
	close(ch3)
}
