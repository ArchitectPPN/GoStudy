package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main()  {
	rand.Seed(time.Now().UnixNano()) // 初始化随机种子
	var scoreMap = make(map[string]int, 200)
	for i := 0; i<= 100; i++ {
		key := fmt.Sprintf("stu%03d", i) // 生成stu开头的字符串
		value := rand.Intn(100)	// 生成0-99 的数字

		scoreMap[key] = value
	}
	fmt.Println(scoreMap)

	// 取出map中所有的key, 存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	fmt.Println(keys)

	// 对切片进行排序
	sort.Strings(keys)
	// 按照排序后的顺序遍历
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}