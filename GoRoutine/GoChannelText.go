package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			ch <- i
			fmt.Println("写入 ", i)
		}
		fmt.Println("关闭通道 ")
		close(ch)

	}()

	for i := 0; i < 6; i++ {
		fmt.Println("读出 ", <-ch)
	}
}
