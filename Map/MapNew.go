package main

import "fmt"

func main()  {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 80
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["张三"])

	// 声明时初始化
	newMap := map[string]string{
		"name":"Architect",
		"provider_code":"TPTE",
	}

	fmt.Println(newMap)

	// 判断某个值是否存在
	existKey := map[string]string{
		"TPTE":"广州",
		"GTKY":"苏州",
	}

	value, exist := existKey["TPTE"]
	if exist {
		fmt.Printf("value: %v", value)
	} else {
		fmt.Println("该值不存在!")
	}

	// forRange
	for key, values := range existKey {
		fmt.Printf("key: %v value: %v", key, values)
	}

}
