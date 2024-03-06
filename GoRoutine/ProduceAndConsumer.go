package main

import (
	"fmt"
	"strconv"
	"strings"
)

type orderInfo struct {
	orderId   string
	orderCode string
}

func producer() <-chan orderInfo {
	ch := make(chan orderInfo, 2)

	// 创建一个新的goroutine执行发送数据
	go func() {
		order := orderInfo{}

		for i := 0; i < 10; i++ {
			var builder strings.Builder
			builder.WriteString("OrderId:")
			builder.WriteString(strconv.Itoa(i))
			order.orderId = builder.String()

			var builder1 strings.Builder
			builder1.WriteString("OrderCode:")
			builder1.WriteString(strconv.Itoa(i))
			order.orderCode = builder1.String()
			ch <- order
		}

		close(ch)
	}()
	return ch
}

func consumer(ch <-chan orderInfo) {
	for value := range ch {
		fmt.Println(value)
	}
}

func main() {
	ch := producer()

	go consumer(ch)
}
