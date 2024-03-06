package main

import "fmt"

type persons struct {
	name string
	age  int
	gender string
	hobby []string
}

// 匿名结构体


func main()  {
	var human persons
	human.name = "architect"
	human.age = 21
	human.gender = "man"
	human.hobby = []string{"Apex", "OW"}
	fmt.Printf("%v\n", human)

	var noNameStruct struct {
		name string
		age int
	}

	noNameStruct.name = "no name"
	noNameStruct.age = 21
}