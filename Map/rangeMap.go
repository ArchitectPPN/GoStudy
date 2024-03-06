package main

import "fmt"

func main() {
	// 使用make定义
	mapDefinedFour := make(map[int]string, 5)

	mapDefinedFour[0] = "zero"
	mapDefinedFour[1] = "one"
	mapDefinedFour[2] = "two"
	mapDefinedFour[3] = "three"
	mapDefinedFour[4] = "four"
	mapDefinedFour[5] = "five"

	fmt.Println("mapDefinedOne的长度：", len(mapDefinedFour))

	// map的遍历
	for key, value := range mapDefinedFour {
		fmt.Println(" key:", key, "value:", value)
	}
}
