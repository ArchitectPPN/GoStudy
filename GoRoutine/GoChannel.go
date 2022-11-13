package main

import (
	"fmt"
	"sync"
)

var channel chan int

var chanWg sync.WaitGroup

func writeToChannelGoRoutine(channel chan int, i int) {
	defer chanWg.Done()
	// 向channel写入数据
	fmt.Println("开始写入数据：", i)
	channel <- i

}

func writeToChannel(channel chan int) {
	// 向channel写入数据
	for i := 0; i < 100; i++ {
		fmt.Println("开始写入数据：", i)
		channel <- i
	}
}

func readFromChannel(channel chan int) {
	defer chanWg.Done()
	// 从channel读取数据
	fmt.Println("读取数据：", <-channel)
}

func main() {
	channel = make(chan int, 10000)
	writeToChannel(channel)

	for i := 0; i < 100; i++ {
		chanWg.Add(1)
		go readFromChannel(channel)
	}

	chanWg.Wait()
}
