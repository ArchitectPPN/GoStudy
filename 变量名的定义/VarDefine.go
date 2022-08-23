package main

import "fmt"

var globalVar = "global var"

// := 不能在函数外使用 var glo := 12

func main()  {
	// 变量名的声明定义
	var xFloat64 float64
	var xInt int

	xFloat64 = 3.165
	xInt = 1

	fmt.Println(xFloat64)
	fmt.Println(xInt)

	// 批量定义
	var (
		stringOne string
		numberOne int
		boolOne bool
		float32One float32
	)

	stringOne = "我是字符串"
	numberOne = 1
	boolOne = true
	float32One = 3.1415962

	fmt.Println(stringOne)
	fmt.Println(numberOne)
	fmt.Println(boolOne)
	fmt.Println(float32One)

	// 类型推导
	var nameString = "推导字符串"
	var ageInt = 26
	var boolFlag = false
	fmt.Println(nameString)
	fmt.Println(ageInt)
	fmt.Println(boolFlag)

	// 一次性定义变量
	var name, age = "我的名字~ string", 26
	fmt.Println("一次性定义变量：string" , name, "int ", age)

	// 短变量声明
	stringTwo := "string"
	intNumber := 46
	fmt.Println("短变量声明：string", stringTwo)
	fmt.Println("短变量声明：int", intNumber)

	// 全局变量
	fmt.Println("全局变量：string", globalVar)

	// 匿名变量
	var anonymousVar, _ = "匿名变量", 26
	fmt.Println("匿名变量： string", anonymousVar)
}
