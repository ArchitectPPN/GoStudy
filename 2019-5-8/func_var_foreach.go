package main

import "fmt"

/*
	1. 定义一个函数类型
	2. 实现定义的函数类型
	3. 作为参数调用
*/

type myFunc func(int) bool

func main(){
	arr := []int {23, 4, 5, 6, 77, 90, 98, 76, 762, 32, 13, 17}
	fmt.Println("slice = ", arr)

	// 获取切片中的奇数元素
	odd := Filter(arr, isOdd)
	fmt.Println("奇数元素: ", odd)

	// 获取切片中的偶数元素
	even := Filter(arr, isEven)
	fmt.Println("偶数元素: ", even)

	// 获取切片中的奇数元素
	oneEven := Filter(arr, addOneIsOdd)
	fmt.Println("+1是奇数:", oneEven)
	// 获取切片中的偶数元素
	oneOdd := Filter(arr, addOneIsEven)
	fmt.Println("+1是偶数:", oneOdd)
}

// 判断整型元素是偶数
func isEven(num int) bool {
	if num % 2 == 0{
		return true
	}

	return false
}

// 判断元素是奇数
func isOdd(num int) bool {
	if num % 2 == 0 {
		return false
	}

	return true
}

// 判断数值+1是否为奇数
func addOneIsOdd(num int) bool {
	if (num + 1) % 2 == 0 {
		return false
	}
	return true
}

func addOneIsEven(num int) bool {
	if (num + 1) % 2 == 0 {
		return true
	}

	return false
}

// 根据函数来处理切片, 实现奇书偶数分组, 返回新的切片
func Filter(arr []int, f myFunc) []int {
	var result []int

	for _, value := range arr {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}






