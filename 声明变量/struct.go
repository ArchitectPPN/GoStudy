package main

import "fmt"

type Man struct {
	name string
	age int
	sex string 
}

type Woman struct {
	name string 
	age int 
	sex string 
}

func main(){

	var XM Man
	XM.name = "源氏"
	XM.age = 30
	XM.sex = "男"

	fmt.Printf("姓名: %-3s 年龄: %-3d 性别: %-3s\n", XM.name, XM.age, XM.sex)

	var JZ Woman

	JZ.name = "天使"
	JZ.age = 21
	JZ.sex = "女"

	fmt.Printf("姓名: %-3s 年龄: %-3d 性别: %-3s", JZ.name, JZ.age, JZ.sex)
}