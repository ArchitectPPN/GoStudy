package main

import "fmt"

// 定义常量 常量名 数据类型 = 值
const WIDTH int = 100
const HEIGHT int = 200

const(
	FIRST = iota
	LENGTH = 1
	FLAG = 2
	SECOND = iota
	THREE = iota
)

const(
	FOUR = 1  << iota
	FIVE = 2 << iota
	SIX = 9 << iota
)

func main(){
	fmt.Printf("面积是: %d \n", WIDTH * HEIGHT)
	fmt.Println(LENGTH, FLAG)
	fmt.Printf("\n")

	fmt.Println(FIRST, SECOND, THREE, FOUR, FIVE, SIX)


}