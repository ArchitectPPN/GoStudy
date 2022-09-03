package main

import "fmt"

type students struct {
	name string
	age, height, weight int
}

func main()  {
	// 使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。
	st := students{name: "ok", age:12, height: 23, weight: 110}

	// 也可以对结构体指针进行键值对初始化
	p6 := &students{name: "architect", age: 12}

	// 初始化结构体的时候可以简写，也就是初始化的时候不写键
	p8 := &students{"string", 12, 23, 220}

	fmt.Println(st)
	fmt.Println(p6)
	fmt.Println(p8)
}
