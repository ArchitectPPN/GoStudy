package main

import "fmt"

var userName string = "Jack"

func test() {

}

func main() {
	// 局部变量和全局变量重名, 优先使用局部变量
	userName := 111
	fmt.Printf("%v", userName)
}
