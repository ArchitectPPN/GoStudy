package main

import "fmt"

func main()  {
	str := "wo|-|ow"

	// 生成一个rune类型的切片
	r := make([]rune, 0, len(str))
	for _, c := range str {
		r = append(r, c)
	}

	// 循环判断
	for index := 0; index < len(r) / 2; index++ {
		if r[index] != r[len(r)- 1 - index] {
			fmt.Printf("不是回文")
			return
		}
	}

	fmt.Printf("回文!")
}