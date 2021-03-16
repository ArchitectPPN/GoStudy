package main

import "fmt"

type person struct {
	name, city string
	age        int8
}

func main() {
	var student = new(person)

	student.name = "architect"
	student.age = 19
	student.city = "上海"

	fmt.Println(student)
}
