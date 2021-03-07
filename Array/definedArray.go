package main

import "fmt"

func main()  {
	//声明方法1
	var person [3]string
	person[0] = "小李子"
	person[1] = "Wonderful"
	person[2] = "KangKang"

	fmt.Printf("person的值： %v\n", person)

	// 声明方法2
	country := [3]string{"中国", "美国", "应该"}
	fmt.Printf("国家列表： %v", country)

	// 声明方法3
	language := [...]string{1:"PHP", 5:"GoLang"}
	fmt.Printf("热门语言： %v", language)

	// 多维数组
	milt := [...][2]string{
		{"北极","伤害"},
		{"OK","PP"},
		{"北极","伤害"},
	}
	fmt.Printf("热门语言： %v", milt)
}
