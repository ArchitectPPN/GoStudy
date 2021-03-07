package main

import "fmt"

func main() {

	big, small := 1, 2
	var now int
	now = max(big, small)

	fmt.Printf("%d", now)

	var str string = "我是字符串哦"

	swap(big, str)
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return  num1
	}
	return num2
}

func swap(string1 int, string2 string) string {
	fmt.Printf("%d %s", string1, string2)
	return string2
}
