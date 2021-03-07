package main

import (
	"fmt"
	"strings"
)

// 定义一个新类型
type CaseFunc func(string) string

func main(){

	str := "asdSDASDrwqrSDFAdWEasdSDWEWIA"
	// fmt.Println(processLetter(str))
	// 函数变量普通用法
	fmt.Println(StringToCase(str, processLetter))

	// 函数变量Type用法
	fmt.Println(StringToCase2(str, processLetter))
}

// 处理字符串, 实现奇偶交替

func processLetter(str string) string {
	result := ""
	for i, value := range str {
		if i % 2 == 0 {
			result += strings.ToUpper(string(value))
		} else {
			result += strings.ToLower(string(value))
		}
	}

	return result
}

// 使用结构
func StringToCase(str string, myfunc func(string) string) string {
	return myfunc(str)
}

// 不使用结构
func StringToCase2(str string, myfunc CaseFunc) string {
	return myfunc(str)
}