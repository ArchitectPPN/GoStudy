package main

import "fmt"

func factorial(n uint64) uint64 {
	if n == 1 {
		return 1
	}

	return n * factorial(n - 1)
}

func main() {
	res := factorial(10)
	fmt.Printf("res: %v\n", res)
}
