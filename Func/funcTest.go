package main

import "fmt"

func main() {
	noNameFunc(yuanShu, "nihao")
	ret := returnNoNameFunc()

	ret(9, 20)
	transfer := transferFunc(yuanShu, "architect")
	low(transfer)
}

func yuanShu(name string) {
	fmt.Printf("%v\n", name)
}

func noNameFunc(f func(string), name string) {
	f(name)
}

func transferFunc(f func(string), name string) func() {
	return func() {
		f(name)
	}
}

// 返回一个匿名函数
func returnNoNameFunc() func(int, int) int {
	return func(x, y int) int {
		fmt.Printf("%v\n", x+y)
		return 9
	}
}

func low(f func()) {
	f()
}