package main

import "fmt"

func main(){
	var b int

	fmt.Printf("请输入b的值!")
	fmt.Scan(&b)

	if b == 0 {
		fmt.Printf("我是真的哦~")
	} else if b <= 100 {
		fmt.Printf("我现在小于200~")
	} else if b >= 200 {
		main()
	}
}