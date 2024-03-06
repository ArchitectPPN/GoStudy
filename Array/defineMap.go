package main

import "fmt"

func main() {
	// 使用make定义一个map， 指定cap容量
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	// 定义时直接初始化
	appInfoMap := map[string]string{
		"name":"爱奇艺",
		"userInfo":"爱奇艺科技",
		"develop":"爱奇艺",
	}

	fmt.Println(appInfoMap)
	fmt.Println(appInfoMap["name"])

	// 检查key是不是存在的
	mapValue, mapKeyExist := appInfoMap["name"]
	if mapKeyExist {
		fmt.Println("查询的key：", mapValue)
	} else {
		fmt.Println("查询的key不存在！")
	}

	// 删除map
	delete(appInfoMap, "name")
	mapValue, mapKeyExist = appInfoMap["name"]
	if mapKeyExist {
		fmt.Println("删除掉查询的key：", mapValue)
	} else {
		fmt.Println("删除掉查询的key不存在！")
	}

	mapValue1, mapKeyExist1 := appInfoMap["查询的key"]
	if mapKeyExist1 {
		fmt.Println("查询的key：", mapValue1)
	} else {
		fmt.Println("查询的key不存在！")
	}

}


