package main

import (
	"fmt"
)

func main()  {
	student := map[string]int{
		"age":21,
		"height":180,
	}

	fmt.Println(student)

	teacher := make(map[string]int, 2)
	teacher["age"] = 90
	teacher["height"] = 80
	teacher["weight"] = 99

	fmt.Println(teacher)
}