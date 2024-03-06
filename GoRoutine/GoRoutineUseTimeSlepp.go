package main

import (
	"fmt"
	"time"
)

func helloWorld() {
	fmt.Println("GoRoutine Hello World~")
}

func main() {
	go helloWorld()

	fmt.Println("Main Hello World~")
	// 使用sleep等待 go routine 执行， 不然可能无法输出 GoRoutine Hello World~
	// main 执行完毕之后， 就会终止我们创建的 go routine
	time.Sleep(time.Second)
}
