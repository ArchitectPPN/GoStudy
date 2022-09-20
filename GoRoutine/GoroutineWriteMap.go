package main

import (
	"fmt"
	"sync"
)

var (
	m = make(map[int]int)
	wgWrite = sync.WaitGroup{}
	m2 = sync.Map{}
)

func get(mKey int)int  {
	return m[mKey]
}

func set(mKey int, mVal int)  {
	m[mKey] = mVal
}

//func main()  {
//	for i:=0; i < 20; i++ {
//		wgWrite.Add(1)
//		go func(i int) {
//			// 设置键值对
//			set(i, i+100)
//			// 获取值
//			fmt.Printf("键 %v 值%v \n", i, get(i))
//			wgWrite.Done()
//		}(i)
//	}
//	wgWrite.Wait()
//}

func main() {
	for i:=0; i < 20; i++ {
		wgWrite.Add(1)
		go func(i int) {
			// 设置键值对
			m2.Store(i, i+100)
			// 获取值
			val, _ := m2.Load(i)
			fmt.Printf("键 %v 值%v \n", i, val)
			wgWrite.Done()
		}(i)
	}
	wgWrite.Wait()
}