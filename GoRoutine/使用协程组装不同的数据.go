package main

import (
	"fmt"
	"sync"
)

type UserInfo struct {
	Username string
	Age      int
	Height   int
}

func main() {
	GetUserInfoInnerOfGoRoutine()
}

// GetUserInfoInnerOfGoRoutine 在协程内部组装数据
func GetUserInfoInnerOfGoRoutine() {
	var wg sync.WaitGroup
	var userInfo UserInfo

	wg.Add(3)
	go func() {
		wg.Done()
		userInfo.Username = "Architect"
	}()

	go func() {
		wg.Done()
		userInfo.Age = 21
	}()

	go func() {
		wg.Done()
		userInfo.Height = 174
	}()

	wg.Wait()

	fmt.Println("用户信息：", userInfo)
}

// GetUserName 获取用户名称
func GetUserName(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	ch <- "Architect"
}

// GetUserAge 获取用户年龄
func GetUserAge(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- 18
}

// GetUserHeight 获取用户身高
func GetUserHeight(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- 171
}

// GetUserInfoOutOfGoRoutine 在协程外面组装数据
func GetUserInfoOutOfGoRoutine() {
	var wg sync.WaitGroup
	var userInfo UserInfo
	var getUserNameCh chan string
	getUserNameCh = make(chan string)
	var getUserAgeCh chan int
	getUserAgeCh = make(chan int)
	var getUserHeightCh chan int
	getUserHeightCh = make(chan int)

	wg.Add(3)
	go GetUserName(&wg, getUserNameCh)
	go GetUserAge(&wg, getUserAgeCh)
	go GetUserHeight(&wg, getUserHeightCh)

	userInfo.Username = <-getUserNameCh
	userInfo.Age = <-getUserAgeCh
	userInfo.Height = <-getUserHeightCh

	wg.Wait()

	// 关闭chan
	close(getUserNameCh)
	close(getUserAgeCh)
	close(getUserHeightCh)

	fmt.Println("用户信息：", userInfo)
}
