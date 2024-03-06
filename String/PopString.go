package main

import "fmt"

func main() {
	var str string
	str = "string"

	fmt.Println("源字符串: ", str)
	fmt.Println("拼接字符串: ", str+"p")
	fmt.Println("移除最后一个字符串: ", str[:len(str)-1])
}
