package main

import (
	"fmt"
	"sync"
)

// goroutine demo
var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("echo Hello~ ", i)
	wg.Done()
}

func main() { // 开启一个主goroutine去执行main函数
	wg.Add(10000) // 计数器+1

	for i := 0; i <= 10000; i++ {
		go hello(i) // 开启一个goroutine去执行main函数
	}

	fmt.Println("Hello main~")
	// 上面的执行结果可能只有一个 Hello main~
	wg.Wait()
}
