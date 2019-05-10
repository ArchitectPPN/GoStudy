package main

import (
	"fmt"
	"math"
)

func main(){
	func(data string){
		fmt.Println(data, ",你好~")
	}("阿牛")

	result := func(data float64) float64 {
		return math.Sqrt(data)
	}(250)
	fmt.Println(result)

	myfunc := func(data string) string {
		return data
	}
	fmt.Println(myfunc("欢迎学习Go语言~"))
}
