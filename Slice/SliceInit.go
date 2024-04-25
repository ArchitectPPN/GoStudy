package main

import "fmt"

func main() {
	// 总结: append 可以对nil切片追加元素

	// 切片初始化
	sliceOne := []int{1, 2, 3}
	fmt.Println("Slice One:", sliceOne)

	sliceTwo := []int{}
	sliceTwo[0] = 1
	fmt.Println("Slice Two:", sliceTwo)

	sliceTwo = append(sliceTwo, 10)
	fmt.Println("Slice Two Append Val:", sliceTwo)

	var sliceThree []int
	sliceThree = append(sliceThree, 10)
	fmt.Println("Slice Three Append Val:", sliceThree)
}
