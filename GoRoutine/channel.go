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

	/**
	* 2. 类型推导
	* 无缓冲区通道, 直接使用会出错 fatal error: all goroutines are asleep - deadlock!
	* 所以需要像ch4一样使用, 在同一时刻同时写入/输出才可以
	 */
	//ch2 := make(chan int)
	//ch2 <- 10
	//p := <-ch2
	//fmt.Println(p)
	//close(ch2)

	// 像
	ch4 := make(chan int)
	go func() {
		ch4 <- 10
		close(ch4)
	}()
	p := <-ch4
	fmt.Println(p)

	// 有缓冲区通道
	//ch3 := make(chan int, 1)
	//ch3 <- 20
	//i := <-ch3
	//fmt.Println(i)
	//// len(ch3) // 取通道中元素的数量
	//// cap(ch3) // 拿到通道的容量
	//close(ch3)
}
