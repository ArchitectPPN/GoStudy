package main

import "fmt"

func testLocalVar(x, y int) {
	fmt.Println(x, y)
	if x > 0 {
		// 未使用变量 z, 编译不通过: z declared but not used
		//z := 100
	}

	// 无法使用if 之内定义的变量, 可能不存在
	//fmt.Printf("%v", z)
}

func main() {
	testLocalVar(100, 2)
}
