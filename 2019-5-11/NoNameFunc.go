package main

import "fmt"

func main(){
	f := func(str string){
		fmt.Println(str)
	}
	f("我是阿牛~")

	for i:=0; i<10; i++{
		fmt.Println("i=:", i)
		fmt.Println(add(i))
	}

}

// 实现计数函数
func add(num int) int {
	sum := 0
	sum += num
	return sum
}



