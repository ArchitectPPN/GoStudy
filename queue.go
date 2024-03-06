package main

import "fmt"

func main() {
	// 队列
	queue := make([]int, 0)
	// 入队
	queue = append(queue, 10)
	fmt.Println("队列: ", queue)

	// 出队
	v := queue[0]
	fmt.Println("v: ", v)
	queue = queue[1:]
	if len(queue) == 0 {
		fmt.Println("队列空了")
	}

	fmt.Println("queue: ", queue, "zero: ")
}
