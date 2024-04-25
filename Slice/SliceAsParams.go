package main

import "fmt"

// 切片传参, 在容量足够的情况下, 函数内部修改其值, 会影响到函数体外的值
func main() {
	var ageSlice []int
	ageSlice = make([]int, 2)
	ageSlice[0] = 12
	ageSlice[1] = 21

	fmt.Println("origin Age: ", ageSlice)

	updateAge(ageSlice)

	fmt.Println("update Age: ", ageSlice)
}

func updateAge(age []int) {
	age[0] = 99
}
