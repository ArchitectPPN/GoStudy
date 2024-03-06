package main

import "fmt"

func main() {
	/**
	* 2. 类型推导
	* 无缓冲区通道, 直接使用会出错 fatal error: all goroutines are asleep - deadlock!
	* 所以需要像ch4一样使用, 在同一时刻同时写入/输出才可以
	 */
	ch2 := make(chan int)
	ch2 <- 10
	p := <-ch2
	fmt.Println(p)
	close(ch2)
}
