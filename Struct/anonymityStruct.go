package main

import "fmt"

func main()  {
	var student struct{Name string; age int}

	student.Name = "Architect"
	student.age = 19

	fmt.Println(student)
}
