package main

import (
	"fmt"
	"sync"
)

var (
	x int64
	wgN sync.WaitGroup
	lockSync sync.Mutex
)

func add()  {
	for i:= 0; i<5000; i++ {
		lockSync.Lock()
		x = x+1
		lockSync.Unlock()
	}
	wgN.Done()
}

// 并发同步和锁
func main() {
	wgN.Add(2)
	go add()
	go add()
	wgN.Wait()
	fmt.Println(x)
}