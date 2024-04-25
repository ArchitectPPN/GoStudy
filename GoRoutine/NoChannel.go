package main

import (
	"fmt"
	"sync"
)

func main() {
	// 新建一个channel
	ch := make(chan int, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		value := <-ch
		fmt.Println("输入的值", value)
	}()

	fmt.Println("发送数据到channel...")
	ch <- 42
	fmt.Println("数据已发送到channel中...")

	wg.Wait()
	fmt.Println("主协程已执行完毕")
}
