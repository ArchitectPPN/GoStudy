package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waitRoutine sync.WaitGroup

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
	waitRoutine.Done()
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
	waitRoutine.Done()
}

func main() {
	runtime.GOMAXPROCS(2) // 只占用一个CPU核心
	waitRoutine.Add(2)
	go a()
	go b()
	waitRoutine.Wait()
}
