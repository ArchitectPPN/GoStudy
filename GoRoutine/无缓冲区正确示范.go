package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var ch chan int
	ch = make(chan int)

	wg.Add(1)
	// 向里面丢入数据
	go func() {
		defer wg.Done()
		ch <- 1
	}()

	inputNum := <-ch
	fmt.Println("拿出的数据： ", inputNum)

	wg.Wait()

}
