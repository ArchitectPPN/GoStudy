package main

import "fmt"

func main() {
	// 创建一个nil切片
	var emptySlice []int
	// 覆盖下标为0的元素值
	emptySlice = append(emptySlice, 1)

	fmt.Println(emptySlice)
}
