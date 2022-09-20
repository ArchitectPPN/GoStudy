package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lockX int64
	lockWg sync.WaitGroup
	lockM sync.Mutex
	lockRw sync.RWMutex
)

// 读写互斥锁
func read() {
	lockM.Lock()
	time.Sleep(time.Millisecond)
	lockM.Unlock()
	lockWg.Done()
}

func write()  {
	lockM.Lock()
	lockX = lockX+1
	time.Sleep(time.Millisecond * 10)
	lockM.Unlock()
	lockWg.Done()
}

// 互斥锁
func rwLockRead() {
	lockRw.RLock()
	time.Sleep(time.Millisecond)
	lockRw.RUnlock()
	lockWg.Done()
}

func rwLockWrite() {
	lockRw.Lock()
	time.Sleep(time.Millisecond * 10)
	lockRw.Unlock()
	lockWg.Done()
}

func main()  {
	start := time.Now()

	for i:=0; i<1000;i++ {
		lockWg.Add(1)
		//go read()
		go rwLockRead()
	}

	for j:=0; j<10; j++ {
		lockWg.Add(1)
		//go write()
		go rwLockWrite()
	}

	lockWg.Wait()
	fmt.Println(time.Now().Sub(start))
}
