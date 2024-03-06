package main

import "fmt"

func main() {
	f1()
	f2()
	f3()
	f4()
}

func f1() {
	fmt.Printf("F1\n")
	defer func() {
		err := recover()// 收集错误
		fmt.Printf("%v\n", err)
	}()
	panic("我出错了!")
}

func f2() {
	fmt.Printf("F2\n")
}

func f3() {
	fmt.Printf("F3\n")
}

func f4() {
	fmt.Printf("F4\n")
}
