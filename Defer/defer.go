package main

import "fmt"

func main()  {
	testDefer()
	fmt.Println(f4())
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)

	return 5
}

func echoString()  {
	fmt.Println("我是defer")
}

func testDefer()  {
	fmt.Println("我执行中~")
	defer echoString()
}
