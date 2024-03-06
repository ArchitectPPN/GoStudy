package main

import "fmt"

// 计算数组里面的数值
var numArr = [5]int{1, 3, 5, 7, 8}

func main() {

	var sumTotal int = 0

	for _, temValue := range numArr {
		sumTotal += temValue
	}

	fmt.Println(sumTotal)
}