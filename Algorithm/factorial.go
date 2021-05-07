package main

import "fmt"

func factorial(n uint64) uint64 {
	if n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

// 上台阶问题
func upStep(n uint64) uint64 {
	if n == 1 {
		// 只有一个台阶
		return 1
	} else if n == 2 {
		// 只有两个台阶
		return 2
	}

	return upStep(n-1) + upStep(n-2)
}

func main() {
	res := factorial(10)
	fmt.Printf("res: %v\n", res)
}
