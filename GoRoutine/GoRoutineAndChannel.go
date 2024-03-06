package main

import "fmt"

// <-chan 限制通道只读 chan<- 限制通道只写
// 通道也可以不用手动关闭， 等通道内无元素时， 会自动东关闭
/*
两个goroutine
 1. 生成0-100的数字发送到ch1
 2. 从ch1中取出数据计算它的平方， 把结果发送到ch2中
*/
func generateRandomNum(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func getNumberFromChan(ch1 chan int, ch2 chan int) {
	// 从通道中取值的方式1
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go generateRandomNum(ch1)
	go getNumberFromChan(ch1, ch2)

	// 从通道中取值的方式2
	for ret := range ch2 {
		fmt.Println(ret)
	}
}
