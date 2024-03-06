package main

import (
	"fmt"
	"sync"
	"time"
)

func sendMessage(channel chan string, wg *sync.WaitGroup, message string, delay time.Duration) {
	defer wg.Done()

	// 延迟一段时间模拟异步操作
	time.Sleep(delay)

	channel <- message
}

func main() {
	var wg sync.WaitGroup

	// 创建两个通道
	channel1 := make(chan string)
	channel2 := make(chan string)

	// 启动两个goroutine向各自的通道发送数据
	wg.Add(2)
	go sendMessage(channel1, &wg, "Message from Channel 3", time.Second*1)
	go sendMessage(channel2, &wg, "Message from Channel 2", time.Second*2)

	// 使用select监听两个通道，并等待所有goroutine完成任务
	go func() {
		wg.Wait()
		close(channel1)
		close(channel2)
	}()

	// 主goroutine使用select监听两个通道
	for {
		select {
		case msg1 := <-channel1:
			fmt.Println("Received message from channel 1:", msg1)
		case msg2 := <-channel2:
			fmt.Println("Received message from channel 2:", msg2)
		case <-time.After(time.Second * 4): // 设置超时，防止死锁
			fmt.Println("No messages received within 3 seconds.")
			return
		}

		// 检查两个通道是否都已关闭
		if len(channel1) == 0 && len(channel2) == 0 {
			fmt.Println("All messages have been received.")
			return
		}
	}
}
