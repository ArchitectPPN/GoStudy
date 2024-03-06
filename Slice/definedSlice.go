package main

import "fmt"

func main()  {
	// 声明切片
	var sliceEx []string

	sliceEx = append(sliceEx, "1111")
	sliceEx = append(sliceEx, "11112222")
	fmt.Printf("%v", sliceEx)

	//
	var a []string
	var b = []string{"1", "2", "3"}
	var c = []bool{false,true}
	var d = []bool{false,true}

	fmt.Printf("%v, %v, %v, %v", a, b, c, d)
}
