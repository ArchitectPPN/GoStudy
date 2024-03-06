package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建两个通道
	channel1 := make(chan string)
	channel2 := make(chan string)

	// 启动一个goroutine向channel1发送数据
	go func() {
		// 延迟一秒模拟异步操作
		time.Sleep(time.Second * 1)
		channel1 <- "Message from Channel 1"
	}()

	// 启动另一个goroutine向channel2发送数据
	go func() {
		// 延迟一秒模拟异步操作
		time.Sleep(time.Second * 1)
		channel2 <- "Message from Channel 2"
	}()

	// 主goroutine使用select监听两个通道，因为使用了return，会在收到任何一条消息后退出
	for {
		select {
		case msg1 := <-channel1:
			fmt.Println("Received message from channel 1:", msg1)
			// 收到channel1的消息后，可以结束循环，
			return
		case msg2 := <-channel2:
			fmt.Println("Received message from channel 2:", msg2)
			// 收到channel2的消息后，也可以结束循环
			return
		case <-time.After(time.Second * 3): // 设置超时，防止死锁
			fmt.Println("No message received within 3 seconds.")
			return
		}
	}
}
