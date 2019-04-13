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
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	return result
}

func swap(string1 int, string2 string) string {
	fmt.Printf("%d %s", string1, string2)
	return string2
}
