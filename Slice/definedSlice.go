package main

import "fmt"

func main() {
	sliceOne := []int{1, 2, 3}

	fmt.Println("值:", sliceOne, "len:", len(sliceOne), "cap:", cap(sliceOne))

	fmt.Printf("未扩容前的地址:%p \n", sliceOne)

	// 引用传值
	appendVal(sliceOne)
	// 输出的值被修改了
	fmt.Println(sliceOne)
}

// 添加值
func appendVal(sliceOne []int) {
	fmt.Printf("函数内部，扩容之前地址: %p \n", sliceOne)
	sliceOne = append(sliceOne, 4)
	fmt.Println("函数内输出值:", sliceOne, "len:", len(sliceOne), "cap:", cap(sliceOne))
	fmt.Printf("扩容后的地址:%p \n", sliceOne)
}
