package main

import "fmt"

func main()  {
	var mapSlice = make([]map[string]string, 3)

	for index, value := range mapSlice {
		fmt.Printf("index: %d, value: %v\n", index, value)
	}

	fmt.Println("after init")

	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 2)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["age"] = "21"
	mapSlice[0]["address"] = "沙河"
	mapSlice[0]["height"] = "200 Cm"

	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}
