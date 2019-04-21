package main

import (
	"fmt"
)

const (
	FLAG = 1
	NAME = "我的年龄"
	// 如果没有给常量赋值, 将会使用上面所定义的
	AGE = iota
	B   = iota
	ID  = iota
	K   = iota
)

type Man struct {
	age, height int
}

func main() {
	man := Man{21, 173}
	//fmt.Printf("常量的数值为: %d \n", FLAG)
	//
	fmt.Println(man)
	//fmt.Print(AGE)
	fmt.Println(AGE, B, ID, K)

	const (
		a1 = '-'
		a2
		a3 = iota
		a4
	)

	fmt.Println(a1, a2, a3, a4)

	const name = iota
	const hh = iota
	fmt.Println(name, hh)

	var integer = 2

	switchTest(integer)

}

func switchTest(test int) {
	switch test {
	case 1:
		fmt.Printf("test: %d", test)
	case 2:
		fmt.Printf("test: %d", test)
	case 3:
		fmt.Printf("test: %d", test)
	}
}
