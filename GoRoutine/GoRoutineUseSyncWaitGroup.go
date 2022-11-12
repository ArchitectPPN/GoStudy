package main

import (
	"fmt"
	"sync"
)

// 声明全局等待组变量
var goRoutine sync.WaitGroup

func GoRoutineHello() {
	defer goRoutine.Done()
	fmt.Println("Go Routine Hello World~")
}

func GoRoutineUserName() {
	defer goRoutine.Done()
	fmt.Println("Go Routine Hello User~")
}

func main() {
	goRoutine.Add(2) // 登记1个goroutine
	go GoRoutineHello()
	go GoRoutineUserName()

	fmt.Println("你好")

	goRoutine.Wait() // 阻塞等待登记的goroutine完成
}
