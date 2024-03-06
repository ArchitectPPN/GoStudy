package main

import "fmt"

func main() {
	// 有缓冲区通道
	ch := make(chan int, 1)
	ch <- 20

	// 放入元素之后， 通道里面的元素数量
	beforeChLen := len(ch)
	fmt.Println("ch 放入通道之后通道中元素的数量：", beforeChLen)

	i := <-ch
	fmt.Println("通道中元素的值", i)

	// 取通道中元素的数量
	afterChLen := len(ch)
	fmt.Println("ch 取出通道之后通道中元素的数量：", afterChLen)

	// 拿到通道的容量
	chCap := cap(ch)
	fmt.Println("ch 可以容纳元素的数量：", chCap)

	close(ch)
}
